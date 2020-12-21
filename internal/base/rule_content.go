package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
	"reflect"
)

type RuleContent struct {
	Statements *Statements
	dataCtx    *context.DataContext
}

func (t *RuleContent) Initialize(dc *context.DataContext) {
	t.dataCtx = dc
	if t.Statements != nil {
		t.Statements.Initialize(dc)
	}
}

func (t *RuleContent) Execute(Vars map[string]reflect.Value) (reflect.Value, error, bool) {
	return t.Statements.Evaluate(Vars)
}

func (t *RuleContent) AcceptStatements(stmts *Statements) error {
	if t.Statements == nil {
		t.Statements = stmts
		return nil
	}
	return errors.New("RuleContent's statements set twice.")
}
