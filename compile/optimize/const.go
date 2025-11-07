package optimize

import (
	"github.com/ycl2018/gs/gen"
)

var _ gen.GsVisitor = (*constOptimizer)(nil)

type constOptimizer struct {
	gen.BaseVisitor
	FoldConstExpr bool
}

func (c constOptimizer) VisitProgram(ctx *gen.ProgramContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitStructDefinition(ctx *gen.StructDefinitionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitFunctionDefinition(ctx *gen.FunctionDefinitionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitBlock(ctx *gen.BlockContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitEmptyStmt(ctx *gen.EmptyStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitStructStmt(ctx *gen.StructStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitAssignStmt(ctx *gen.AssignStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitSelfOpAssignStmt(ctx *gen.SelfOpAssignStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitIncrDecrStmt(ctx *gen.IncrDecrStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitReturnStmt(ctx *gen.ReturnStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitPrintStmt(ctx *gen.PrintStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitPrintfStmt(ctx *gen.PrintfStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitIfStmt(ctx *gen.IfStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitForCStyleStmt(ctx *gen.ForCStyleStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitForRangeStmt(ctx *gen.ForRangeStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitForCondStmt(ctx *gen.ForCondStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitCallStmt(ctx *gen.CallStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitBreakStmt(ctx *gen.BreakStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitContinueStmt(ctx *gen.ContinueStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitAssign(ctx *gen.AssignContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitIncrDecr(ctx *gen.IncrDecrContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitSingleIter(ctx *gen.SingleIterContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitDoubleIter(ctx *gen.DoubleIterContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitForInit(ctx *gen.ForInitContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitForUpdate(ctx *gen.ForUpdateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitSelfAssign(ctx *gen.SelfAssignContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitSelfUpdate(ctx *gen.SelfUpdateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitIncrDecrUpdate(ctx *gen.IncrDecrUpdateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitAssignUpdate(ctx *gen.AssignUpdateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitSelfAssignOp(ctx *gen.SelfAssignOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitCall(ctx *gen.CallContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitExpr(ctx *gen.ExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitLogicalOrExpr(ctx *gen.LogicalOrExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitLogicalAndExpr(ctx *gen.LogicalAndExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitComparisonExpr(ctx *gen.ComparisonExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitAddExpr(ctx *gen.AddExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitBinExpr(ctx *gen.BinExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitMulExpr(ctx *gen.MulExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitPowExpr(ctx *gen.PowExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitNegAtom(ctx *gen.NegAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitIntAtom(ctx *gen.IntAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitFloatAtom(ctx *gen.FloatAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitCharAtom(ctx *gen.CharAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitStringAtom(ctx *gen.StringAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitTrueAtom(ctx *gen.TrueAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitFalseAtom(ctx *gen.FalseAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitNilAtom(ctx *gen.NilAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitNotAtom(ctx *gen.NotAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitQidAtom(ctx *gen.QidAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitCallAtom(ctx *gen.CallAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitInstanceAtom(ctx *gen.InstanceAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitArrayAtom(ctx *gen.ArrayAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitIndexAccessAtom(ctx *gen.IndexAccessAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitDictAtom(ctx *gen.DictAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitParenAtom(ctx *gen.ParenAtomContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitArrayLiteral(ctx *gen.ArrayLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitIndexAccess(ctx *gen.IndexAccessContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitSliceExpr(ctx *gen.SliceExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitDictLiteral(ctx *gen.DictLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitStrKeyEntry(ctx *gen.StrKeyEntryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitIdKeyEntry(ctx *gen.IdKeyEntryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitInstance(ctx *gen.InstanceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitQid(ctx *gen.QidContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitPrimary(ctx *gen.PrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitCompOp(ctx *gen.CompOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitAddOp(ctx *gen.AddOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitBitOp(ctx *gen.BitOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitMulOp(ctx *gen.MulOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c constOptimizer) VisitPowOp(ctx *gen.PowOpContext) interface{} {
	//TODO implement me
	panic("implement me")
}
