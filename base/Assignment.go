package base

import (
	"gengine/context"
	"gengine/core/errors"
	"reflect"
)

// := or =
type Assignment struct {
	Variable         string
	MathExpression   *MathExpression
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (a *Assignment) Evaluate(Vars map[string]interface{}) (reflect.Value, error) {
	v, err := a.MathExpression.Evaluate(Vars)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	err = a.dataCtx.SetValue(Vars, a.Variable, v)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return reflect.ValueOf(nil), nil
}

func (a *Assignment) Initialize(kc *KnowledgeContext, dc *context.DataContext) {
	a.knowledgeContext = kc
	a.dataCtx = dc

	if a.MathExpression != nil{
		a.MathExpression.Initialize(kc, dc)
	}
}

func (a *Assignment)AcceptMathExpression(me *MathExpression) error{
	if a.MathExpression == nil {
		a.MathExpression = me
		return nil
	}
	return errors.Errorf("%s","MathExpression already set twice")
}

func (a *Assignment)AcceptVariable(name string) error{
	if len(a.Variable) == 0 {
		a.Variable = name
		return nil
	}
	return errors.Errorf("Variable already set twice")
}