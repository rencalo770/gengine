package base

import (
	"fmt"
	"gengine/context"
	"gengine/internal/core"
	"gengine/internal/core/errors"
	"runtime"
	"strings"
)

// := or =
type Assignment struct {
	SourceCode
	Variable       string
	MapVar         *MapVar
	AssignOperator string
	MathExpression *MathExpression
	Expression     *Expression
	dataCtx        *context.DataContext
}

func (a *Assignment) Evaluate(Vars map[string]interface{}) (value interface{}, err error) {

	defer func() {
		if e := recover(); e != nil {
			size := 1 << 10 * 10
			buf := make([]byte, size)
			rs := runtime.Stack(buf, false)
			if rs > size {
				rs = size
			}
			buf = buf[:rs]
			eMsg := fmt.Sprintf("line %d, column %d, code: %s, %+v \n%s", a.LineNum, a.Column, a.Code, e, string(buf))
			eMsg = strings.Replace(eMsg,"panic","error", 10)
			err = errors.New(eMsg)
		}
	}()

	var mv interface{}

	if a.MathExpression != nil {
		mv, err = a.MathExpression.Evaluate(Vars)
		if err != nil {
			return nil, err
		}
	}

	if a.Expression != nil {
		mv, err = a.Expression.Evaluate(Vars)
		if err != nil {
			return nil, err
		}
	}

	var sv interface{}

	if a.AssignOperator == "=" || a.AssignOperator == ":=" {
		goto END
	}

	if len(a.Variable) > 0 {
		sv, err = a.dataCtx.GetValue(Vars, a.Variable)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column:%d, code: %s, %+v:", a.LineNum, a.Column, a.Code, err))
		}
	}

	if a.MapVar != nil {
		sv, err = a.MapVar.Evaluate(Vars)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column:%d, code: %s, %+v:", a.LineNum, a.Column, a.Code, err))
		}
	}

	//var pv interface{}
	if a.AssignOperator == "+=" {
		mv, err = core.Add(sv, mv)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column:%d, code: %s, %+v:", a.LineNum, a.Column, a.Code, err))
		}
		goto END
	}

	if a.AssignOperator == "-=" {
		mv, err = core.Sub(sv, mv)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column:%d, code: %s, %+v:", a.LineNum, a.Column, a.Code, err))
		}
		goto END
	}

	if a.AssignOperator == "*=" {
		mv, err = core.Mul(sv, mv)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column:%d, code: %s, %+v:", a.LineNum, a.Column, a.Code, err))
		}
		goto END
	}

	if a.AssignOperator == "/=" {
		mv, err = core.Div(sv, mv)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column:%d, code: %s, %+v:", a.LineNum, a.Column, a.Code, err))
		}
		goto END
	}

END:
	if len(a.Variable) > 0 {
		err = a.dataCtx.SetValue(Vars, a.Variable, mv)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", a.LineNum, a.Column, a.Code, err))
		}
		return
	}

	if a.MapVar != nil {
		err = a.dataCtx.SetMapVarValue(Vars, a.MapVar.Name, a.MapVar.Strkey, a.MapVar.Varkey, a.MapVar.Intkey, mv)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column:%d, code: %s, %+v:", a.LineNum, a.Column, a.Code, err))
		}
		return
	}

	return
}

func (a *Assignment) Initialize(dc *context.DataContext) {
	a.dataCtx = dc

	if a.MathExpression != nil {
		a.MathExpression.Initialize(dc)
	}

	if a.MapVar != nil {
		a.MapVar.Initialize(dc)
	}

	if a.Expression != nil {
		a.Expression.Initialize(dc)
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

func (a *Assignment) AcceptExpression(exp *Expression) error {
	if a.Expression == nil {
		a.Expression = exp
		return nil
	}
	return errors.New("Expression already set twice!")
}
