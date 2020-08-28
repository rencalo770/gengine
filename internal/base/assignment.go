package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
)

// := or =
type Assignment struct {
	Variable       string
	MapVar         *MapVar
	MathExpression *MathExpression
	dataCtx *context.DataContext
}

func (a *Assignment) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	v, err := a.MathExpression.Evaluate(Vars)
	if err != nil {
		return nil, err
	}

	if len(a.Variable) > 0 {
		err = a.dataCtx.SetValue(Vars, a.Variable, v)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	if a.MapVar != nil {
		err := a.dataCtx.SetMapVarValue(Vars, a.MapVar.Name, a.MapVar.Strkey, a.MapVar.Varkey, a.MapVar.Intkey, v)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (a *Assignment) Initialize(dc *context.DataContext) {
	a.dataCtx = dc

	if a.MathExpression != nil {
		a.MathExpression.Initialize(dc)
	}

	if a.MapVar != nil {
		a.MapVar.Initialize(dc)
	}
}

func (a *Assignment) AcceptMathExpression(me *MathExpression) error {
	if a.MathExpression == nil {
		a.MathExpression = me
		return nil
	}
	return errors.New("MathExpression already set twice!")
}

func (a *Assignment) AcceptVariable(name string) error {
	if len(a.Variable) == 0 {
		a.Variable = name
		return nil
	}
	return errors.New("Variable already set twice!")
}

func (a *Assignment) AcceptMapVar(mapVar *MapVar) error {
	if a.MapVar == nil {
		a.MapVar = mapVar
		return nil
	}
	return errors.New("MapVar already set twice!")
}
