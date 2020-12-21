package base

import (
	"fmt"
	"gengine/context"
	"gengine/internal/core/errors"
	"reflect"
	"runtime"
	"strings"
)

type FunctionCall struct {
	SourceCode
	FunctionName string
	FunctionArgs *Args
	dataCtx      *context.DataContext
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

func (fc *FunctionCall) Evaluate(Vars map[string]reflect.Value) (res reflect.Value, err error) {

	defer func() {
		if e := recover(); e != nil {
			size := 1 << 10 * 10
			buf := make([]byte, size)
			rs := runtime.Stack(buf, false)
			if rs > size {
				rs = size
			}
			buf = buf[:rs]
			eMsg := fmt.Sprintf("line %d, column %d, code: %s, %+v \n%s", fc.LineNum, fc.Column, fc.Code, e, string(buf))
			eMsg = strings.ReplaceAll(eMsg, "panic", "error")
			err = errors.New(eMsg)
		}
	}()

	var argumentValues []reflect.Value
	if fc.FunctionArgs == nil {
		argumentValues = nil
	} else {
		av, err := fc.FunctionArgs.Evaluate(Vars)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		argumentValues = av
	}

	res, e := fc.dataCtx.ExecFunc(fc.FunctionName, argumentValues)
	if e != nil {
		return reflect.ValueOf(nil), errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", fc.LineNum, fc.Column, fc.Code, e))
	}
	return //res, nil
}
