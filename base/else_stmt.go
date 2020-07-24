package base

import (
	"gengine/context"
	"gengine/core/errors"
	"reflect"
)

type ElseStmt struct {
	StatementList    *Statements
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (e *ElseStmt) Evaluate(Vars map[string]interface{}) (reflect.Value, error) {

	if e.StatementList != nil {
		value, err:= e.StatementList.Evaluate(Vars)
		if err != nil {
			return reflect.ValueOf(nil),err
		}
		return value,nil
	}else {
		return reflect.ValueOf(nil),nil
	}
}

func (e *ElseStmt) Initialize(kc *KnowledgeContext, dc *context.DataContext) {
	e.knowledgeContext = kc
	e.dataCtx = dc
	if e.StatementList != nil {
		e.StatementList.Initialize(kc, dc)
	}
}

func (e *ElseStmt)AcceptStatements(stmts *Statements ) error {
	if e.StatementList  == nil {
		e.StatementList = stmts
		return nil
	}
	return errors.New("ElseStmt set twice!")
}
