package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
	"reflect"
)

var TypeMap = map[string]string{
	"int":     "int",
	"int8":    "int8",
	"int16":   "int16",
	"int32":   "int32",
	"int64":   "int64",
	"uint":    "uint",
	"uint8":   "uint8",
	"uint16":  "uint16",
	"uint32":  "uint32",
	"uint64":  "uint64",
	"float32": "float32",
	"float64": "float64",
}

type Expression struct {
	ExpressionLeft     *Expression
	ExpressionRight    *Expression
	ExpressionAtom     *ExpressionAtom
	MathExpression     *MathExpression
	LogicalOperator    string
	ComparisonOperator string
	NotOperator        string
	dataCtx *context.DataContext
}

func (e *Expression) Initialize(dc *context.DataContext) {
	e.dataCtx = dc

	if e.ExpressionLeft != nil {
		e.ExpressionLeft.Initialize(dc)
	}
	if e.ExpressionRight != nil {
		e.ExpressionRight.Initialize(dc)
	}

	if e.ExpressionAtom != nil {
		e.ExpressionAtom.Initialize(dc)
	}

	if e.MathExpression != nil {
		e.MathExpression.Initialize(dc)
	}
}

func (e *Expression) AcceptExpressionAtom(atom *ExpressionAtom) error {
	if e.ExpressionAtom == nil {
		e.ExpressionAtom = atom
		return nil
	}
	return errors.New("ExpressionAtom already set twice!")
}

func (e *Expression) AcceptMathExpression(atom *MathExpression) error {
	if e.MathExpression == nil {
		e.MathExpression = atom
		return nil
	}
	return errors.New(" Expression's MathExpression set twice")
}

func (e *Expression) AcceptExpression(expression *Expression) error {
	if e.ExpressionLeft == nil {
		e.ExpressionLeft = expression
		return nil
	}

	if e.ExpressionRight == nil {
		e.ExpressionRight = expression
		return nil
	}
	return errors.New("Expression already set twice! ")
}

func (e *Expression) Evaluate(Vars map[string]interface{}) (interface{}, error) {

	//priority to calculate single value
	var math interface{}
	if e.MathExpression != nil {
		evl, err := e.MathExpression.Evaluate(Vars)
		if err != nil {
			return nil, err
		}
		math = evl
	}

	var atom interface{}
	if e.ExpressionAtom != nil {
		evl, err := e.ExpressionAtom.Evaluate(Vars)
		if err != nil {
			return nil, err
		}
		atom = evl
	}

	var b interface{}
	if e.ExpressionRight == nil {
		if e.ExpressionLeft != nil {
			left, err := e.ExpressionLeft.Evaluate(Vars)
			if err != nil {
				return nil, err
			}
			b = left
		}
	}

	// && ||  just only to be used between boolean
	if e.LogicalOperator != "" {

		lv, err := e.ExpressionLeft.Evaluate(Vars)
		if err != nil {
			return nil, err
		}

		rv, err := e.ExpressionRight.Evaluate(Vars)
		if err != nil {
			return nil, err
		}

		//
		flv := reflect.ValueOf(lv)
		frv := reflect.ValueOf(rv)

		if reflect.TypeOf(lv).String() == "bool" && reflect.TypeOf(rv).String() == "bool" {
			if e.LogicalOperator == "&&" {
				b = flv.Bool() && frv.Bool()
			}
			if e.LogicalOperator == "||" {
				b = flv.Bool() || frv.Bool()
			}
		} else {
			return nil, errors.Errorf(" || or && can't be used between %s and %s", flv.Kind().String(), frv.Kind().String())
		}
	}

	// == > < != >= <=  just only to be used between number and number, string and string, bool and bool
	if e.ComparisonOperator != "" {

		lv, err := e.ExpressionLeft.Evaluate(Vars)
		if err != nil {
			return nil, err
		}

		rv, err := e.ExpressionRight.Evaluate(Vars)
		if err != nil {
			return nil, err
		}

		//
		flv := reflect.ValueOf(lv)
		frv := reflect.ValueOf(rv)

		//string compare
		tlv := reflect.TypeOf(lv).String()
		trv := reflect.TypeOf(rv).String()
		if tlv == "string" && trv == "string" {
			switch e.ComparisonOperator {
			case "==":
				b = flv.String() == frv.String()
				break
			case "!=":
				b = flv.String() != frv.String()
				break
			case ">":
				b = flv.String() > frv.String()
				break
			case "<":
				b = flv.String() < frv.String()
				break
			case ">=":
				b = flv.String() >= frv.String()
				break
			case "<=":
				b = flv.String() <= frv.String()
				break
			default:
				return nil, errors.Errorf("Can't be recognized ComparisonOperator: %s", e.ComparisonOperator)
			}
			goto LAST
		}

		//data compare
		if l, ok1 := TypeMap[tlv]; ok1 {
			if r, ok2 := TypeMap[trv]; ok2 {
				var ll float64
				switch l {
				case "int", "int8", "int16", "int32", "int64":
					ll = float64(flv.Int())
					break
				case "uint", "uint8", "uint16", "uint32", "uint64":
					ll = float64(flv.Uint())
					break
				case "float32", "float64":
					ll = flv.Float()
					break
				}

				var rr float64
				switch r {
				case "int", "int8", "int16", "int32", "int64":
					rr = float64(frv.Int())
					break
				case "uint", "uint8", "uint16", "uint32", "uint64":
					rr = float64(frv.Uint())
					break
				case "float32", "float64":
					rr = frv.Float()
					break
				}

				switch e.ComparisonOperator {
				case "==":
					b = ll == rr
					break
				case "!=":
					b = ll != rr
					break
				case ">":
					b = ll > rr
					break
				case "<":
					b = ll < rr
					break
				case ">=":
					b = ll >= rr
					break
				case "<=":
					b = ll <= rr
					break
				default:
					return nil, errors.Errorf("Can't be recognized ComparisonOperator: %s", e.ComparisonOperator)
				}
			}
			goto LAST
		}

		if tlv == "bool" && trv == "bool" {
			switch e.ComparisonOperator {
			case "==":
				b = flv.Bool() == frv.Bool()
				break
			case "!=":
				b = flv.Bool() != frv.Bool()
				break
			default:
				return nil, errors.Errorf("Can't be recognized ComparisonOperator: %s", e.ComparisonOperator)
			}
			goto LAST
		}
	}

LAST:
	if e.NotOperator == "!" {

		if math != nil {
			return !reflect.ValueOf(math).Bool(), nil
		}

		if atom != nil {
			return !reflect.ValueOf(atom).Bool(), nil
		}

		if b != nil {
			return !reflect.ValueOf(b).Bool(), nil
		}
	} else {
		if math != nil {
			return math, nil
		}

		if atom != nil {
			return atom, nil
		}

		if b != nil {
			return b, nil
		}
	}
	return nil, errors.New("evaluate Expression err!")
}
