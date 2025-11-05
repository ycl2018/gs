// Code generated from github.com/ycl2018/gs/Gs.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // Gs
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by GsParser.
type GsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GsParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GsParser#structDefinition.
	VisitStructDefinition(ctx *StructDefinitionContext) interface{}

	// Visit a parse tree produced by GsParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by GsParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GsParser#emptyStmt.
	VisitEmptyStmt(ctx *EmptyStmtContext) interface{}

	// Visit a parse tree produced by GsParser#structStmt.
	VisitStructStmt(ctx *StructStmtContext) interface{}

	// Visit a parse tree produced by GsParser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by GsParser#selfOpAssignStmt.
	VisitSelfOpAssignStmt(ctx *SelfOpAssignStmtContext) interface{}

	// Visit a parse tree produced by GsParser#incrDecrStmt.
	VisitIncrDecrStmt(ctx *IncrDecrStmtContext) interface{}

	// Visit a parse tree produced by GsParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by GsParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by GsParser#printfStmt.
	VisitPrintfStmt(ctx *PrintfStmtContext) interface{}

	// Visit a parse tree produced by GsParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by GsParser#forCStyleStmt.
	VisitForCStyleStmt(ctx *ForCStyleStmtContext) interface{}

	// Visit a parse tree produced by GsParser#forRangeStmt.
	VisitForRangeStmt(ctx *ForRangeStmtContext) interface{}

	// Visit a parse tree produced by GsParser#forCondStmt.
	VisitForCondStmt(ctx *ForCondStmtContext) interface{}

	// Visit a parse tree produced by GsParser#callStmt.
	VisitCallStmt(ctx *CallStmtContext) interface{}

	// Visit a parse tree produced by GsParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by GsParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by GsParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by GsParser#incrDecr.
	VisitIncrDecr(ctx *IncrDecrContext) interface{}

	// Visit a parse tree produced by GsParser#singleIter.
	VisitSingleIter(ctx *SingleIterContext) interface{}

	// Visit a parse tree produced by GsParser#doubleIter.
	VisitDoubleIter(ctx *DoubleIterContext) interface{}

	// Visit a parse tree produced by GsParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by GsParser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by GsParser#selfAssign.
	VisitSelfAssign(ctx *SelfAssignContext) interface{}

	// Visit a parse tree produced by GsParser#selfUpdate.
	VisitSelfUpdate(ctx *SelfUpdateContext) interface{}

	// Visit a parse tree produced by GsParser#incrDecrUpdate.
	VisitIncrDecrUpdate(ctx *IncrDecrUpdateContext) interface{}

	// Visit a parse tree produced by GsParser#assignUpdate.
	VisitAssignUpdate(ctx *AssignUpdateContext) interface{}

	// Visit a parse tree produced by GsParser#selfAssignOp.
	VisitSelfAssignOp(ctx *SelfAssignOpContext) interface{}

	// Visit a parse tree produced by GsParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by GsParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by GsParser#logicalOrExpr.
	VisitLogicalOrExpr(ctx *LogicalOrExprContext) interface{}

	// Visit a parse tree produced by GsParser#logicalAndExpr.
	VisitLogicalAndExpr(ctx *LogicalAndExprContext) interface{}

	// Visit a parse tree produced by GsParser#comparisonExpr.
	VisitComparisonExpr(ctx *ComparisonExprContext) interface{}

	// Visit a parse tree produced by GsParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by GsParser#binExpr.
	VisitBinExpr(ctx *BinExprContext) interface{}

	// Visit a parse tree produced by GsParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by GsParser#powExpr.
	VisitPowExpr(ctx *PowExprContext) interface{}

	// Visit a parse tree produced by GsParser#negAtom.
	VisitNegAtom(ctx *NegAtomContext) interface{}

	// Visit a parse tree produced by GsParser#intAtom.
	VisitIntAtom(ctx *IntAtomContext) interface{}

	// Visit a parse tree produced by GsParser#floatAtom.
	VisitFloatAtom(ctx *FloatAtomContext) interface{}

	// Visit a parse tree produced by GsParser#charAtom.
	VisitCharAtom(ctx *CharAtomContext) interface{}

	// Visit a parse tree produced by GsParser#stringAtom.
	VisitStringAtom(ctx *StringAtomContext) interface{}

	// Visit a parse tree produced by GsParser#trueAtom.
	VisitTrueAtom(ctx *TrueAtomContext) interface{}

	// Visit a parse tree produced by GsParser#falseAtom.
	VisitFalseAtom(ctx *FalseAtomContext) interface{}

	// Visit a parse tree produced by GsParser#nilAtom.
	VisitNilAtom(ctx *NilAtomContext) interface{}

	// Visit a parse tree produced by GsParser#notAtom.
	VisitNotAtom(ctx *NotAtomContext) interface{}

	// Visit a parse tree produced by GsParser#qidAtom.
	VisitQidAtom(ctx *QidAtomContext) interface{}

	// Visit a parse tree produced by GsParser#callAtom.
	VisitCallAtom(ctx *CallAtomContext) interface{}

	// Visit a parse tree produced by GsParser#instanceAtom.
	VisitInstanceAtom(ctx *InstanceAtomContext) interface{}

	// Visit a parse tree produced by GsParser#arrayAtom.
	VisitArrayAtom(ctx *ArrayAtomContext) interface{}

	// Visit a parse tree produced by GsParser#indexAccessAtom.
	VisitIndexAccessAtom(ctx *IndexAccessAtomContext) interface{}

	// Visit a parse tree produced by GsParser#dictAtom.
	VisitDictAtom(ctx *DictAtomContext) interface{}

	// Visit a parse tree produced by GsParser#parenAtom.
	VisitParenAtom(ctx *ParenAtomContext) interface{}

	// Visit a parse tree produced by GsParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by GsParser#indexAccess.
	VisitIndexAccess(ctx *IndexAccessContext) interface{}

	// Visit a parse tree produced by GsParser#sliceExpr.
	VisitSliceExpr(ctx *SliceExprContext) interface{}

	// Visit a parse tree produced by GsParser#dictLiteral.
	VisitDictLiteral(ctx *DictLiteralContext) interface{}

	// Visit a parse tree produced by GsParser#strKeyEntry.
	VisitStrKeyEntry(ctx *StrKeyEntryContext) interface{}

	// Visit a parse tree produced by GsParser#idKeyEntry.
	VisitIdKeyEntry(ctx *IdKeyEntryContext) interface{}

	// Visit a parse tree produced by GsParser#instance.
	VisitInstance(ctx *InstanceContext) interface{}

	// Visit a parse tree produced by GsParser#qid.
	VisitQid(ctx *QidContext) interface{}

	// Visit a parse tree produced by GsParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by GsParser#compOp.
	VisitCompOp(ctx *CompOpContext) interface{}

	// Visit a parse tree produced by GsParser#addOp.
	VisitAddOp(ctx *AddOpContext) interface{}

	// Visit a parse tree produced by GsParser#bitOp.
	VisitBitOp(ctx *BitOpContext) interface{}

	// Visit a parse tree produced by GsParser#mulOp.
	VisitMulOp(ctx *MulOpContext) interface{}

	// Visit a parse tree produced by GsParser#powOp.
	VisitPowOp(ctx *PowOpContext) interface{}

}