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

func (s *StackCompileVisitor) EmitStore(l *VariableSymbol, token antlr.Token) {
	if l.Scope().GetName() == GlobalScopeName {
		s.Write(consts.InstrGStore, token, l.Address)
	} else {
		s.Write(consts.InstrStore, token, l.Address)
	}
}

func (s *StackCompileVisitor) EmitLoad(v *VariableSymbol, token antlr.Token) {
	if v.Scope().GetName() == GlobalScopeName {
		s.Write(consts.InstrGLoad, token, v.Address)
	} else {
		s.Write(consts.InstrLoad, token, v.Address)
	}
}

func (s *StackCompileVisitor) VisitConstNode(v *consts.ConstNode) interface{} {
	switch v.Kind {
	case consts.ConstNodeKindInt:
		s.Write(consts.InstrIConst, v.GetStart(), v.Value.(int))
	case consts.ConstNodeKindFloat:
		s.Write(consts.InstrFConst, v.GetStart(), defineFloatConst(consts.ToFloatValue(v), s.GlobalScope).GetAddress())
	case consts.ConstNodeKindList:
		s.Write(consts.InstrSliceConst, v.GetStart(), defineSliceConst(v.Value.(*consts.SliceLiteralConst), s.GlobalScope).GetAddress())
	case consts.ConstNodeKindMap:
		s.Write(consts.InstrMapConst, v.GetStart(), defineMapConst(v.Value.(*consts.MapLiteralConst), s.GlobalScope).GetAddress())
	case consts.ConstNodeKindString:
		s.Write(consts.InstrSConst, v.GetStart(), s.defineStringConst(v.Value.(string)))
	case consts.ConstNodeKindBool:
		if v.Value.(bool) {
			s.Write(consts.InstrTrue, v.GetStart())
		} else {
			s.Write(consts.InstrFalse, v.GetStart())
		}
	default:
		panic(fmt.Sprintf("unknown constant type:%d", v.Kind))
	}
	return nil
}

func (s *StackCompileVisitor) defineStringConst(val string) int {
	return defineStringConst(val, s.GlobalScope).GetAddress()
}

func (s *StackCompileVisitor) Write(code consts.Instr, token antlr.Token, operands ...int) {
	operand := -11111
	if len(operands) > 0 {
		operand = operands[0]
	}
	s.CurFunc.Code = append(s.CurFunc.Code, consts.NewStackInstr(code, operand))
	s.CurFunc.Debugger.Table = append(s.CurFunc.Debugger.Table, consts.Info{
		Line: token.GetLine(),
	})
}

func (s *StackCompileVisitor) FillTarget(instrs ...*consts.StackInstr) {
	for _, instr := range instrs {
		instr.Operands = len(s.CurFunc.Code)
	}
}

func (s *StackCompileVisitor) storeLvalue(lvalue *gen.LvalueContext) {
	// lvalue: qid | * lvalue;
	if len(lvalue.GetChildren()) == 2 {
		s.Write(consts.InstrNewPtrValue, lvalue.GetStart())
		s.storeLvalue(lvalue.Lvalue().(*gen.LvalueContext))
		return
	}
	qid := lvalue.Qid()
	s.storeQid(qid)
}

