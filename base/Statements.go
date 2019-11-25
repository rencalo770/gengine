package base

import (
	"gengine/context"
	"reflect"
)

type Statements struct {
	StatementList    []*Statement
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (s *Statements) Evaluate(Vars map[string]interface{}) (reflect.Value, error) {
	for _, val := range s.StatementList{
		_, err := val.Evaluate(Vars)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
	}
	return reflect.ValueOf(nil), nil
}

func (s *Statements) Initialize(kc *KnowledgeContext, dc *context.DataContext) {
	s.knowledgeContext = kc
	s.dataCtx = dc

	if s.StatementList != nil {
		for _, val := range s.StatementList{
			val.Initialize(kc, dc)
		}
	}
}









