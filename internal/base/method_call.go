package base

import (
	"fmt"
	"gengine/context"
	"gengine/internal/core/errors"
	"runtime"
	"strings"
)

type MethodCall struct {
	SourceCode
	MethodName string
	MethodArgs *Args
	dataCtx    *context.DataContext
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

func (mc *MethodCall) Evaluate(Vars map[string]interface{}) (mr interface{}, err error) {

	defer func() {
		if e := recover(); e != nil {
			size := 1 << 10 * 10
			buf := make([]byte, size)
			rs := runtime.Stack(buf, false)
			if rs > size {
				rs = size
			}
			buf = buf[:rs]
			eMsg := fmt.Sprintf("line %d, column %d, code: %s, %+v \n%s", mc.LineNum, mc.Column, mc.Code, e, string(buf))
			eMsg = strings.Replace(eMsg,"panic", "error",10)
			err = errors.New(eMsg)
		}
	}()

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

	mr, err = mc.dataCtx.ExecMethod(mc.MethodName, argumentValues)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", mc.LineNum, mc.Column, mc.Code, err))
	}
	return
}
