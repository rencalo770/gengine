package base

import (
	"gengine/context"
	"reflect"
)

type Statements struct {
	StatementList   []*Statement
	ReturnStatement *ReturnStatement
	//dataCtx         *context.DataContext
}

func (s *Statements) Evaluate(dc *context.DataContext, Vars map[string]reflect.Value) (reflect.Value, error, bool) {
	for _, statement := range s.StatementList {
		v, err, b := statement.Evaluate(dc, Vars)
		if err != nil {
			return reflect.ValueOf(nil), err, false
		}

		if b {
			//important: meet return，not continue to execute
			return v, nil, b
		}
	}
	if s.ReturnStatement != nil {
		return s.ReturnStatement.Evaluate(dc, Vars)
	} else {
		return reflect.ValueOf(nil), nil, false
	}

}
/*
func (s *Statements) Initialize(dc *context.DataContext) {
	s.dataCtx = dc

	if s.StatementList != nil {
		for _, val := range s.StatementList {
			val.Initialize(dc)
		}
	}

	if s.ReturnStatement != nil {
		s.ReturnStatement.Initialize(dc)
	}
}
*/