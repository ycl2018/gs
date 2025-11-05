package compile

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/gen"
	"github.com/ycl2018/gs/vm"
)

const placeholder = -1

func (s *StackCompileVisitor) PushLoopStack(loop *ForLoop) {
	s.LoopStack = append(s.LoopStack, loop)
}

func (s *StackCompileVisitor) PopLoopStack() *ForLoop {
	loop := s.LoopStack[len(s.LoopStack)-1]
	s.LoopStack = s.LoopStack[:len(s.LoopStack)-1]
	return loop
}

func (s *StackCompileVisitor) GetCurrentLoop() *ForLoop {
	if len(s.LoopStack) == 0 {
		return nil
	}
	return s.LoopStack[len(s.LoopStack)-1]
}

func (s *StackCompileVisitor) EmitStore(l *VariableSymbol) {
	if l.Scope().GetName() == GlobalScopeName {
		s.Write(vm.InstrGStore, l.Address)
	} else {
		s.Write(vm.InstrStore, l.Address)
	}
}

func (s *StackCompileVisitor) EmitLoad(v *VariableSymbol) {
	if v.Scope().GetName() == GlobalScopeName {
		s.Write(vm.InstrGLoad, v.Address)
	} else {
		s.Write(vm.InstrLoad, v.Address)
	}
}

func (s *StackCompileVisitor) defineStringConst(val string) int {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%s", vm.ConstString, val),
		Kind:  vm.ConstString,
		Value: val,
	}
	cSymbol, _ := s.GlobalScope.DefineOrGetConst(constSymbol)
	return cSymbol.GetAddress()
}

func (s *StackCompileVisitor) Write(code int, operands ...int) {
	s.CurFunc.Code = append(s.CurFunc.Code, vm.NewStackInstr(code, operands...))
}

func (s *StackCompileVisitor) FillTarget(instrs ...*vm.StackInstr) {
	for _, instr := range instrs {
		instr.Operands[0] = len(s.CurFunc.Code)
	}
}

func (s *StackCompileVisitor) storeQid(qid gen.IQidContext) {
	// insert check nil br to pop left value
	brNil := vm.NewStackInstr(vm.InstrBRNil, placeholder)
	s.WriteInstr(brNil)
	var ids []string
	var query []int // tokenType
	var qExprs []*gen.ExprContext
	for i, child := range qid.GetChildren() {
		if i == 0 {
			ids = append(ids, child.(*gen.PrimaryContext).ID().GetText())
			continue
		}
		if node, ok := child.(antlr.TerminalNode); ok {
			switch t := node.GetSymbol().GetTokenType(); t {
			case gen.GsLexerDOT, gen.GsLexerSAFE_DOT, gen.GsLexerLBRACK, gen.GsLexerSAFE_LBRACK:
				{
					query = append(query, t)
				}
			case gen.GsLexerID:
				ids = append(ids, node.GetText())
			}
		} else if e, ok := child.(*gen.ExprContext); ok {
			qExprs = append(qExprs, e)
		}
	}
	primarySymbol := s.Scopes[qid].Resolve(ids[0]).(*VariableSymbol)
	if len(ids) == 1 {
		// no need load primarySymbol
		s.EmitStore(primarySymbol)
		brOK := vm.NewStackInstr(vm.InstrBR, placeholder)
		s.WriteInstr(brOK)
		s.FillTarget(brNil)
		s.Write(vm.InstrPop, 1)
		s.FillTarget(brOK)
		return
	}
	s.EmitLoad(primarySymbol)
	var brNils []*vm.StackInstr
	var brOk []*vm.StackInstr
	for i := 0; i < len(query)-1; i++ {
		switch query[i] {
		case gen.GsLexerSAFE_DOT:
			// if true, then fieldLoad
			brNils = append(brNils, vm.NewStackInstr(vm.InstrBRNil, placeholder))
			fallthrough
		case gen.GsLexerDOT:
			// fieldLoad
			fieldName := ids[i+1]
			operand := s.defineStringConst(fieldName)
			s.Write(vm.InstrFLoad, operand)
		case gen.GsLexerSAFE_LBRACK:
			brNils = append(brNils, vm.NewStackInstr(vm.InstrBRNil, placeholder))
			fallthrough
		case gen.GsLexerLBRACK:
			// arrayLoad/mapLoad
			expr := qExprs[i]
			expr.Accept(s)
			s.Write(vm.InstrIndexAccess)
		}
	}
	// lastQuery
	switch query[len(query)-1] {
	case gen.GsLexerSAFE_DOT:
		brNils = append(brNils, vm.NewStackInstr(vm.InstrBRNil, placeholder))
	case gen.GsLexerDOT:
		fieldName := ids[len(ids)-1]
		operand := s.defineStringConst(fieldName)
		s.Write(vm.InstrFStore, operand)
		brOK := vm.NewStackInstr(vm.InstrBR, placeholder)
		s.WriteInstr(brOK)
		brOk = append(brOk, brOK)
	case gen.GsLexerSAFE_LBRACK:
		brNils = append(brNils, vm.NewStackInstr(vm.InstrBRNil, placeholder))
		fallthrough
	case gen.GsLexerLBRACK:
		// arrayStore/mapStore
		expr := qExprs[len(qExprs)-1]
		expr.Accept(s)
		s.Write(vm.InstrIndexStore)
		brOK := vm.NewStackInstr(vm.InstrBR, placeholder)
		s.WriteInstr(brOK)
		brOk = append(brOk, brOK)
	}
	for _, nilInstr := range brNils {
		s.FillTarget(nilInstr)
	}
	s.Write(vm.InstrPop, 1) // pop right value
	s.FillTarget(brNil)
	s.Write(vm.InstrPop, 1) // pop left value
	s.FillTarget(brOk...)
}

