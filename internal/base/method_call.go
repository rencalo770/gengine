package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
)

type MethodCall struct {
	MethodName string
	MethodArgs *Args
	dataCtx *context.DataContext
}

func (mc *MethodCall) Initialize(dataCtx *context.DataContext) {
	mc.dataCtx = dataCtx

	if mc.MethodArgs != nil {
		mc.MethodArgs.Initialize(dataCtx)
	}
}

func (mc *MethodCall) AcceptArgs(funcArg *Args) error {
	if mc.MethodArgs == nil {
		mc.MethodArgs = funcArg
		return nil
	}
	return errors.New("methodArgs set twice!")
}

func (mc *MethodCall) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	var argumentValues []interface{}
	if mc.MethodArgs == nil {
		argumentValues = make([]interface{}, 0)
	} else {
		av, err := mc.MethodArgs.Evaluate(Vars)
		if err != nil {
			return nil, err
		}
		argumentValues = av
	}

	return mc.dataCtx.ExecMethod(mc.MethodName, argumentValues)
}
