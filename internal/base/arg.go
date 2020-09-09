package base

import (
	"errors"
	"gengine/context"
)

type Arg struct {
	Constant     *Constant
	Variable     string
	FunctionCall *FunctionCall
	MethodCall   *MethodCall
	MapVar       *MapVar
	dataCtx      *context.DataContext
}

func (a *Arg) Initialize(dc *context.DataContext) {

	a.dataCtx = dc

	if a.Constant != nil {
		a.Constant.Initialize(dc)
	}
	if a.FunctionCall != nil {
		a.FunctionCall.Initialize(dc)
	}
	if a.MethodCall != nil {
		a.MethodCall.Initialize(dc)
	}
	if a.MapVar != nil {
		a.MapVar.Initialize(dc)
	}

}

func (a *Arg) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	if len(a.Variable) > 0 {
		return a.dataCtx.GetValue(Vars, a.Variable)
	}

	if a.Constant != nil {
		return a.Constant.Evaluate(Vars)
	}

	if a.FunctionCall != nil {
		return a.FunctionCall.Evaluate(Vars)
	}

	if a.MethodCall != nil {
		return a.MethodCall.Evaluate(Vars)
	}

	if a.MapVar != nil {
		return a.MapVar.Evaluate(Vars)
	}

	return nil, errors.New("argHolder holder has more values than want！")
}
