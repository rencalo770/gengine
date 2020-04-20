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
				return errors.Errorf("rule: %s executed, error: %+v ",r.RuleName, err)
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
		rr := r
		go func() {
			e := rr.Execute()
			if e != nil {
				logrus.Errorf("in rule:%s execute rule err:  %+v", r.RuleName, e)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}


/*
 mix model to execute rules
 in this mode, it will not consider the priority，and it also concurrently to execute rules
 first to execute the most high priority rule，then concurrently to execute last rules without consider the priority
*/
func (g *Gengine) ExecuteMixModel(rb * builder.RuleBuilder){
	rules := rb.Kc.SortRules
	if len(rules) > 0 {
		e := rules[0].Execute()
		if e != nil {
			logrus.Errorf("the most high priority rule: [%s]  exe err:%+v",rules[0].RuleName, e)
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(rules) - 1)
	for _,r := range rules[1:] {
		rr := r
		go func() {
			e := rr.Execute()
			if e != nil {
				logrus.Errorf("in rule:%s execute rule err:  %+v", r.RuleName, e)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}