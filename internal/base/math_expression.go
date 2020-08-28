package base

import (
	"gengine/context"
	"gengine/internal/core"
	"gengine/internal/core/errors"
)

type MathExpression struct {
	MathExpressionLeft       *MathExpression
	MathPmOperator           string
	MathMdOperator           string
	MathExpressionRight      *MathExpression
	ExpressionAtom           *ExpressionAtom
//	knowledgeContext         *KnowledgeContext
	dataCtx                  *context.DataContext
}

func (e *MathExpression) AcceptMathExpression(atom *MathExpression) error {
	if e.MathExpressionLeft == nil {
		e.MathExpressionLeft = atom
		return nil
	}
	if e.MathExpressionRight == nil{
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

func (e *MathExpression) AcceptExpressionAtom(atom *ExpressionAtom) error{
	if e.ExpressionAtom == nil {
		e.ExpressionAtom = atom
		return nil
	}
	return errors.New("ExpressionAtom already set twice!")
}


func (e *MathExpression) Evaluate(Vars map[string]interface{}) (interface{}, error) {

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
		return nil, err
	}
	rv, err := e.MathExpressionRight.Evaluate(Vars)
	if err != nil {
		return nil, err
	}

	if e.MathPmOperator == "+" {
		return core.Add(lv, rv)
	}

	if e.MathPmOperator == "-" {
		return core.Sub(lv, rv)
	}

	if e.MathMdOperator == "*" {
		return core.Mul(lv, rv)
	}

	if e.MathMdOperator == "/" {
		return core.Div(lv, rv)
	}
	return nil, errors.Errorf("MathExpression calculate evaluate error: %s, %s", e.MathPmOperator, e.MathPmOperator)
}


