package base

import (
	"gengine/context"
	"gengine/core/errors"
	"reflect"
)

type IfStmt struct {
	Expression       *Expression
	StatementList    *Statements
	ElseStmt         *ElseStmt
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (i *IfStmt) Evaluate(Vars map[string]interface{}) (reflect.Value, error) {

	it ,err := i.Expression.Evaluate(Vars)
	if err != nil {
		return reflect.ValueOf(nil), err
	}

	if reflect.ValueOf(it).Bool() {
		if i.StatementList == nil{
			return reflect.ValueOf(nil),nil
		}else {
			return i.StatementList.Evaluate(Vars)
		}

	}else {
		if i.ElseStmt != nil{
			return i.ElseStmt.Evaluate(Vars)
		}else {
			return reflect.ValueOf(nil),nil
		}
	}
}

func (i *IfStmt) Initialize(kc *KnowledgeContext,  dc *context.DataContext) {
	i.knowledgeContext = kc
	i.dataCtx = dc

	if i.Expression != nil {
		i.Expression.Initialize(kc, dc)
	}
	if i.StatementList != nil {
		i.StatementList.Initialize(kc, dc)
	}

	if i.ElseStmt != nil {
		i.ElseStmt.Initialize(kc, dc)
	}
}

func (i *IfStmt)AcceptExpression(expr *Expression) error{
	if i.Expression == nil {
		i.Expression = expr
		return nil
	}
	return errors.New("IfStmt Expression set twice!")
}

func (i *IfStmt)AcceptStatements(stmts *Statements)error{
	if i.StatementList == nil {
		i.StatementList = stmts
		return nil
	}
	return errors.New("ifStmt's statements set twice!")
}