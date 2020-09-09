package engine

import (
	"fmt"
	"gengine/builder"
	"gengine/internal/base"
	"gengine/internal/core/errors"
	"github.com/google/martian/log"
	"sort"
	"sync"
)

type Gengine struct {
}

func NewGengine() *Gengine {
	return &Gengine{}
}

type Stag struct {
	StopTag bool
}

/**
sort execute model

	when b is true it means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
*/
func (g *Gengine) Execute(rb *builder.RuleBuilder, b bool) error {
	if len(rb.Kc.RuleEntities) == 0 {
		return nil
	}

	for _, r := range rb.Kc.SortRules {
		err := r.Execute()
		if err != nil {
			if b {
				log.Errorf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, err)
			} else {
				return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, err))
			}
		}
	}
	return nil
}

/**
sort execute model

when b is true it means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule;
if stopTag become true,it will not continue to execute

sTag is a struct given by user, and user can use it  to control rules execute behavior in rules, it can improve performance

it used in this scene:
where some high priority rules execute finished, you don't want to execute to the last rules, you can use sTag to control it out of gengine
*/
func (g *Gengine) ExecuteWithStopTagDirect(rb *builder.RuleBuilder, b bool, sTag *Stag) error {
	if len(rb.Kc.RuleEntities) == 0 {
		return nil
	}

	for _, r := range rb.Kc.SortRules {
		err := r.Execute()
		if err != nil {
			if b {
				log.Errorf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, err)
			} else {
				return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, err))
			}
		}

		if sTag.StopTag {
			break
		}
	}
	return nil
}

/*
 concurrent execute model
 in this mode, it will not consider the priority  and not consider err control
*/
func (g *Gengine) ExecuteConcurrent(rb *builder.RuleBuilder) {
	if len(rb.Kc.RuleEntities) >= 1 {
		var wg sync.WaitGroup
		wg.Add(len(rb.Kc.RuleEntities))
		for _, r := range rb.Kc.RuleEntities {
			rr := r
			go func() {
				e := rr.Execute()
				if e != nil {
					log.Errorf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, e)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

/*
 mix model to execute rules

 in this mode, it will not consider the priority，and it also concurrently to execute rules
 first to execute the most high priority rule，then concurrently to execute last rules without consider the priority
*/
func (g *Gengine) ExecuteMixModel(rb *builder.RuleBuilder) {
	rules := rb.Kc.SortRules
	if len(rules) > 0 {
		e := rules[0].Execute()
		if e != nil {
			log.Errorf("the most high priority rule: \"%s\"  executed, error:\n %+v", rules[0].RuleName, e)
		}
	} else {
		return
	}

	if (len(rules) - 1) >= 1 {
		var wg sync.WaitGroup
		wg.Add(len(rules) - 1)
		for _, r := range rules[1:] {
			rr := r
			go func() {
				e := rr.Execute()
				if e != nil {
					log.Errorf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, e)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

/**
 mix execute model

base type :golang translate value
not base type: golang translate pointer

if stopTag become true,it will not continue to execute
stopTag is a name given by user, and user can use it  to control rules execute behavior in rules, it can improve performance

it used in this scene:
where the first rule execute finished, you don't want to execute to the last rules, you can use sTag to control it out of gengine

*/
func (g *Gengine) ExecuteMixModelWithStopTagDirect(rb *builder.RuleBuilder, sTag *Stag) {

	rules := rb.Kc.SortRules
	if len(rules) > 0 {
		e := rules[0].Execute()
		if e != nil {
			log.Errorf("the most high priority rule: \"%s\"  executed, error:\n %+v", rules[0].RuleName, e)
		}
	} else {
		return
	}

	if !sTag.StopTag {
		if (len(rules) - 1) >= 1 {
			var wg sync.WaitGroup
			wg.Add(len(rules) - 1)
			for _, r := range rules[1:] {
				rr := r
				go func() {
					e := rr.Execute()
					if e != nil {
						log.Errorf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, e)
					}
					wg.Done()
				}()
			}
			wg.Wait()
		}
	}
}

/**
user can choose specified name rules to run with sort
*/
func (g *Gengine) ExecuteSelectedRules(rb *builder.RuleBuilder, names []string) {
	rules := []*base.RuleEntity{}
	for _, name := range names {
		if ruleEntity, ok := rb.Kc.RuleEntities[name]; ok {
			rr := ruleEntity
			rules = append(rules, rr)
		} else {
			log.Infof("no such rule named: \"%s\"", name)
		}
	}

	if len(rules) < 1 {
		return
	}

	if len(rules) >= 2 {
		sort.SliceStable(rules, func(i, j int) bool {
			return rules[i].Salience > rules[j].Salience
		})
	}

	for _, rule := range rules {
		rr := rule
		e := rr.Execute()
		if e != nil {
			log.Errorf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e)
		}
	}
}

/**
user can choose specified name rules to concurrent run
*/
func (g *Gengine) ExecuteSelectedRulesConcurrent(rb *builder.RuleBuilder, names []string) {

	rules := []*base.RuleEntity{}
	for _, name := range names {
		if ruleEntity, ok := rb.Kc.RuleEntities[name]; ok {
			rr := ruleEntity
			rules = append(rules, rr)
		} else {
			log.Infof("no such rule named: \"%s\"", name)
		}
	}

	if len(rules) < 1 {
		return
	}

	if len(rules) <= 1 {
		e := rules[0].Execute()
		if e != nil {
			log.Errorf("rule: \"%s\" executed, error:\n %+v ", rules[0].RuleName, e)
		}
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(rules))
	for _, r := range rules {
		rr := r
		go func() {
			e := rr.Execute()
			if e != nil {
				log.Errorf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, e)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
