package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
	"reflect"
)

type ElseIfStmt struct {
	Expression    *Expression
	StatementList *Statements
	dataCtx *context.DataContext
}

func (ef *ElseIfStmt) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	it, err := ef.Expression.Evaluate(Vars)
	if err != nil {
		return nil, err
	}

	if reflect.ValueOf(it).Bool() {
		if ef.StatementList == nil {
			return nil, nil
		} else {
			return ef.StatementList.Evaluate(Vars)
		}
	} else {
		return nil, nil
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
