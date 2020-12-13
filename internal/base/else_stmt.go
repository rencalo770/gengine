package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
)

type ElseStmt struct {
	StatementList *Statements
	dataCtx       *context.DataContext
}

func (e *ElseStmt) Evaluate(Vars map[string]interface{}) (interface{}, error, bool) {

	if e.StatementList != nil {
		return e.StatementList.Evaluate(Vars)
	} else {
		return nil, nil, false
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
