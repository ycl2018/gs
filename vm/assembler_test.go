package vm

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs-go/interpreter/asm"
	"github.com/ycl2018/gs-go/interpreter/asm/gen"
)

func TestAssembler(t *testing.T) {
	codeAssembler := asm.NewByteCodeAssembler(Instructions)
	asmCode := `
.def fact: args=1, locals=0
;	if n < 2 return 1
	load 0
	iconst 2
	lt
	brf cont
	iconst 1
	ret
cont:
;	return n * fact(n-1)
	load 0
	load 0
	iconst 1
	sub
	call fact()
	mul
	ret

.def main: args=0, locals=0
; print fact(10)
	iconst 10
	call fact()
	print
	halt
`
	lexer := gen.NewAssemblerLexer(antlr.NewInputStream(asmCode))
	parser := gen.NewAssemblerParser(antlr.NewCommonTokenStream(lexer, 0))
	parser.AsmGenerator = codeAssembler
	parser.Program()

	if parser.GetError() != nil {
		t.Errorf("parser err:%v", parser.GetError())
		return
	}
	// disAssembler
	disAssembler := asm.NewDisAssembler(Instructions, codeAssembler.Code, codeAssembler.ConstPool, codeAssembler.FuncConsPool)
	assemble, err := disAssembler.DisAssemble()
	if err != nil {
		t.Errorf("diaAssemble err:%v", err)
		return
	}
	t.Log(assemble)
}