func (s *StackCompileVisitor) storeQid(qid gen.IQidContext) {
	var primaryText string
	primary, accessors := qid.Primary().(*gen.PrimaryContext), qid.AllAccessor()
	if env := primary.ENV(); env != nil {
		if s.Env != nil {
			s.storeQidToEnv(qid)
			return
		} else {
			primaryText = env.GetText()
		}
	} else {
		primaryText = primary.ID().GetText()
	}
	// qid: primary accessor*
	if primaryText == "$" {
		s.Write(consts.InstrLoadEnv, qid.GetStart())
	} else {
		primarySymbol := s.CurScope.Resolve(primaryText).(*VariableSymbol)
		if len(accessors) == 0 {
			// no need load primarySymbol
			s.EmitStore(primarySymbol, qid.GetStart())
			return
		}
		s.EmitLoad(primarySymbol, qid.GetStart())
	}
	for i, accessor := range accessors {
		switch a := accessor.(type) {
		case *gen.PropertyAccessContext:
			// fieldLoad
			// checkDOT
			if a.SAFE_DOT() != nil {
				s.Log.ErrorToken(a.SAFE_DOT().GetSymbol(), "syntax error:can't use ? in assign left side")
				return
			}
			fieldName := a.ID().GetText()
			operand := s.defineStringConst(fieldName)
			if i < len(accessors)-1 {
				s.Write(consts.InstrFLoad, a.GetStart(), operand)
			} else {
				s.Write(consts.InstrFStore, a.GetStart(), operand)
			}
		case *gen.IndexAccessContext:
			// checkLBRACK
			if a.SAFE_LBRACK() != nil {
				s.Log.ErrorToken(a.SAFE_LBRACK().GetSymbol(), "syntax error:can't use ? in assign left side")
				return
			}
			switch t := a.GetChild(1).(type) {
			case *gen.ExprContext:
				t.Accept(s)
				if i < len(accessors)-1 {
					s.Write(consts.InstrIndexLoad, a.GetStart())
				} else {
					s.Write(consts.InstrIndexStore, a.GetStart())
				}
			case *gen.SliceExprContext:
				if i == len(accessors)-1 {
					s.Log.ErrorToken(a.GetStart(), "syntax error:can't assign to slice split")
					return
				}
				t.Accept(s)
			}
		}
	}
}

func (s *StackCompileVisitor) loadLvalue(lvalue *gen.LvalueContext) {
	// lvalue: qid | * lvalue;
	if len(lvalue.GetChildren()) == 2 {
		s.loadLvalue(lvalue.Lvalue().(*gen.LvalueContext))
		s.Write(consts.InstrDeref, lvalue.GetStart())
		return
	}
	qid := lvalue.Qid()
	s.loadQid(qid)
}

func (s *StackCompileVisitor) loadQid(qid gen.IQidContext) {
	var primaryText string
	primary, accessors := qid.Primary().(*gen.PrimaryContext), qid.AllAccessor()
	if env := primary.ENV(); env != nil {
		if s.Env != nil {
			s.loadQidFromEnv(qid)
			return
		} else {
			primaryText = env.GetText()
		}
	} else {
		primaryText = primary.ID().GetText()
	}
	// qid: primary accessor*
	if primaryText == "$" {
		s.Write(consts.InstrLoadEnv, qid.GetStart())
	} else {
		scope := s.CurScope
		primarySymbol, ok := scope.Resolve(primaryText).(*VariableSymbol)
		if !ok {
			s.Log.ErrorToken(qid.GetStart(), "undefined symbol: %s", primaryText)
			return
		}
		s.EmitLoad(primarySymbol, qid.GetStart())
	}
	var brNils []*consts.StackInstr
	for _, accessor := range accessors {
		switch a := accessor.(type) {
		case *gen.PropertyAccessContext:
			// fieldLoad
			// checkDOT
			if a.SAFE_DOT() != nil {
				// if true, then fieldLoad
				brNil := consts.NewStackInstr(consts.InstrBRNil, placeholder)
				brNils = append(brNils, brNil)
				s.WriteInstr(brNil, a.GetStart())
			}
			fieldName := a.ID().GetText()
			operand := s.defineStringConst(fieldName)
			s.Write(consts.InstrFLoad, a.GetStart(), operand)
		case *gen.IndexAccessContext:
			// checkLBRACK
			if a.SAFE_LBRACK() != nil {
				brNil := consts.NewStackInstr(consts.InstrBRNil, placeholder)
				brNils = append(brNils, brNil)
				s.WriteInstr(brNil, a.GetStart())
			}
			switch t := a.GetChild(1).(type) {
			case *gen.ExprContext:
				t.Accept(s)
				s.Write(consts.InstrIndexLoad, a.GetStart())
			case *gen.SliceExprContext:
				t.Accept(s)
			}
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
			Debugger:   f.Debugger,
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
