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
	A         string
	B         *string
	Map       map[any]any
	Slice     []any
	StructMap map[*MyEnv]string
}

type BasicValue struct {
	Uint     uint
	Uint8    uint8
	Uint16   uint16
	Uint32   uint32
	Uint64   uint64
	Int8     int8
	Int16    int16
	Int32    int32
	Int      int
	Int64    int64
	Float32  float32
	Float64  float64
	String   string
	SliceAny []any
	MapAny   map[any]any
}

func TestGsInterpreter_Interp(t *testing.T) {
	tests := []struct {
		name    string
		program string
		env     any
		expect  string
	}{
		{
			name: "apple.gs",
			program: `
i = 0
for i<5 {
	i = i + 1
	if i<3 {
		print i , " is less than 3"
	}else {
		print "foo"
	}
}
	`,
			expect: `
1 is less than 3
2 is less than 3
foo
foo
foo
`},
		{
			name: "cheery.gs",
			program: `
func f(x) {return 2*x}
print f(4)
`,
			expect: "8",
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

print fact(10)
`,
			expect: "3628800",
		},
		{
			name: "forward.gs",
			program: `
print f(4)           // references definition on next line
func f(x) {return 2*x}
print new User{}       // references definition on next line
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
    print u                    // prints "{x=null, y=null}"
 }                            // end body of f
print new User{}                 // prints "{name=null, password=null}"
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
    print x     // access parameter; prints 10
    y = 2       // create local variable
}
func g() {        // define g in global space
    x = 3       // set global variable
}
f(10)
g()
print x         // prints 3 (g alters global value)
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
print "looped ",n," times."
`,
			expect: "looped 100 times.",
		},
		{
			name: "struct.gs",
			program: `
type User struct { name, password }
u = new User{}
u.name = "chenglong"
print "Login: "+u.name
print u
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
print u.addr.street
print u.addr.city
print u.addr
`,
			expect: `
123 Main St
Chicago
struct Address { street: 123 Main St, city: Chicago, state: <nil>, zip: <nil> }
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
print sum
	`,
			expect: `
6
`,
		},
		{
			name: "cstyle_for.gs",
			program: `
for i = 0; i<3; i = i + 1 {
	print i
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
		print i
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
	print i
} else {
	print "i is not less than 10"
}
	`,
			expect: `i is not less than 10`,
		},
		{
			name: "slice_split.gs",
			program: `
s = ["a", "b", "c"]
print s[0]
print s[1:]
print s[:2]
print s[1:2]
print s[:]
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
print d["a"]
	`,
			expect: "1",
		},
		{
			name: "calculate.gs",
			program: `
a = 1
b = 2
print a + b
print a - b
print a * b
print a / b
print a % b
print a < b
print a > b
print a <= b
print a >= b
print a != b
print a == b
print -a
print a & b
print a | b
print a ^ b
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
print c
		`,
			expect: "false",
		},
		{
			name: "self_update.gs",
			program: `
a = 0
a+=1
print a
a-=2
print a
a*=3
print a
a%=4
print a
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
print a
a--
print a
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
print a
`,
			expect: "<nil>",
		},
		{
			name: "const_optimizer.gs",
			program: `
a= (1*2+3*3) + 100 // 111
print a
`,
			expect: "111",
		},
		{
			name: "slice.gs",
			program: `
b=["111","222","333"]
for i = range b {
	print b[i]
}
for i, v = range b {
 print i," ",v
}
`,
			expect: `
111
222
333
0 111
1 222
2 333
`,
		},
		{
			name: "env_map.gs",
			program: `
print $["a"]["name"] + $["a"]["addr"]
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
print $[0]+$[1]+$[2]
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
print $.C["A+B"]
`,
			expect: `ab`,
		},
		{
			name: "pointer_access.gs",
			program: `
a = "chenglong"
*$.B = a
print *$.B
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
print $.Map
m["d"] = 4
print "$.Map=", $.Map
print "m=", m
		`,
			env: &MyEnv{
				A:   "a",
				Map: nil,
			},
			expect: `
map[a:1 b:2 c:3]
$.Map=map[a:1 b:2 c:3 4:d]
m=map[a:1 b:2 c:3 4:d]
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

print vsum == "AB" || vsum == "BA"
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
$.Uint 	   = 1
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
$.SliceAny = [1,"sss",.2]
$.MapAny   = {1:"1",true:false,"string":"string"}
print $
`,
			env: &BasicValue{},
			expect: `
&{1 2 3 4 5 6 7 8 9 10 0.2 0.3 str [1 sss 0.2] map[string:string 1:1 true:false]}
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
			vm.NewInterpreter(code, vm.WithEnv(tt.env), vm.WithPrintTo(out)).Run()
			got := out.String()
			if strings.TrimSpace(got) != strings.TrimSpace(tt.expect) {
				t.Errorf("got:\n%s\nwant:\n%s", got, tt.expect)
			}
		})
	}
}
