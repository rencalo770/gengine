package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
	"reflect"
)

type ElseIfStmt struct {
	Expression    *Expression
	StatementList *Statements
	dataCtx       *context.DataContext
}

func (ef *ElseIfStmt) Evaluate(Vars map[string]reflect.Value) (reflect.Value, error, bool) {
	it, err := ef.Expression.Evaluate(Vars)
	if err != nil {
		return reflect.ValueOf(nil), err, false
	}

	if it.Bool() {
		if ef.StatementList == nil {
			return reflect.ValueOf(nil), nil, false
		} else {
			return ef.StatementList.Evaluate(Vars)
		}
	} else {
		return reflect.ValueOf(nil), nil, false
	}
}

func (ef *ElseIfStmt) Initialize(dc *context.DataContext) {
	ef.dataCtx = dc

	if ef.Expression != nil {
		ef.Expression.Initialize(dc)
	}

	if ef.StatementList != nil {
		ef.StatementList.Initialize(dc)
	}
}

func (ef *ElseIfStmt) AcceptExpression(expr *Expression) error {

	if ef.Expression == nil {
		ef.Expression = expr
		return nil
	}
	return errors.New("ElseIfStmt's Expression set twice!")
}

func (ef *ElseIfStmt) AcceptStatements(stmts *Statements) error {
	if ef.StatementList == nil {
		ef.StatementList = stmts
		return nil
	}
	return errors.New("ElseIfStmt's statements set twice!")
}