func (s *StackCompileVisitor) loadQid(qid gen.IQidContext) {
	var ids []string
	var query []int // tokenType
	var qExprs []*gen.ExprContext
	for i, child := range qid.GetChildren() {
		if i == 0 {
			ids = append(ids, child.(*gen.PrimaryContext).ID().GetText())
			continue
		}
		if node, ok := child.(antlr.TerminalNode); ok {
			switch t := node.GetSymbol().GetTokenType(); t {
			case gen.GsLexerDOT, gen.GsLexerSAFE_DOT, gen.GsLexerLBRACK, gen.GsLexerSAFE_LBRACK:
				{
					query = append(query, t)
				}
			case gen.GsLexerID:
				ids = append(ids, node.GetText())
			}
		} else if e, ok := child.(*gen.ExprContext); ok {
			qExprs = append(qExprs, e)
		}
	}
	primarySymbol, ok := s.Scopes[qid].Resolve(ids[0]).(*VariableSymbol)
	if !ok {
		s.Log.ErrorToken(qid.GetStart(), "undefined symbol: %s", ids[0])
		return
	}
	if len(ids) == 0 {
		// no need load primarySymbol
		s.EmitStore(primarySymbol)
	}
	s.EmitLoad(primarySymbol)
	var brNils []*vm.StackInstr
	for i := 0; i < len(query); i++ {
		switch query[i] {
		case gen.GsLexerSAFE_DOT:
			// if true, then fieldLoad
			brNils = append(brNils, vm.NewStackInstr(vm.InstrBRNil, placeholder))
			fallthrough
		case gen.GsLexerDOT:
			// fieldLoad
			fieldName := ids[i+1]
			operand := s.defineStringConst(fieldName)
			s.Write(vm.InstrFLoad, operand)
		case gen.GsLexerSAFE_LBRACK:
			brNils = append(brNils, vm.NewStackInstr(vm.InstrBRNil, placeholder))
			fallthrough
		case gen.GsLexerLBRACK:
			// arrayLoad/mapLoad
			expr := qExprs[i]
			expr.Accept(s)
			s.Write(vm.InstrIndexAccess)
		}
	}
	s.FillTarget(brNils...)
}
