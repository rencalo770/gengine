package base

import (
	"errors"
	"gengine/context"
	"reflect"
)

type ElseStmt struct {
	StatementList *Statements
	dataCtx       *context.DataContext
}

func (e *ElseStmt) Evaluate(Vars map[string]reflect.Value) (reflect.Value, error, bool) {

	if e.StatementList != nil {
		return e.StatementList.Evaluate(Vars)
	} else {
		return reflect.ValueOf(nil), nil, false
	}
}

func (e *ElseStmt) Initialize(dc *context.DataContext) {
	e.dataCtx = dc
	if e.StatementList != nil {
		e.StatementList.Initialize(dc)
	}
}

func (e *ElseStmt) AcceptStatements(stmts *Statements) error {
	if e.StatementList == nil {
		e.StatementList = stmts
		return nil
	}
	return errors.New("ElseStmt set twice!")
}
