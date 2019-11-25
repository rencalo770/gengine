package base

import "reflect"

type Evaluator interface {
	Evaluate(Vars map[string]interface{}) (reflect.Value, error)
}