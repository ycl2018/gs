package compile

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/consts"
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
		s.Write(consts.InstrGStore, l.Address)
	} else {
		s.Write(consts.InstrStore, l.Address)
	}
}

func (s *StackCompileVisitor) EmitLoad(v *VariableSymbol) {
	if v.Scope().GetName() == GlobalScopeName {
		s.Write(consts.InstrGLoad, v.Address)
	} else {
		s.Write(consts.InstrLoad, v.Address)
	}
}

func (s *StackCompileVisitor) VisitConstNode(v *consts.ConstNode) interface{} {
	switch v.Kind {
	case consts.ConstNodeKindInt:
		s.Write(consts.InstrIConst, v.Value.(int))
	case consts.ConstNodeKindFloat:
		s.Write(consts.InstrFConst, defineFloatConst(consts.ToFloatValue(v), s.GlobalScope).GetAddress())
	case consts.ConstNodeKindList:
		s.Write(consts.InstrSliceConst, defineSliceConst(v.Value.(*consts.SliceLiteralConst), s.GlobalScope).GetAddress())
	case consts.ConstNodeKindMap:
		s.Write(consts.InstrMapConst, defineMapConst(v.Value.(*consts.MapLiteralConst), s.GlobalScope).GetAddress())
	case consts.ConstNodeKindString:
		s.Write(consts.InstrSConst, s.defineStringConst(v.Value.(string)))
	case consts.ConstNodeKindBool:
		if v.Value.(bool) {
			s.Write(consts.InstrTrue)
		} else {
			s.Write(consts.InstrFalse)
		}
	default:
		panic(fmt.Sprintf("unknown constant type:%d", v.Kind))
	}
	return nil
}

func (s *StackCompileVisitor) defineStringConst(val string) int {
	return defineStringConst(val, s.GlobalScope).GetAddress()
}

func (s *StackCompileVisitor) Write(code consts.Instr, operands ...int) {
	operand := -11111
	if len(operands) > 0 {
		operand = operands[0]
	}
	s.CurFunc.Code = append(s.CurFunc.Code, consts.NewStackInstr(code, operand))
}

func (s *StackCompileVisitor) FillTarget(instrs ...*consts.StackInstr) {
	for _, instr := range instrs {
		instr.Operands = len(s.CurFunc.Code)
	}
}

func (s *StackCompileVisitor) storeQid(qid gen.IQidContext) {
	var ids []string
	var query []int // tokenType
	var qExprs []*gen.ExprContext
	var j int // index of qExprs
	for i, child := range qid.GetChildren() {
		if i == 0 {
			if env := child.(*gen.PrimaryContext).ENV(); env != nil {
				if s.Env != nil {
					s.storeQidToEnv(qid)
					return
				} else {
					ids = append(ids, "$")
					continue
				}
			} else {
				ids = append(ids, child.(*gen.PrimaryContext).ID().GetText())
			}
			continue
		}
		if node, ok := child.(antlr.TerminalNode); ok {
			switch t := node.GetSymbol().GetTokenType(); t {
			case gen.GsLexerDOT, gen.GsLexerLBRACK:
				query = append(query, t)
			case gen.GsLexerSAFE_DOT, gen.GsLexerSAFE_LBRACK:
				s.Log.ErrorToken(node.GetSymbol(), "syntax error:can't use %s in assign left side", node.GetSymbol().GetText())
				return
			case gen.GsLexerID:
				ids = append(ids, node.GetText())
			}
		} else if e, ok := child.(*gen.ExprContext); ok {
			qExprs = append(qExprs, e)
		}
	}
	if ids[0] == "$" {
		s.Write(consts.InstrLoadEnv)
	} else {
		primarySymbol := s.Scopes[qid].Resolve(ids[0]).(*VariableSymbol)
		if len(ids) == 1 {
			// no need load primarySymbol
			s.EmitStore(primarySymbol)
			return
		}
		s.EmitLoad(primarySymbol)
	}
	for i := 0; i < len(query)-1; i++ {
		switch query[i] {
		case gen.GsLexerDOT:
			// fieldLoad
			fieldName := ids[i+1]
			operand := s.defineStringConst(fieldName)
			s.Write(consts.InstrFLoad, operand)
		case gen.GsLexerLBRACK:
			// arrayLoad/mapLoad
			expr := qExprs[j]
			j++
			expr.Accept(s)
			s.Write(consts.InstrIndexLoad)
		}
	}
	// lastQuery
	switch query[len(query)-1] {
	case gen.GsLexerDOT:
		fieldName := ids[len(ids)-1]
		operand := s.defineStringConst(fieldName)
		s.Write(consts.InstrFStore, operand)
	case gen.GsLexerLBRACK:
		// arrayStore/mapStore
		expr := qExprs[j]
		expr.Accept(s)
		s.Write(consts.InstrIndexStore)
	}
}

