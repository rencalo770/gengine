package engine

import (
	"gengine/builder"
	"gengine/core/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

type Gengine struct {
}

func NewGengine() *Gengine {
	return &Gengine{}
}

/**
	when b is true it means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
 */
func (g *Gengine) Execute(rb *builder.RuleBuilder, b bool) error {
	if len(rb.Kc.RuleEntities) == 0 {
		return nil
	}

	for _,r := range rb.Kc.SortRules{
		err := r.Execute()
		if err != nil {
			if b {
				logrus.Errorf("rule: %s executed, error: %+v ",r.RuleName, err)
			} else {
				return errors.Errorf("rule: %s executed, error: %v ",r.RuleName, err)
			}
		}
	}
	return nil
}

/*
 concurrent execute rules
 in this mode, it will not consider the salience  and err control
 */
func (g *Gengine) ExecuteConcurrent(rb * builder.RuleBuilder){
	var wg sync.WaitGroup
	wg.Add(len(rb.Kc.RuleEntities))
	for _,r := range rb.Kc.RuleEntities {
		go func() {
			e := r.Execute()
			if e != nil {
				logrus.Errorf("in rule:%s execute rule err:  %+v", r.RuleName, e)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}