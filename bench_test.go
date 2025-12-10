package gs

import (
	"io"
	"testing"
)

//goos: darwin
//goarch: arm64
//pkg: github.com/ycl2018/gs
//cpu: Apple M3 Pro
//BenchmarkWithCache
//BenchmarkWithCache-12    	  407086	      3042 ns/op
//BenchmarkNoCache
//BenchmarkNoCache-12      	  263005	      4589 ns/op
//PASS

var program = `
for i = range 10 {
	println($.Itf.Hello() == "Hello World")
}
`

func BenchmarkWithCache(b *testing.B) {
	code, err := Compile(program)
	if err != nil {
		b.Fatal(err)
		return
	}
	env := &MyEnv{
		Itf: &MyItf{
			Name: "World",
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Run(code, env, Output(io.Discard))
		if err != nil {
			b.Fatal(err)
			return
		}
	}
}

func BenchmarkNoCache(b *testing.B) {
	code, err := Compile(program)
	if err != nil {
		b.Fatal(err)
		return
	}
	env := &MyEnv{
		Itf: &MyItf{
			Name: "World",
		},
	}
	noCacheOp := NoCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Run(code, env, noCacheOp, Output(io.Discard))
		if err != nil {
			b.Fatal(err)
			return
		}
	}
}
