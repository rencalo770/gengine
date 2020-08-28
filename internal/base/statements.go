package base

import (
	"gengine/context"
)

type Statements struct {
	StatementList []*Statement
	//	knowledgeContext *KnowledgeContext
	dataCtx *context.DataContext
}

func (s *Statements) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	for _, statement := range s.StatementList {
		_, err := statement.Evaluate(Vars)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (s *Statements) Initialize(dc *context.DataContext) {
	s.dataCtx = dc

	if s.StatementList != nil {
		for _, val := range s.StatementList {
			val.Initialize(dc)
		}
	}
}
