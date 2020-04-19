package base

import (
	"fmt"
	"gengine/context"
)

type Arg struct {
	Constant         *Constant
	Variable         string
	FunctionCall     *FunctionCall
	MethodCall       *MethodCall
	MapVar           *MapVar
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (a *Arg) Initialize(kc *KnowledgeContext, dc *context.DataContext) {

	a.knowledgeContext = kc
	a.dataCtx = dc

	if a.Constant != nil {
		a.Constant.Initialize(kc, dc)
	}
	if a.FunctionCall != nil {
		a.FunctionCall.Initialize(kc, dc)
	}
	if a.MethodCall != nil {
		a.MethodCall.Initialize(kc, dc)
	}
	if a.MapVar != nil {
		a.MapVar.Initialize(kc, dc)
	}

}

func (a *Arg) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	if len(a.Variable) > 0 {
		//单值
		return a.dataCtx.GetValue(Vars, a.Variable)
	}

	if a.Constant != nil {
		//单值
		return a.Constant.Evaluate(Vars)
	}

	if a.FunctionCall != nil {
		//单值
		return a.FunctionCall.Evaluate(Vars)
	}

	if a.MethodCall != nil {
		//单值
		return a.MethodCall.Evaluate(Vars)
	}

	if a.MapVar != nil {
		return a.MapVar.Evaluate(Vars)
	}

	return nil, fmt.Errorf("argHolder holder has more values than want！")
}
