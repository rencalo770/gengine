package base

import (
	"gengine/context"
	"gengine/core/errors"
)

type Statement struct {
	IfStmt           *IfStmt
	MethodCall       *MethodCall
	FunctionCall     *FunctionCall
	Assignment       *Assignment
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (s *Statement) Evaluate(Vars map[string]interface{}) (interface{}, error) {

	if s.IfStmt != nil {
		return s.IfStmt.Evaluate(Vars)
	}

	if s.MethodCall != nil {
		return s.MethodCall.Evaluate(Vars)
	}

	if s.FunctionCall != nil {
		return s.FunctionCall.Evaluate(Vars)
	}

	if s.Assignment != nil {
		return s.Assignment.Evaluate(Vars)
	}

	return nil,errors.New("Statement evaluate error!")
}

func (s *Statement) Initialize(kc *KnowledgeContext, dc *context.DataContext) {
	s.knowledgeContext = kc
	s.dataCtx = dc

	if s.IfStmt != nil {
		s.IfStmt.Initialize(kc, dc)
	}

	if s.FunctionCall != nil {
		s.FunctionCall.Initialize(kc, dc)
	}

	if s.MethodCall != nil {
		s.MethodCall.Initialize(kc, dc)
	}

	if s.Assignment != nil {
		s.Assignment.Initialize(kc, dc)
	}

}

func (s *Statement)AcceptFunctionCall(funcCall *FunctionCall) error  {
	if s.FunctionCall == nil {
		s.FunctionCall = funcCall
		return nil
	}
	return errors.New("FunctionCall already defined!")
}


func (s *Statement)AcceptMethodCall(methodCall *MethodCall) error{
	if s.MethodCall == nil {
		s.MethodCall = methodCall
		return nil
	}
	return errors.New("MethodCall already defined!")
}

