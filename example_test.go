package gs

import (
	"bytes"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/compile"
	"github.com/ycl2018/gs/vm"
)

type MyEnv struct {
	Embed
	A           string
	B           *string
	Map         map[any]any
	Slice       []any
	StringSlice []string
	StructMap   map[*MyEnv]string
}

type Embed struct {
	D string
	C int
}

type BasicValue struct {
	Uint        uint
	Uint8       uint8
	Uint16      uint16
	Uint32      uint32
	Uint64      uint64
	Int8        int8
	Int16       int16
	Int32       int32
	Int         int
	Int64       int64
	Float32     float32
	Float64     float64
	String      string
	SliceAny    []any
	SliceString []string
	MapAny      map[any]any
}

func TestGsInterpreter_Interp(t *testing.T) {
	tests := []struct {
		name    string
		program string
		env     any
		expect  string
		trace   bool
	}{
		{
			name: "apple.gs",
			program: `
i = 0
for i<5 {
	i = i + 1
	if i<3 {
		println( i , " is less than 3" )
	} else {
		println( "foo" )
	}
}
	`,
			expect: `
1  is less than 3
2  is less than 3
foo
foo
foo
`},
		{
			name: "cheery.gs",
			program: `
func f(x) {return 2*x}
println( f(4) )
`,
			expect: "8",
		},
		{
			name: "multi_return.gs",
			program: `
func swap (x,y) {
	return y, x
}

println( swap(1,2) )
x,y,z,w = swap(1,2),swap(3,4)
println( x,y,z,w ) // 2143
`,
			expect: `
(2,1)
2 1 4 3
`,
		},
		{
			name: "embeded.gs",
			program: `
$.D = "chenglong"
$.C = 5.2
println( $.D )
println( $.C )
`,
			expect: `
chenglong
5
`,
			env: &MyEnv{},
		},
		{
			name: "factorials.gs",
			program: `
func fact(n) {
	if n < 2 {
		return 1
	}
	return n * fact(n-1)
}

println( fact(10) )
`,
			expect: "3628800",
		},
		{
			name: "forward.gs",
			program: `
println( f(4) )           // references definition on next line
func f(x) {return 2*x}
println( new User{} )       // references definition on next line
type User struct { name, password }
`,
			expect: `
8
struct User { name: <nil>, password: <nil> }
`,
		},
		{
			name: "localstruct.gs",
			program: `
type User struct { name, password } // define global struct
func f() {                       // define f
    type User struct { x, y }       // hides global User def
    u = new User{}               // create new User instance, put in local u
    println( u )                    // println(s "{x=null, y=null}"
 }                            // end body of f
println( new User{} )                 // println(s "{name=null, password=null}"
f()                            // call f
`,
			expect: `
struct User { name: <nil>, password: <nil> }
struct User { x: <nil>, y: <nil> }
`,
		},
		{
			name: "lookup.gs",
			program: `
x = 1           // create global variable
func f(x) {       // define f in global space
    println( x )     // access parameter; println(s 10
    y = 2       // create local variable
}
func g() {        // define g in global space
    x = 3       // set global variable
}
f(10)
g()
println( x )         // println(s 3 (g alters global value)
`,
			expect: `
10
3
`,
		},
		{
			name: "loop.gs",
			program: `
n = 100
i = 0
for i<n {
	i = i + 1
}
println( "looped ",n," times." )
`,
			expect: "looped  100  times.",
		},
		{
			name: "struct.gs",
			program: `
type User struct { name, password }
u = new User{}
u.name = "chenglong"
println( "Login: "+u.name )
println( u )
`,
			expect: `
Login: chenglong
struct User { name: chenglong, password: <nil> }
`,
		},
		{
			name: "nested.gs",
			program: `
type User struct { name, addr }
type Address struct { street, city, state, zip }
u = new User{}
u.name = "parrt"
addr = new Address{}
addr.street = "123 Main St"
u.addr =  addr
u.addr.city = "Chicago"
println( u.addr.street )
println( u.addr.city )
println( u.addr )
`,
			expect: `
123 Main St
Chicago
struct Address { street: 123 Main St, city: Chicago, state: <nil>, zip: <nil> }
 `,
		},
		{
			name: "slice.gs",
			program: `
arr = []
for i = range 10 {
	arr = append(arr, i)
}
println(arr)
println(arr[0])
arr[0] = "a"
println(arr)
`,
			expect: `
[0 1 2 3 4 5 6 7 8 9]
0
[a 1 2 3 4 5 6 7 8 9]
`,
		},
		{
			name: "range_kv.gs",
			program: `
c = {"a": 1, "b": 2, "c": 3}
sum = 0
for k, v = range c {
	sum += v
}
println( sum )
	`,
			expect: `
6
`,
		},
		{
			name: "cstyle_for.gs",
			program: `
for i = 0; i<3; i = i + 1 {
	println( i )
}
	`,
			expect: `
0
1
2
`,
		},
		{
			name: "continue_break.gs",
			program: `
for i = range 10 {
	if i %2 == 0 {
		continue
	} else {
		if i == 7 {
			break
		}
		println( i )
	}
	
}
	`,
			expect: `
1
3
5`,
		},
		{
			name: "if_else.gs",
			program: `
a = 9
b = 2
if i = a + b; i < 10 {
	println( i )
} else {
	println( "i is not less than 10" )
}
	`,
			expect: `i is not less than 10`,
		},
		{
			name: "slice_split.gs",
			program: `
s = ["a", "b", "c"]
println( s[0] )
println( s[1:] )
println( s[:2] )
println( s[1:2] )
println( s[:] )
	`,
			expect: `
a
[b c]
[a b]
[b]
[a b c]
`,
		},
		{
			name: "map.gs",
			program: `
d = {"a": 1, "b": 2, "c": 3}
println( d["a"] )
k, v = 1 + 2, d["a"]
mm = {k:v}
println( mm )
	`,
			expect: `
1
map[3:1]
`,
		},
		{
			name: "calculate.gs",
			program: `
a = 1
b = 2
println( a + b )
println( a - b )
println( a * b )
println( a / b )
println( a % b )
println( a < b )
println( a > b )
println( a <= b )
println( a >= b )
println( a != b )
println( a == b )
println( -a )
println( a & b )
println( a | b )
println( a ^ b )
		`,
			expect: `
3
-1
2
0
1
true
false
true
false
true
false
-1
0
3
3
`,
		},
		{
			name: "expression2.gs",
			program: `
a = 1
b = 2
c = a + b == 0 || a + b == 1 && a + b == 3
println( c )
		`,
			expect: "false",
		},
		{
			name: "self_update.gs",
			program: `
a = 0
a+=1
println( a )
a-=2
println( a )
a*=3
println( a )
a%=4
println( a )
		`,
			expect: `
1
-1
-3
-3
`,
		},
		{
			name: "incr_decr.gs",
			program: `
a = 0
a++
println( a )
a--
println( a )
		`,
			expect: `
1
0
`,
		},
		{
			name: "safe_access.gs",
			program: `
type Student struct {
x,y,z
}

s = new Student{}
a = s.x?.z
println( a )
`,
			expect: "<nil>",
		},
		{
			name: "const_optimizer.gs",
			program: `
a= (1*2+3*3) + 100 // 111
println( a )
`,
			expect: "111",
		},
		{
			name: "slice.gs",
			program: `
b=["111","222","333"]
for i = range b {
	println( b[i] )
}
for i, v = range b {
 println( i," ",v )
}
`,
			expect: `
111
222
333
0   111
1   222
2   333
`,
		},
		{
			name: "env_map.gs",
			program: `
println( $["a"]["name"] + $["a"]["addr"] )
`,
			env: map[string]any{
				"a": map[string]any{
					"name": "chenglong",
					"addr": "123 Main St",
				},
			},
			expect: `chenglong123 Main St`,
		},
		{
			name: "env_slice.gs",
			env:  []any{"a", "b", "c"},
			program: `
println( $[0]+$[1]+$[2] )
`,
			expect: `abc`,
		},
		{
			name: "env_struct.gs",
			env: &struct {
				A string
				B string
				C map[string]string
			}{
				A: "a",
				B: "b",
				C: map[string]string{},
			},
			program: `
$.C["A+B"] = $.A + $.B
println( $.C["A+B"] )
`,
			expect: `ab`,
		},
		{
			name: "pointer_access.gs",
			program: `
a = "chenglong"
*$.B = a
println( *$.B )
		`,
			env: &MyEnv{
				A: "a",
				B: nil,
			},
			expect: `chenglong`,
		},
		{
			name: "map.gs",
			program: `
m = {"a": 1, "b": 2, "c": 3}
$.Map = m
println( $.Map )
m["d"] = 4
println( "$.Map=", $.Map )
println( "m=", m )
m = [1,2,3]
println(m)
println($.Map)
		`,
			env: &MyEnv{
				A:   "a",
				Map: nil,
			},
			expect: `
map[a:1 b:2 c:3]
$.Map= map[a:1 b:2 c:3 d:4]
m= map[a:1 b:2 c:3 d:4]
[1 2 3]
map[a:1 b:2 c:3 d:4]
`,
		},
		{
			name: "copy_map.gs",
			program: `
for k,v = range $.Map {
	$.StructMap[k] = v
}
vsum = ""
for k,v = range $.StructMap {
	vsum += v
}

println( vsum == "AB" || vsum == "BA" )
		`,
			expect: `
true
`,
			env: &MyEnv{
				Map: map[any]any{
					&MyEnv{
						A: "A",
					}: "A",
					&MyEnv{
						A: "B",
					}: "B",
				},
				StructMap: make(map[*MyEnv]string),
			},
		},
		{
			name: "set_value.gs",
			program: `
$.Uint     = 1
$.Uint8    = 2
$.Uint16   = 3
$.Uint32   = 4
$.Uint64   = 5
$.Int8     = 6
$.Int16    = 7
$.Int32    = 8
$.Int	   = 9
$.Int64    = 10
$.Float32  = .2
$.Float64  = .3
$.String   = "str"
println($.Uint + $.Uint8) // 3
println($.Int + $.Int8) // 15
println($.Uint + $.Float32) // 1.2
println($.String + $.String) // strstr
$.SliceAny = append($.SliceAny, $.String)
println($.SliceAny) // [str]
$.SliceAny = [1,"sss",.2]
println($.SliceAny) // [1 sss 0.2]
$.MapAny = {1:"1",true:false,"string":"string"}
println( $.MapAny )
$.SliceString = append($.SliceString, "a")
println($.SliceString) // [a]
`,
			env: &BasicValue{},
			expect: `
3
15
1.2000000029802322
strstr
[str]
[1 sss 0.2]
map[string:string 1:1 true:false]
[a]
`,
		},
		{
			name: "builtin.gs",
			program: `
println (len($.A)) // 3
println (len($.Map)) // 3
println (len($.StringSlice)) // 0
println($.Map)
delete($.Map, "1")
println($.Map)
delete($.Map, uint(4))
println($.Map)
$.Slice = append($.Slice, 4)
println($.Slice)
$.Slice[0] = 4
print("print:",$.Slice,"\n")
println("println:",$.Slice)
printf("printf:$.Slice=%v\n", $.Slice)
` + "println(`raw literal:\nfirst line\nsecond line\n\"哈喽\"`)",
			env: &MyEnv{
				A:           "chenglong",
				Map:         map[any]any{"1": 1, "2": 2, "3": 3, uint(4): "4"},
				Slice:       []any{1, 2, 3},
				StringSlice: nil,
				StructMap:   nil,
			},
			expect: `
9
4
0
map[1:1 2:2 3:3 4:4]
map[2:2 3:3 4:4]
map[2:2 3:3]
[1 2 3 4]
print:[4 2 3 4]
println: [4 2 3 4]
printf:$.Slice=[4 2 3 4]
raw literal:
first line
second line
"哈喽"
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			p := compile.NewGsCompiler()
			code, err := p.Compile(antlr.NewInputStream(tt.program), tt.env)
			if err != nil {
				t.Fatal(err)
			}
			var ops = []vm.Option{
				vm.WithEnv(tt.env), vm.WithPrintTo(out),
			}
			if tt.trace {
				ops = append(ops, vm.WithEnableTrace())
			}
			interpreter := vm.NewInterpreter(code, ops...)
			interpreter.Run()
			want, got := strings.TrimSpace(tt.expect), strings.TrimSpace(out.String())
			if want != got {
				t.Errorf("got:\n%s\nwant:\n%s", got, tt.expect)
			}
		})
	}
}
