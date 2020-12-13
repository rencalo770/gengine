package base

import (
	"gengine/context"
)

type Statements struct {
	StatementList   []*Statement
	ReturnStatement *ReturnStatement
	dataCtx         *context.DataContext
}

func (s *Statements) Evaluate(Vars map[string]interface{}) (interface{}, error, bool) {
	for _, statement := range s.StatementList {
		v, err, b := statement.Evaluate(Vars)
		if err != nil {
			return nil, err, false
		}

		if b {
			//important: meet return，not continue to execute
			return v, nil, b
		}
	}
	if s.ReturnStatement != nil {
		return s.ReturnStatement.Evaluate(Vars)
	} else {
		return nil, nil, false
	}

}

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
