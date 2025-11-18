package compile

import (
	"fmt"
	"io"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type InterpreterListener interface {
	Infof(string, ...any)
	Errorf(string, ...any)
	ErrorToken(antlr.Token, string, ...any)
}

var _ InterpreterListener = (*DefaultInterpreterListener)(nil)

type DefaultInterpreterListener struct {
	Writer      io.Writer
	ErrorWriter io.Writer
}

func (d DefaultInterpreterListener) Infof(s string, a ...any) {
	_, _ = fmt.Fprintf(d.Writer, s+"\n", a...)
}

func (d DefaultInterpreterListener) Errorf(s string, a ...any) {
	_, _ = fmt.Fprintf(d.ErrorWriter, s+"\n", a...)
}

func (d DefaultInterpreterListener) ErrorToken(token antlr.Token, s string, a ...any) {
	var sb strings.Builder
	lineStr := fmt.Sprintf("\n<line %d> %s: ", token.GetLine(), token.GetText())
	sb.WriteString(lineStr)
	sb.WriteString(fmt.Sprintf(s, a...))
	str := sb.String()
	_, _ = fmt.Fprintf(d.ErrorWriter, str+"\n")
	start, end := strings.Index(str, ">"), strings.Index(str, ":")
	fmt.Fprintf(d.ErrorWriter, strings.Repeat(" ", start+1)+strings.Repeat("^", end-start-2)+"\n")
}
