package base

import (
	"gengine/context"
)

type Constant struct {
	ConstantValue interface{}
	dataCtx       *context.DataContext
}

func (cons *Constant) AcceptString(str string) error {
	cons.ConstantValue = str
	return nil
}

func (cons *Constant) Initialize(dc *context.DataContext) {
	cons.dataCtx = dc
}

func (cons *Constant) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	return cons.ConstantValue, nil
}

func (cons *Constant) AcceptInteger(i64 int64) error {
	cons.ConstantValue = i64
	return nil
}

//receive rule's name
func (cons *Constant) AcceptName(name string) error {
	cons.ConstantValue = name
	return nil
}