func (s *StackCompileVisitor) loadQid(qid gen.IQidContext) {
	var ids []string
	var query []int // tokenType
	var qExprs []*gen.ExprContext
	var j int // index of qExprs
	for i, child := range qid.GetChildren() {
		if i == 0 {
			if env := child.(*gen.PrimaryContext).ENV(); env != nil {
				if s.Env != nil {
					s.loadQidFromEnv(qid)
					return
				} else {
					ids = append(ids, "$")
					continue
				}
			} else {
				ids = append(ids, child.(*gen.PrimaryContext).ID().GetText())
			}
			continue
		}
		if node, ok := child.(antlr.TerminalNode); ok {
			switch t := node.GetSymbol().GetTokenType(); t {
			case gen.GsLexerDOT, gen.GsLexerSAFE_DOT, gen.GsLexerLBRACK, gen.GsLexerSAFE_LBRACK:
				query = append(query, t)
			case gen.GsLexerID:
				ids = append(ids, node.GetText())
			}
		} else if e, ok := child.(*gen.ExprContext); ok {
			qExprs = append(qExprs, e)
		}
	}
	if ids[0] == "$" {
		if len(query) == 0 {
			s.Write(consts.InstrLoadEnv)
			return
		}
	} else {
		if s.Scopes[qid] == nil {
			panic("scope not found")
		}
		primarySymbol, ok := s.Scopes[qid].Resolve(ids[0]).(*VariableSymbol)
		if !ok {
			s.Log.ErrorToken(qid.GetStart(), "undefined symbol: %s", ids[0])
			return
		}
		s.EmitLoad(primarySymbol)
	}
	var brNils []*consts.StackInstr
	for i := 0; i < len(query); i++ {
		switch query[i] {
		case gen.GsLexerSAFE_DOT:
			// if true, then fieldLoad
			brNil := consts.NewStackInstr(consts.InstrBRNil, placeholder)
			brNils = append(brNils, brNil)
			s.WriteInstr(brNil)
			fallthrough
		case gen.GsLexerDOT:
			// fieldLoad
			fieldName := ids[i+1]
			operand := s.defineStringConst(fieldName)
			s.Write(consts.InstrFLoad, operand)
		case gen.GsLexerSAFE_LBRACK:
			brNil := consts.NewStackInstr(consts.InstrBRNil, placeholder)
			brNils = append(brNils, brNil)
			s.WriteInstr(brNil)
			fallthrough
		case gen.GsLexerLBRACK:
			// arrayLoad/mapLoad
			expr := qExprs[j]
			j++
			expr.Accept(s)
			s.Write(consts.InstrIndexLoad)
		}
	}
	s.FillTarget(brNils...)
}

func (s *StackCompileVisitor) Code() vm.Code {
	var constPoll []consts.Const
	var cs []*ConstSymbol
	for _, c := range s.GlobalScope.Consts {
		cs = append(cs, c)
	}
	// sort
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Address < cs[j].Address
	})
	// fill const poll
	toFuncConst := func(f *FunctionSymbol, addr int) consts.FunctionConst {
		var codes = make([]consts.StackInstr, len(f.Code))
		for i, instr := range f.Code {
			codes[i] = *instr
		}
		return consts.FunctionConst{
			Name:       f.Name,
			ParamCount: len(f.FormalArgs),
			LocalCount: f.LocalNums(),
			Addr:       addr,
			Code:       codes,
		}
	}
	for i, c := range cs {
		if c.Kind == consts.ConstFunc {
			// fill code
			name := c.Name[strings.Index(c.Name, "::")+2:]
			c.Value = toFuncConst(s.GlobalScope.Resolve(name).(*FunctionSymbol), i)
		}
		constPoll = append(constPoll, consts.Const{
			Value: c.Value,
			Kind:  c.Kind,
		})
	}

	var envType reflect.Type
	if s.Env != nil {
		envType = s.Env.RType
	}
	return vm.Code{
		Globals:      int(s.GlobalScope.LocalVarAllocator),
		ConstPool:    constPoll,
		MainFunc:     toFuncConst(s.MainFunc, -1),
		BuildEnvType: envType,
	}
}
