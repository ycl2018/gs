package compile

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func TestGsInterpreter_Interp(t *testing.T) {
	tests := []struct {
		name    string
		program string
	}{
		{
			name: "apple.gs",
			program: `
i = 0
for i<10 {
	print i*3.2
	i = i + 1
	if i<5 {
		print i + " is less than 5"
	}else {
		print "foo"
	}
}
	`},
		{
			name: "cheery.gs",
			program: `
func f(x) {return 2*x}
print f(4)
`,
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
		},
		{
			name: "forward.gs",
			program: `
print f(4)           // references definition on next line
func f(x) {return 2*x}
print new User{}       // references definition on next line
type User struct { name, password }
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
		},
		{
			name: "loop.gs",
			program: `
n = 100
i = 0
for i<n {
        i = i + 1
}

print "looped "+n+" times."
`,
		},
		{
			name: "struct.gs",
			program: `
type User struct { name, password }
u = new User{}
u.name = "parrt"
print "Login: "+u.name
print u
`,
		},
		{
			name: "structerr.gs",
			program: `
type User struct { name, password }
u = new User{}
u.name = "parrt"    // make u.name a string
u.name.y = "parrt"  // u.name is a string not a struct
u.x = 3             // x isn't a field of User; can't write to it
print u.x           // check for unknown field in expr as well
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
		},
		{
			name: "range.gs",
			program: `
for i = range 10 {
	print i
}
c = {"a": 1, "b": 2, "c": 3}
for k, v = range c {
	print k, v
}
	`,
		},
		{
			name: "cStyleRange.gs",
			program: `
for i = 0; i<10; i = i + 1 {
	print i
}
	`,
		},
		{
			name: "forRangeBreak.gs",
			program: `
for i = range 10 {
	if i == 5 {
		break
	} else {
		if i == 7 {
			continue
		}
	}
	print i
}
	`,
		},
		{
			name: "ifElse.gs",
			program: `
a = 1
b = 2
if i = a + b; i< 10 {
		print i
} else {
		print "i is not less than 10"
}
	`,
		},
		{
			name: "slice.gs",
			program: `
s = ["a", "b", "c"]
print s[0]
print s[1:]
print s[:2]
print s[1:2]
print s[:]
	`,
		},
		{
			name: "dict.gs",
			program: `
d = {"a": 1, "b": 2, "c": 3}
print d["a"]
	`,
		},
		{
			name: "expression.gs",
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
print a || b
print a && b
print a ** b
print !a
print -a
print a & b
print a | b
print a ^ b
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
		},
		{
			name: "selfUpdate.gs",
			program: `
a = 0
a+=1
a-=2
a*=3
a%=4
print a
		`,
		},
		{
			name: "selfIncrDecr.gs",
			program: `
a = 0
a++
a--
print a
		`,
		},
		{
			name: "safeAccess.gs",
			program: `
type Student struct {
x,y,z
}

s = new Student{}
a = s.x?.z
print a
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGsCompiler()
			_, err := p.Compile(antlr.NewInputStream(tt.program))
			if err != nil {
				t.Fatal(err)
			}
			t.Log(p.Dump())
		})
	}
}
