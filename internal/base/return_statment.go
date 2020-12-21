package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
	"reflect"
)

type ReturnStatement struct {
	Expression *Expression
	dataCtx    *context.DataContext
}

func (rs *ReturnStatement) Evaluate(Vars map[string]reflect.Value) (reflect.Value, error, bool) {
	if rs.Expression != nil {
		value, e := rs.Expression.Evaluate(Vars)
		return value, e, true
	}
	return reflect.ValueOf(nil), nil, true
}

func (rs *ReturnStatement) Initialize(dc *context.DataContext) {
	rs.dataCtx = dc
	if rs.Expression != nil {
		rs.Expression.Initialize(dc)
	}
}

func (rs *ReturnStatement) AcceptExpression(expr *Expression) error {
	if rs.Expression == nil {
		rs.Expression = expr
		return nil
	}
	return errors.New("Expression already set twice!")
}
