package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
	"reflect"
)

type Statement struct {
	IfStmt        *IfStmt
	MethodCall    *MethodCall
	FunctionCall  *FunctionCall
	Assignment    *Assignment
	ConcStatement *ConcStatement
	dataCtx       *context.DataContext
}

func (s *Statement) Evaluate(Vars map[string]reflect.Value) (reflect.Value, error, bool) {

	if s.IfStmt != nil {
		return s.IfStmt.Evaluate(Vars)
	}

	if s.MethodCall != nil {
		v, e := s.MethodCall.Evaluate(Vars)
		return v, e, false
	}

	if s.FunctionCall != nil {
		v, e := s.FunctionCall.Evaluate(Vars)
		return v, e, false
	}

	if s.Assignment != nil {
		v, e := s.Assignment.Evaluate(Vars)
		return v, e, false
	}

	if s.ConcStatement != nil {
		v, e := s.ConcStatement.Evaluate(Vars)
		return v, e, false
	}

	return reflect.ValueOf(nil), errors.New("Statement evaluate error!"), false
}

func (s *Statement) Initialize(dc *context.DataContext) {
	s.dataCtx = dc

	if s.IfStmt != nil {
		s.IfStmt.Initialize(dc)
	}

	if s.FunctionCall != nil {
		s.FunctionCall.Initialize(dc)
	}

	if s.MethodCall != nil {
		s.MethodCall.Initialize(dc)
	}

	if s.Assignment != nil {
		s.Assignment.Initialize(dc)
	}

	if s.ConcStatement != nil {
		s.ConcStatement.Initialize(dc)
	}
}

func (s *Statement) AcceptFunctionCall(funcCall *FunctionCall) error {
	if s.FunctionCall == nil {
		s.FunctionCall = funcCall
		return nil
	}
	return errors.New("FunctionCall already defined!")
}

func (s *Statement) AcceptMethodCall(methodCall *MethodCall) error {
	if s.MethodCall == nil {
		s.MethodCall = methodCall
		return nil
	}
	return errors.New("MethodCall already defined!")
}

func (s *Statement) AcceptAssignment(assignment *Assignment) error {
	if s.Assignment == nil {
		s.Assignment = assignment
		return nil
	}
	return errors.New("Assignment already defined!")
}
