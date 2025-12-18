package gs

import (
	"testing"

	"github.com/expr-lang/expr"
)

type TestEnv struct {
	Str    string
	Int    int64
	Bool   bool
	Slice  []string
	Map    map[string]string
	Struct *TestEnv
}

func BenchmarkEvalGS(b *testing.B) {
	program := `
	return $.Str == "string" && $.Int == 100 && $.Bool == true && in($.Slice, "item") && $.Map["key"] == "value"
	`
	env := &TestEnv{
		Str:   "string",
		Int:   100,
		Bool:  true,
		Slice: []string{"a", "foo", "bar", "item"},
		Map:   map[string]string{"key": "value"},
	}
	b.ResetTimer()
	for b.Loop() {
		ret, err := Eval(program, env)
		if err != nil {
			b.Fatal(err)
		}
		if !ret.MustBool() {
			b.Fatal("eval failed")
		}
	}
}

func BenchmarkEvalExpr(b *testing.B) {
	program := `
Str == "string" && Int == 100 && Bool == true && "item" in Slice && Map["key"] == "value"

	`
	env := &TestEnv{
		Str:   "string",
		Int:   100,
		Bool:  true,
		Slice: []string{"item"},
		Map:   map[string]string{"key": "value"},
	}
	b.ResetTimer()
	for b.Loop() {
		ret, err := expr.Eval(program, env)
		if err != nil {
			b.Fatal(err)
		}
		if !ret.(bool) {
			b.Fatal("eval failed")
		}
	}
}

func BenchmarkCompileMapExpr(b *testing.B) {
	program := `
Str == "string" && Int == 100 && Bool == true  && Map["key"] == "value"
	`
	env := map[string]any{
		"Str":   "string",
		"Int":   int64(100),
		"Bool":  true,
		"Slice": []string{"item"},
		"Map":   map[string]string{"key": "value"},
	}
	code, err := expr.Compile(program, expr.Env(env), expr.AsBool())
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.StartTimer()
	for b.Loop() {
		_, err := expr.Run(code, env)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompileMapGS(b *testing.B) {
	program := `
	return $["Str"] == "string" && $["Int"] == 100 && $["Bool"] == true  && $["Map"]["key"] == "value"
	`
	env := map[string]any{
		"Str":   "string",
		"Int":   int64(100),
		"Bool":  true,
		"Slice": []string{"item"},
		"Map":   map[string]string{"key": "value"},
	}
	code, err := Compile(program, Env(env))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.StartTimer()
	for b.Loop() {
		_, err := Run(code, env)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompileStructExpr(b *testing.B) {
	program := `
Str == "string" && Int + 100 == 200 && Bool == true && "item" in Slice && Map["key"] == "value" && Struct.Str == "string" && Struct.Int == 100 && Struct.Bool == true 
	`
	env := &TestEnv{
		Str:   "string",
		Int:   100,
		Bool:  true,
		Slice: []string{"foo", "bar", "item", "222"},
		Map:   map[string]string{"key": "value"},
		Struct: &TestEnv{
			Str:   "string",
			Int:   100,
			Bool:  true,
			Slice: []string{"foo", "bar", "item", "222"},
			Map:   map[string]string{"key": "value"},
		},
	}
	code, err := expr.Compile(program, expr.Env(env), expr.AsBool())
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.StartTimer()
	for b.Loop() {
		_, err := expr.Run(code, env)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompileStructGS(b *testing.B) {
	program := `
	return $.Str == "string" && $.Int + 100 == 200 && $.Bool == true  && $.Map["key"] == "value" && $.Struct.Str == "string" && $.Struct.Int == 100 && $.Struct.Bool == true 
	`
	env := &TestEnv{
		Str:   "string",
		Int:   100,
		Bool:  true,
		Slice: []string{"foo", "bar", "item", "222"},
		Map:   map[string]string{"key": "value"},
		Struct: &TestEnv{
			Str:   "string",
			Int:   100,
			Bool:  true,
			Slice: []string{"foo", "bar", "item", "222"},
			Map:   map[string]string{"key": "value"},
		},
	}
	code, err := Compile(program, Env(env))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.StartTimer()
	for b.Loop() {
		_, err := Run(code, env)
		if err != nil {
			b.Fatal(err)
		}
	}
}
