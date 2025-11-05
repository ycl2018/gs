// Code generated from github.com/ycl2018/gs/Gs.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // Gs
import "github.com/antlr4-go/antlr/v4"


type BaseGsVisitor struct {
	antlr.ParseTreeVisitor
}

func (v *BaseGsVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitStructDefinition(ctx *StructDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitEmptyStmt(ctx *EmptyStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitStructStmt(ctx *StructStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitSelfOpAssignStmt(ctx *SelfOpAssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitIncrDecrStmt(ctx *IncrDecrStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitPrintfStmt(ctx *PrintfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitForCStyleStmt(ctx *ForCStyleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitForRangeStmt(ctx *ForRangeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitForCondStmt(ctx *ForCondStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitCallStmt(ctx *CallStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitIncrDecr(ctx *IncrDecrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitSingleIter(ctx *SingleIterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitDoubleIter(ctx *DoubleIterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitForUpdate(ctx *ForUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitSelfAssign(ctx *SelfAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitSelfUpdate(ctx *SelfUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitIncrDecrUpdate(ctx *IncrDecrUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitAssignUpdate(ctx *AssignUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitSelfAssignOp(ctx *SelfAssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitLogicalOrExpr(ctx *LogicalOrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitLogicalAndExpr(ctx *LogicalAndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitComparisonExpr(ctx *ComparisonExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitBinExpr(ctx *BinExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitPowExpr(ctx *PowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitNegAtom(ctx *NegAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitIntAtom(ctx *IntAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitFloatAtom(ctx *FloatAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitCharAtom(ctx *CharAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitStringAtom(ctx *StringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitTrueAtom(ctx *TrueAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitFalseAtom(ctx *FalseAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitNilAtom(ctx *NilAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitNotAtom(ctx *NotAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitQidAtom(ctx *QidAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitCallAtom(ctx *CallAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitInstanceAtom(ctx *InstanceAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitArrayAtom(ctx *ArrayAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitIndexAccessAtom(ctx *IndexAccessAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitDictAtom(ctx *DictAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitParenAtom(ctx *ParenAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitIndexAccess(ctx *IndexAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitSliceExpr(ctx *SliceExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitDictLiteral(ctx *DictLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitStrKeyEntry(ctx *StrKeyEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitIdKeyEntry(ctx *IdKeyEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitInstance(ctx *InstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitQid(ctx *QidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitCompOp(ctx *CompOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitAddOp(ctx *AddOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitBitOp(ctx *BitOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitMulOp(ctx *MulOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGsVisitor) VisitPowOp(ctx *PowOpContext) interface{} {
	return v.VisitChildren(ctx)
}
