package gs

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/ycl2018/gs/consts"
)

type MyEnv struct {
	Embed
	A           string
	B           *string
	Arr         [3]int
	Map         map[any]any
	Slice       []any
	StringSlice []string
	StructMap   map[*MyEnv]string
	Fn          any
	StringAlias StringAlias
	Itf         Itf
	MyEnv       *MyEnv
}

type StringAlias string

type Itf interface {
	Hello() string
}

type MyItf struct {
	Name string
}

func (m *MyItf) Hello() string {
	return "Hello " + m.Name
}

func (m *MyItf) Hello2() string {
	return "Hello2 " + m.Name
}

func (env *MyEnv) SayHello() string {
	return "Hello World " + env.A
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
	Byte        byte
	SliceAny    []any
	SliceString []string
	MapAny      map[any]any
	Bool        bool
}

func TestGsInterpreter_Interp(t *testing.T) {
	tests := []struct {
		name       string
		program    string
		env        any
		expect     string
		expectErr  string
		defineFunc []Func
		fastFunc   []FastFunc
		trace      bool
		dump       bool
		ret        any
		onlyRunEnv bool
	}{
		{
			name: "apple.gs",
			program: `
i = 0
for i<5 {
	i = i + 1
	if i<3 {
		println(i , " is less than 3")
	} else {
		println("foo")
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
println(f(4))
`,
			expect: "8",
		},
		{
			name: "multi_return.gs",
			program: `
func swap (x,y) {
	return y, x
}

println(swap(1,2))
x,y,z,w = swap(1,2),swap(3,4)
println(x,y,z,w) // 2143
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
println($.D)
println($.C)
println($.MyEnv.A)
*$.MyEnv.B = "chenglong"
println(*$.MyEnv.B)
`,
			expect: `
chenglong
5
chenglong
chenglong
`,
			env: &MyEnv{MyEnv: &MyEnv{
				A: "chenglong",
			}},
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

println(fact(10))
`,
			expect: "3628800",
		},
		{
			name: "forward.gs",
			program: `
println(f(4))           // references definition on next line
func f(x) {return 2*x}
println(new User{})       // references definition on next line
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
type User struct { name, password } 	// define global struct
func f() {                       		// define f
    type User struct { x, y }       	// hides global User def
    u = new User{}               		// create new User instance, put in local u
    println(u)                    		// struct User { name: <nil>, password: <nil> }
}                            			// end body of f
println(new User{})                 	// struct User { x: <nil>, y: <nil> }
f()                            			// call f
`,
			expect: `
struct User { name: <nil>, password: <nil> }
struct User { x: <nil>, y: <nil> }
`,
		},
		{
			name: "lookup.gs",
			program: `
x = 1           	// create global variable
func f(x) {
    println(x)     	// access parameter; println 10
	x += 4			// local x
	println(x) 		// 14
}
func g() {
    x = 3       	// set local variable
}
func setGlobalX() {
	global x		// declare global x
	x = 100
}
f(10)
g()
println(x)         // 1
setGlobalX()
println(x)         // 100
`,
			expect: `
10
14
1
100
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
println("looped ",n," times.")
`,
			expect: "looped  100  times.",
		},
		{
			name: "struct.gs",
			program: `
type User struct { name, password }
u = new User{}
u.name = "chenglong"
println("Login: "+u.name)
println(u)
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
println(u.addr.street)
println(u.addr.city)
println(u.addr)
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
println(sum)
	`,
			expect: `
6
`,
		},
		{
			name: "cstyle_for.gs",
			program: `
for i = 0; i<3; i = i + 1 {
	println(i)
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
		println(i)
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
	println(i)
} else {
	println("i is not less than 10")
}
	`,
			expect: `i is not less than 10`,
		},
		{
			name: "slice_split.gs",
			program: `
s = ["a", "b", "c"]
println(s[0])
println(s[1:])
println(s[:2])
println(s[1:2])
println(s[:])
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
			env: &MyEnv{
				Map: make(map[any]any),
			},
			program: `
d = {"a": 1, "b": 2, "c": 3}
println(d["a"])
k, v = 1 + 2, d["a"]
mm = {k:v}
println(mm)
println($.Map["a"])
$.Map["a"] = 1
println($.Map["a"])
	`,
			expect: `
1
map[3:1]
<nil>
1
`,
		},
		{
			name: "calculate.gs",
			program: `
a = 1
b = 2
println(a + b)
println(a - b)
println(a * b)
println(a / b)
println(a % b)
println(a < b)
println(a > b)
println(a <= b)
println(a >= b)
println(a != b)
println(a == b)
println(-a)
println(a & b)
println(a | b)
println(a ^ b)
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
			name: "or.gs",
			program: `
a = 1
b = 2
println(a == 2) // false
println(a + b == 3) // true
println(a + b == 1 || a + b == 2 || a == 3) // false
println(a + b == 1 || a + b == 2 || a + b == 3) // true
println(false || true)// true
		`,
			expect: `
false
true
false
true
true
`,
		},
		{
			name: "and.gs",
			program: `
a = 1
b = 2
println(a + b == 3 && a - b == -1) // true
println(a + b == 1 && a + b == 3) // false
println(a + b == 3 && a - b == -1 && a == 0) // false
println(true && a == 2)
`,
			expect: `
true
false
false
false
`,
		},
		{
			name: "self_update.gs",
			program: `
a = 0
a+=1
println(a)
a-=2
println(a)
a*=3
println(a)
a%=4
println(a)
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
println(a)
a--
println(a)
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
a = s.x
println(a)
`,
			expect: "<nil>",
		},
		{
			name: "const_optimizer.gs",
			program: `
a= (1*2+3*3) + 100 // 111
println(a)
`,
			expect: "111",
		},
		{
			name: "slice.gs",
			program: `
b=["111","222","333"]
for i = range b {
	println(b[i])
}
for i, v = range b {
 println(i," ",v)
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
println($["a"]["name"] + $["a"]["addr"])
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
println($[0]+$[1]+$[2])
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
println($.C["A+B"])
`,
			expect: `ab`,
		},
		{
			name: "pointer_access.gs",
			program: `
a = "chenglong"
*$.B = a
println(*$.B)
		`,
			env: &MyEnv{
				A: "a",
				B: nil,
			},
			expect: `chenglong`,
		},
		{
			name: "env_slice.gs",
			program: `
println($.Arr[0])
println($.Arr[1])
println($.Arr[2])
println($.Arr[:2][0])
		`,
			env: &MyEnv{
				Arr: [3]int{1, 2, 3},
			},
			expect: `
1
2
3
1
`,
		},
		{
			name: "map.gs",
			program: `
m = {"a": 1, "b": 2, "c": 3}
$.Map = m
println($.Map)
m["d"] = 4
println("$.Map=", $.Map)
println("m=", m)
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

println(vsum == "AB" || vsum == "BA")
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
			name: "convert.gs",
			program: `
println(int($.Int))
println(uint($.Int))
println(uint8($.Int))
println(uint16($.Int))
println(uint32($.Int))
println(uint64($.Int))
println(int32($.Int))
println(int64($.Int))
println(int8($.Int))
println(int16($.Int))
println(float32($.Int))
println(float64($.Int))
println(string($.String))
println(bool($.Bool))
`,
			env: &BasicValue{
				Int:    12,
				String: "str",
			},
			expect: `
12
12
12
12
12
12
12
12
12
12
12
12
str
false
`,
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
$.Bool     = true
$.Byte     = 12
println($.Uint + $.Uint8) // 3
println($.Int + $.Int8) // 15
println($.Uint + $.Float32) // 1.2
println($.String + $.String) // strstr
println($.Byte)
println($.Bool)
$.SliceAny = append($.SliceAny, $.String, "hello")
println($.SliceAny) // [str]
$.SliceAny = [1,"sss",.2]
println($.SliceAny) // [1 sss 0.2]
$.MapAny = {1:"1",true:false,"string":"string"}
println($.MapAny)
$.SliceString = append($.SliceString, "a")
println($.SliceString) // [a]
`,
			env: &BasicValue{},
			expect: `
3
15
1.2000000029802322
strstr
12
true
[str hello]
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
println (len($.StringSlice)) // 2
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
str = sprintf("hello:%s", "chenglong")
println(str)
arr = []
arr = append(arr, $.StringSlice...)
println(arr) // [hello world]
arr = ["cheng", "long"]
$.StringSlice = append($.StringSlice, arr...)
println($.StringSlice) // [hello world cheng long]
` + "println(`raw literal:\nfirst line\nsecond line\n\"哈喽\"`)",
			env: &MyEnv{
				A:           "chenglong",
				Map:         map[any]any{"1": 1, "2": 2, "3": 3, uint(4): "4"},
				Slice:       []any{1, 2, 3},
				StringSlice: []string{"hello", "world"},
				StructMap:   nil,
			},
			expect: `
9
4
2
map[1:1 2:2 3:3 4:4]
map[2:2 3:3 4:4]
map[2:2 3:3]
[1 2 3 4]
print:[4 2 3 4]
println: [4 2 3 4]
printf:$.Slice=[4 2 3 4]
hello:chenglong
[hello world]
[hello world cheng long]
raw literal:
first line
second line
"哈喽"
`,
		},
		{
			name: "err.gs",
			program: `
func divide(a, b) {
	return a / b
}
divide(1, 0)
`,
			expectErr: `
panic: division by zero
stack trace:
	at divide args:(1, 0) line:3
	at main args:() line:5
`,
		},
		{
			name: "outer_func.gs",
			program: `
ret = $.Fn("hello"," world")
println(ret)
println($.SayHello())
println($.Map["fn"]("hello", " world"))
`,
			env: &MyEnv{
				A:  "chenglong",
				Fn: func(a, b string) string { return a + b },
				Map: map[any]any{
					"fn": func(a, b string) string {
						return a + b
					},
				},
			},
			expect: `
hello world
Hello World chenglong
hello world
`,
		},
		{
			name: "define_funcs.gs",
			defineFunc: []Func{
				{
					Name: "add",
					Fn: func(a, b int) int {
						return a + b
					},
				},
				{
					Name: "sub",
					Fn: func(a, b int) int {
						return a - b
					},
				},
			},
			program: `
println(sub(1,2))
sum = add(1,2)
println(sum)
`,
			expect: `
-1
3
`,
		},
		{
			name: "define_fastfuncs.gs",
			fastFunc: []FastFunc{
				{
					Name:   "add",
					NumIn:  2,
					NumOut: 1,
					Fn: func(args []any) []any {
						return []any{args[0].(int) + args[1].(int)}
					},
				},
				{
					Name:   "sub",
					NumIn:  2,
					NumOut: 1,
					Fn: func(args []any) []any {
						return []any{args[0].(int) - args[1].(int)}
					},
				},
			},
			program: `
println(sub(1,2))
sum = add(1,2)
println(sum)
`,
			expect: `
-1
3
`,
		},
		{
			name: "eq_err.gs",
			program: `
println($.StringAlias == "hello")
		`,
			env: &MyEnv{
				StringAlias: StringAlias("hello"),
			},
			expectErr: `
panic: invalid operation '==' (mismatched types gs.StringAlias and string)
stack trace:
	at main args:() line:2
`,
		},
		{
			name: "alias.gs",
			program: `
println(string($.StringAlias) == "hello")
		`,
			env: &MyEnv{
				StringAlias: StringAlias("hello"),
			},
			expect: `true`,
		},
		{
			name: "nil.gs",
			program: `
println($.Itf == nil)
println($.B == nil)
println($.Map == nil)
println($.Slice == nil)
println($.Fn == nil)
		`,
			env: &MyEnv{
				Itf: nil,
			},
			expect: `
true
true
true
true
true
`,
		},
		{
			name: "itf.gs",
			program: `
println($.Itf.Hello() == "Hello World")
		`,
			env: &MyEnv{
				Itf: &MyItf{
					Name: "World",
				},
			},
			expect: `true`,
		},
		{
			name: "itf_hide.gs",
			program: `
println($.Itf.Hello2() == "Hello2 World")
		`,
			env: &MyEnv{
				Itf: &MyItf{
					Name: "World",
				},
			},
			expectErr: `
<line 2> .: syntax error:invalid $.Itf.Hello2 on type:gs.Itf
		 ^`,
		},
		{
			name: "only_run_env.gs",
			program: `
println($.Itf.Hello() == "Hello World")
println($.Itf.Hello() == "Hello World")
		`,
			env: &MyEnv{
				Itf: &MyItf{
					Name: "World",
				},
			},
			expect: `
true
true
`,
			onlyRunEnv: true,
		},
		{
			name: "predefine_func.gs",
			program: `
println(in([1,2,4], 1)) 				// true
println(in($.StringSlice, "hello")) 	// true
println(in($.Map, "foo")) 				// true
println(in($.Map, "bar")) 				// false
println(index($.StringSlice, "hello")) 	// 0
println(index($.StringSlice, "world")) 	// 1
println(index($.StringSlice, "long")) 	// -1
println(hasPrefix("hello", "he")) 		// true
println(hasPrefix("hello", "lo")) 		// false
println(hasSuffix("hello", "lo")) 		// true
println(hasSuffix("hello", "he")) 		// false
println(trim(" ###hello", " #")) 		// hello
println(trimPrefix("hello", "he")) 		// llo
println(trimSuffix("hello", "lo")) 		// hel
println(trimSpace(" hello ")) 			// hello
println(trimLeft(" ###hello", " #")) 	// hello
println(trimRight("hello ###", " #")) 	// hello
println(toLower("HELLO")) 				// hello
println(toUpper("hello")) 				// HELLO
println(split("hello world", " ")) 		// [hello world]
println(join(["hello", "world"], " ")) 	// hello world
println(parseTime("2006-01-02T15:04:05Z07:00", "2025-01-02T15:04:05+07:00")) //(2025-01-02 15:04:05 +0700 +0700,<nil>)
println(parseDuration("1h2m3s")) 		// (1h2m3s,<nil>)
println(duration(111)) 					// 111ns
println(toJson($.StringSlice)) 			// (["hello","world"],<nil>)
println(fromJson("{\"foo\":\"bar\"}")) 	// (map[foo:bar],<nil>)
println(unmarshalJson("{\"foo\":\"bar\"}",$.Map)) // <nil>
		`,
			env: &MyEnv{
				StringSlice: []string{"hello", "world"},
				Map:         map[any]any{"foo": "bar"},
			},
			expect: `
true
true
true
false
0
1
-1
true
false
true
false
hello
llo
hel
hello
hello
hello
hello
HELLO
[hello world]
hello world
(2025-01-02 15:04:05 +0700 +0700,<nil>)
(1h2m3s,<nil>)
111ns
(["hello","world"],<nil>)
(map[foo:bar],<nil>)
<nil>
`,
		},
		{
			name: "return.gs",
			program: `
return 1,2
		`,
			ret:    consts.Tuple{Values: []any{1, 2}, Num: 2},
			expect: ``,
		},
		{
			name: "in.gs",
			program: `
println(in([1,2,3], int($.Uint32))) // true
`,
			env:    &BasicValue{Uint32: 3},
			expect: `true`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			var ops []CompileOption
			if tt.env != nil && !tt.onlyRunEnv {
				ops = append(ops, Env(tt.env))
			}
			if tt.dump {
				ops = append(ops, DumpCode())
			}
			if len(tt.defineFunc) > 0 {
				ops = append(ops, DefineFuncs(tt.defineFunc...))
			}
			if len(tt.fastFunc) > 0 {
				ops = append(ops, DefineFast(tt.fastFunc...))
			}
			code, err := Compile(tt.program, ops...)
			if err != nil {
				if tt.expectErr == "" {
					t.Fatal(err)
				}
				if tt.expectErr != err.Error() {
					return
				}
				return
			}
			var runOps = []RunOption{
				Output(out),
			}
			if tt.trace {
				runOps = append(runOps, Trace())
			}
			ret, err := Run(code, tt.env, runOps...)
			if err != nil && tt.expectErr == "" {
				t.Fatal(err)
				return
			}
			if tt.expectErr != "" {
				if err == nil {
					t.Errorf("expect err: %s, but got nil", tt.expectErr)
					return
				}
				want, got := strings.TrimSpace(tt.expectErr), strings.TrimSpace(err.(*consts.CrashError).CodeTrace)
				if want != got {
					t.Errorf("got:\n%s\nwant:\n%s", got, tt.expectErr)
					return
				}
			} else {
				want, got := strings.TrimSpace(tt.expect), strings.TrimSpace(out.String())
				if want != got {
					t.Errorf("got:\n%s\nwant:\n%s", got, tt.expect)
					return
				}
			}
			if tt.ret != nil {
				if !reflect.DeepEqual(tt.ret, ret.RawValue()) {
					t.Errorf("expect ret: %v, but got %v", tt.ret, ret.RawValue())
					return
				}
			}
		})
	}
}
