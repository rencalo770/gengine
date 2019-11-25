package iparser

import (
	"gengine/base"
	"gengine/core/errors"
	parser "gengine/iantlr/alr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang-collections/collections/stack"
	"strconv"
	"strings"
)

func NewGengineParserListener(ctx *base.KnowledgeContext) *GengineParserListener {
	return &GengineParserListener{
		Stack:         stack.New(),
		KnowledgeContext: ctx,
		ParseErrors:   make([]string, 0),
	}
}

type GengineParserListener struct {
	parser.BasegengineListener
	ParseErrors []string

	KnowledgeContext *base.KnowledgeContext
	Stack *stack.Stack
}

func (g *GengineParserListener)AddError(e error)  {
	g.ParseErrors = append(g.ParseErrors, e.Error())
}

func (g *GengineParserListener) VisitTerminal(node antlr.TerminalNode) {}

func (g *GengineParserListener) VisitErrorNode(node antlr.ErrorNode) {
	g.AddError(errors.Errorf("cannot recognize '"+ node.GetText()+ "' "))
}

func (g *GengineParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

func (g *GengineParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (g *GengineParserListener) EnterPrimary(ctx *parser.PrimaryContext) {}

func (g *GengineParserListener) ExitPrimary(ctx *parser.PrimaryContext) {}

func (g *GengineParserListener) EnterRuleEntity(ctx *parser.RuleEntityContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	entry := &base.RuleEntity{}
	g.Stack.Push(entry)
}

func (g *GengineParserListener) ExitRuleEntity(ctx *parser.RuleEntityContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	entry := g.Stack.Pop().(*base.RuleEntity)
	if _, ok := g.KnowledgeContext.RuleEntities[entry.RuleName]; ok {
		g.AddError(errors.Errorf("already existed entity's name \"%s\"", entry.RuleName))
		return
	}
	g.KnowledgeContext.RuleEntities[entry.RuleName] = entry
}

func (g *GengineParserListener) EnterRuleName(ctx *parser.RuleNameContext) {}

func (g *GengineParserListener) ExitRuleName(ctx *parser.RuleNameContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	ruleName := ctx.GetText()
	entry := g.Stack.Peek().(*base.RuleEntity)
	entry.RuleName = ruleName
}

func (g *GengineParserListener) EnterSalience(ctx *parser.SalienceContext) {
}

func (g *GengineParserListener) ExitSalience(ctx *parser.SalienceContext) {}

func (g *GengineParserListener) EnterRuleDescription(ctx *parser.RuleDescriptionContext) {}

func (g *GengineParserListener) ExitRuleDescription(ctx *parser.RuleDescriptionContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	ruleDescription := ctx.GetText()
	entry := g.Stack.Peek().(*base.RuleEntity)
	entry.RuleDescription = ruleDescription
}

func (g *GengineParserListener) EnterRuleContent(ctx *parser.RuleContentContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	ruleContent := &base.RuleContent{}
	g.Stack.Push(ruleContent)
}

func (g *GengineParserListener) ExitRuleContent(ctx *parser.RuleContentContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	ruleContent := g.Stack.Pop().(*base.RuleContent)
	entry := g.Stack.Peek().(*base.RuleEntity)
	entry.RuleContent = ruleContent
}

func (g *GengineParserListener) EnterAssignment(ctx *parser.AssignmentContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	assignment := &base.Assignment{}
	g.Stack.Push(assignment)
}

func (g *GengineParserListener) ExitAssignment(ctx *parser.AssignmentContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	assignment := g.Stack.Pop().(*base.Assignment)
	statement := g.Stack.Peek().(*base.Statement)
	statement.Assignment = assignment
}

func (g *GengineParserListener) EnterMathExpression(ctx *parser.MathExpressionContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	me := &base.MathExpression{}
	g.Stack.Push(me)

}

func (g *GengineParserListener) ExitMathExpression(ctx *parser.MathExpressionContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	expr := g.Stack.Pop().(*base.MathExpression)
	holder := g.Stack.Peek().(base.MathExpressionHolder)
	err := holder.AcceptMathExpression(expr)
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener) EnterExpression(ctx *parser.ExpressionContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	expression := &base.Expression{}
	g.Stack.Push(expression)
}

func (g *GengineParserListener) ExitExpression(ctx *parser.ExpressionContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	expr := g.Stack.Pop().(*base.Expression)
	holder := g.Stack.Peek().(base.ExpressionHolder)
	err := holder.AcceptExpression(expr)
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener) EnterExpressionAtom(ctx *parser.ExpressionAtomContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	exprAtom := &base.ExpressionAtom{}
	g.Stack.Push(exprAtom)
}

func (g *GengineParserListener) ExitExpressionAtom(ctx *parser.ExpressionAtomContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	exprAtom := g.Stack.Pop().(*base.ExpressionAtom)
	holder := g.Stack.Peek().(base.ExpressionAtomHolder)
	err := holder.AcceptExpressionAtom(exprAtom)
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener) EnterMethodCall(ctx *parser.MethodCallContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	funcCall := &base.MethodCall{
		MethodName: ctx.DOTTEDNAME().GetText(),
	}
	g.Stack.Push(funcCall)
}

func (g *GengineParserListener) ExitMethodCall(ctx *parser.MethodCallContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	methodCall := g.Stack.Pop().(*base.MethodCall)
	holder := g.Stack.Peek().(base.MethodCallHolder)
	err := holder.AcceptMethodCall(methodCall)
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	funcCall := &base.FunctionCall{
		FunctionName: ctx.SIMPLENAME().GetText(),
	}
	g.Stack.Push(funcCall)
}

func (g *GengineParserListener) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	funcCall := g.Stack.Pop().(*base.FunctionCall)
	holder := g.Stack.Peek().(base.FunctionCallHolder)
	err := holder.AcceptFunctionCall(funcCall)
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener) EnterFunctionArgs(ctx *parser.FunctionArgsContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	funcArg := &base.Args{
		ArgList: make([]*base.Arg, 0),
	}
	g.Stack.Push(funcArg)
}

func (g *GengineParserListener) ExitFunctionArgs(ctx *parser.FunctionArgsContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	funcArgs := g.Stack.Pop().(*base.Args)
	argHolder := g.Stack.Peek().(base.ArgsHolder)
	err := argHolder.AcceptArgs(funcArgs)
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener) EnterLogicalOperator(ctx *parser.LogicalOperatorContext) {}

func (g *GengineParserListener) ExitLogicalOperator(ctx *parser.LogicalOperatorContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	expr := g.Stack.Peek().(*base.Expression)
	// && and ||
	expr.LogicalOperator = ctx.GetText()
}

func (g *GengineParserListener) EnterNotOperator(ctx *parser.NotOperatorContext) {}

func (g *GengineParserListener) ExitNotOperator(ctx *parser.NotOperatorContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	expr := g.Stack.Peek().(*base.Expression)
	// !
	expr.NotOperator = ctx.GetText()
}

func (g *GengineParserListener) EnterVariable(ctx *parser.VariableContext) {}

func (g *GengineParserListener) ExitVariable(ctx *parser.VariableContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	varName := ctx.GetText()
	holder := g.Stack.Peek().(base.VariableHolder)
	err := holder.AcceptVariable(varName)
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener) EnterMathPmOperator(ctx *parser.MathPmOperatorContext) {}

func (g *GengineParserListener) ExitMathPmOperator(ctx *parser.MathPmOperatorContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	expr := g.Stack.Peek().(*base.MathExpression)
	// + -
	expr.MathPmOperator = ctx.GetText()
}

func (g *GengineParserListener) EnterMathMdOperator(ctx *parser.MathMdOperatorContext) {}

func (g *GengineParserListener) ExitMathMdOperator(ctx *parser.MathMdOperatorContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	expr := g.Stack.Peek().(*base.MathExpression)
	// * /
	expr.MathMdOperator = ctx.GetText()
}

func (g *GengineParserListener) EnterComparisonOperator(ctx *parser.ComparisonOperatorContext) {}

func (g *GengineParserListener) ExitComparisonOperator(ctx *parser.ComparisonOperatorContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	expr := g.Stack.Peek().(*base.Expression)
	// ==  !=  <  > <= >=
	expr.ComparisonOperator = ctx.GetText()
}

func (g *GengineParserListener) EnterConstant(ctx *parser.ConstantContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	cons := &base.Constant{}
	g.Stack.Push(cons)
}

func (g *GengineParserListener) ExitConstant(ctx *parser.ConstantContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	cons := g.Stack.Pop().(*base.Constant)
	holder := g.Stack.Peek().(base.ConstantHolder)
	err := holder.AcceptConstant(cons)
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {}

func (g *GengineParserListener) ExitStringLiteral(ctx *parser.StringLiteralContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	holder := g.Stack.Peek().(base.StringHolder)
	text := ctx.GetText()
	err := holder.AcceptString(strings.Trim(text, "\""))
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener) EnterBooleanLiteral(ctx *parser.BooleanLiteralContext) {}

func (g *GengineParserListener) ExitBooleanLiteral(ctx *parser.BooleanLiteralContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	cons := g.Stack.Peek().(*base.Constant)
	b, e := strconv.ParseBool(ctx.GetText())
	if e != nil{
		g.AddError(e)
		return
	}
	cons.ConstantValue = b
}

func (g *GengineParserListener) EnterRealLiteral(ctx *parser.RealLiteralContext) {}

func (g *GengineParserListener) ExitRealLiteral(ctx *parser.RealLiteralContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	cons := g.Stack.Peek().(*base.Constant)
	flo, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		g.AddError(errors.Errorf("string to float conversion error. String is not real type '%s'", ctx.GetText()))
		return
	}
	cons.ConstantValue = flo
}

func (g *GengineParserListener) EnterIfStmt(ctx *parser.IfStmtContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	ifStmt := &base.IfStmt{}
	g.Stack.Push(ifStmt)
}

func (g *GengineParserListener) ExitIfStmt(ctx *parser.IfStmtContext) {
	if len(g.ParseErrors) > 0 {
		return
	}
	ifStmt := g.Stack.Pop().(*base.IfStmt)
	statement := g.Stack.Peek().(*base.Statement)
	statement.IfStmt = ifStmt
}

func (g *GengineParserListener)EnterStatement(ctx *parser.StatementContext){
	if len(g.ParseErrors) > 0 {
		return
	}
	statement := &base.Statement{}
	g.Stack.Push(statement)
}

func (g *GengineParserListener)ExitStatement(ctx *parser.StatementContext){
	if len(g.ParseErrors) > 0 {
		return
	}
	statement := g.Stack.Pop().(*base.Statement)
	statements := g.Stack.Peek().(*base.Statements)
	statements.StatementList = append(statements.StatementList, statement)
}

func(g *GengineParserListener)EnterStatements(ctx *parser.StatementsContext){
	if len(g.ParseErrors) > 0 {
		return
	}
	statements := &base.Statements{
		StatementList: make([]*base.Statement, 0),
	}
	g.Stack.Push(statements)
}

func(g *GengineParserListener)ExitStatements(ctx *parser.StatementsContext){
	if len(g.ParseErrors) > 0 {
		return
	}
	statements := g.Stack.Pop().(*base.Statements)
	holder := g.Stack.Peek().(base.StatementsHolder)
	err := holder.AcceptStatements(statements)
	if err != nil {
		g.AddError(err)
	}
}

func (g *GengineParserListener)EnterAssignOperator(ctx *parser.AssignOperatorContext)  {}

func (g *GengineParserListener)ExitAssignOperator(ctx *parser.AssignOperatorContext)  {}

func (g *GengineParserListener)EnterSetOperator(ctx *parser.SetOperatorContext){}

func (g *GengineParserListener)ExitSetOperator(ctx *parser.SetOperatorContext){}

func (g *GengineParserListener)EnterElseStmt(ctx *parser.ElseStmtContext){
	if len(g.ParseErrors) > 0{
		return
	}
	elseStmt := &base.ElseStmt{}
	g.Stack.Push(elseStmt)
}

func (g *GengineParserListener)ExitElseStmt(ctx *parser.ElseStmtContext){
	if len(g.ParseErrors) > 0 {
		return
	}
	elseStmt := g.Stack.Pop().(*base.ElseStmt)
	ifStmt := g.Stack.Peek().(*base.IfStmt)
	ifStmt.ElseStmt = elseStmt
}

func (g *GengineParserListener) ExitInteger(ctx *parser.IntegerContext)  {
	if len(g.ParseErrors) > 0 {
		return
	}
	val,err := strconv.ParseInt(ctx.GetText(),10,64)
	if err !=nil {
		g.AddError(err)
		return
	}
	holder := g.Stack.Peek().(base.IntegerHolder)
	err = holder.AcceptInteger(val)
	if err != nil {
		g.AddError(err)
	}
}
func (g *GengineParserListener) EnterInteger(ctx *parser.IntegerContext)  {}