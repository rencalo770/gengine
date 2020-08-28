package base

import (
	"gengine/context"
)

type FunctionCall struct {
	FunctionName string
	FunctionArgs *Args
	//	knowledgeContext  *KnowledgeContext
	dataCtx *context.DataContext
}

func (fc *FunctionCall) AcceptArgs(funcArg *Args) error {
	fc.FunctionArgs = funcArg
	return nil
}

func (fc *FunctionCall) Initialize(dc *context.DataContext) {
	fc.dataCtx = dc

	if fc.FunctionArgs != nil {
		fc.FunctionArgs.Initialize(dc)
	}
}

func (fc *FunctionCall) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	var argumentValues []interface{}
	if fc.FunctionArgs == nil {
		argumentValues = nil
	} else {
		av, err := fc.FunctionArgs.Evaluate(Vars)
		if err != nil {
			return nil, err
		}
		argumentValues = av
	}

	return fc.dataCtx.ExecFunc(fc.FunctionName, argumentValues)
}
