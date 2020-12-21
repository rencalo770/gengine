package base

import (
	"fmt"
	"gengine/context"
	"gengine/internal/core"
	"gengine/internal/core/errors"
	"reflect"
)

type MathExpression struct {
	SourceCode
	MathExpressionLeft  *MathExpression
	MathPmOperator      string
	MathMdOperator      string
	MathExpressionRight *MathExpression
	ExpressionAtom      *ExpressionAtom
	dataCtx             *context.DataContext
}

func (e *MathExpression) AcceptMathExpression(atom *MathExpression) error {
	if e.MathExpressionLeft == nil {
		e.MathExpressionLeft = atom
		return nil
	}
	if e.MathExpressionRight == nil {
		e.MathExpressionRight = atom
		return nil
	}
	return errors.New("expressionAtom set twice")
}

func (e *MathExpression) Initialize(dc *context.DataContext) {
	e.dataCtx = dc

	if e.MathExpressionLeft != nil {
		e.MathExpressionLeft.Initialize(dc)
	}

	if e.MathExpressionRight != nil {
		e.MathExpressionRight.Initialize(dc)
	}

	if e.ExpressionAtom != nil {
		e.ExpressionAtom.Initialize(dc)
	}
}

func (e *MathExpression) AcceptExpressionAtom(atom *ExpressionAtom) error {
	if e.ExpressionAtom == nil {
		e.ExpressionAtom = atom
		return nil
	}
	return errors.New("ExpressionAtom already set twice!")
}

func (e *MathExpression) Evaluate(Vars map[string]reflect.Value) (reflect.Value, error) {

	//priority to calculate single value
	if e.ExpressionAtom != nil {
		return e.ExpressionAtom.Evaluate(Vars)
	}

	// check the right whether is nil
	if e.MathExpressionRight == nil {
		return e.MathExpressionLeft.Evaluate(Vars)
	}

	lv, err := e.MathExpressionLeft.Evaluate(Vars)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	rv, err := e.MathExpressionRight.Evaluate(Vars)
	if err != nil {
		return reflect.ValueOf(nil), err
	}

	if e.MathPmOperator == "+" {
		add, err := core.Add(lv, rv)
		if err != nil {
			return reflect.ValueOf(nil), errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", e.LineNum, e.Column, e.Code, err))
		}
		return reflect.ValueOf(add), nil
	}

	if e.MathPmOperator == "-" {
		sub, err := core.Sub(lv, rv)
		if err != nil {
			return reflect.ValueOf(nil), errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", e.LineNum, e.Column, e.Code, err))
		}
		return reflect.ValueOf(sub), nil
	}

	if e.MathMdOperator == "*" {
		mul, err := core.Mul(lv, rv)
		if err != nil {
			return reflect.ValueOf(nil), errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", e.LineNum, e.Column, e.Code, err))
		}
		return reflect.ValueOf(mul), nil
	}

	if e.MathMdOperator == "/" {
		div, err := core.Div(lv, rv)
		if err != nil {
			return reflect.ValueOf(nil), errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", e.LineNum, e.Column, e.Code, err))
		}
		return reflect.ValueOf(div), nil
	}
	return reflect.ValueOf(nil), errors.New(fmt.Sprintf("line %d, column %d, code: %s, MathExpression calculate evaluate error", e.LineNum, e.Column, e.Code))
}
