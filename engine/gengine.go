package engine

import (
	"fmt"
	"gengine/builder"
	"gengine/internal/base"
	"gengine/internal/core/errors"
	"sort"
	"sync"

	"github.com/google/martian/log"
)

type Gengine struct {
	lock         sync.Mutex
	returnResult map[string]interface{}
}

func NewGengine() *Gengine {
	return &Gengine{
		returnResult: make(map[string]interface{}),
	}
}

type Stag struct {
	StopTag bool
}

func (g *Gengine) addResult(name string, returnResult interface{}) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.returnResult[name] = returnResult
}

func (g *Gengine) GetRulesResultMap() (map[string]interface{}, error) {
	return g.returnResult, nil
}

/**
sort execute model

when b is true it means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
*/
func (g *Gengine) Execute(rb *builder.RuleBuilder, b bool) error {
	if len(rb.Kc.SortRules) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	var eMsg []string
	for _, r := range rb.Kc.SortRules {
		v, err, bx := r.Execute()
		if bx {
			g.addResult(r.RuleName, v)
		}

		if err != nil {
			if b {
				eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, err))
			} else {
				return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, err))
			}
		}
	}

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
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
	if len(rb.Kc.SortRules) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	var eMsg []string
	for _, r := range rb.Kc.SortRules {
		v, err, bx := r.Execute()
		if bx {
			g.addResult(r.RuleName, v)
		}
		if err != nil {
			if b {
				eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, err))
			} else {
				return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, err))
			}
		}

		if sTag.StopTag {
			break
		}
	}

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}

	return nil
}

/*
 concurrent execute model
 in this mode, it will not consider the priority  and not consider err control
*/
func (g *Gengine) ExecuteConcurrent(rb *builder.RuleBuilder) error {
	if len(rb.Kc.RuleEntities) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	var errLock sync.Mutex
	var eMsg []string

	var wg sync.WaitGroup
	wg.Add(len(rb.Kc.RuleEntities))
	for _, r := range rb.Kc.RuleEntities {
		rr := r
		go func() {
			v, e, bx := rr.Execute()
			if bx {
				g.addResult(rr.RuleName, v)
			}
			if e != nil {
				errLock.Lock()
				eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
				errLock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}
	return nil
}

/*
 mix model to execute rules

 in this mode, it will not consider the priority，and it also concurrently to execute rules
 first to execute the most high priority rule，then concurrently to execute last rules without consider the priority
*/
func (g *Gengine) ExecuteMixModel(rb *builder.RuleBuilder) error {
	if len(rb.Kc.SortRules) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	rules := rb.Kc.SortRules
	v, e, bx := rules[0].Execute()
	if bx {
		g.addResult(rules[0].RuleName, v)
	}

	if e != nil {
		return errors.New(fmt.Sprintf("the most high priority rule: \"%s\"  executed, error:\n %+v", rules[0].RuleName, e))
	}

	var errLock sync.Mutex
	var eMsg []string

	if (len(rules) - 1) >= 1 {
		var wg sync.WaitGroup
		wg.Add(len(rules) - 1)
		for _, r := range rules[1:] {
			rr := r
			go func() {
				v, e, bx := rr.Execute()
				if bx {
					g.addResult(rr.RuleName, v)
				}
				if e != nil {
					errLock.Lock()
					eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
					errLock.Unlock()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}
	return nil
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
func (g *Gengine) ExecuteMixModelWithStopTagDirect(rb *builder.RuleBuilder, sTag *Stag) error {
	if len(rb.Kc.SortRules) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	rules := rb.Kc.SortRules
	v, e, bx := rules[0].Execute()
	if bx {
		g.addResult(rules[0].RuleName, v)
	}
	if e != nil {
		return errors.New(fmt.Sprintf("the most high priority rule: \"%s\"  executed, error:\n %+v", rules[0].RuleName, e))
	}

	var errLock sync.Mutex
	var eMsg []string

	if !sTag.StopTag {
		if (len(rules) - 1) >= 1 {
			var wg sync.WaitGroup
			wg.Add(len(rules) - 1)
			for _, r := range rules[1:] {
				rr := r
				go func() {
					v, e, bx := rr.Execute()
					if bx {
						g.addResult(rr.RuleName, v)
					}
					if e != nil {
						errLock.Lock()
						eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
						errLock.Unlock()
					}
					wg.Done()
				}()
			}
			wg.Wait()
		}
	}

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}
	return nil
}

/**
user can choose specified name rules to run with sort, and it will continue to execute the last rules,even if there rule execute error
*/
func (g *Gengine) ExecuteSelectedRules(rb *builder.RuleBuilder, names []string) error {
	if len(rb.Kc.RuleEntities) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	var rules []*base.RuleEntity
	for _, name := range names {
		if ruleEntity, ok := rb.Kc.RuleEntities[name]; ok {
			rr := ruleEntity
			rules = append(rules, rr)
		} else {
			log.Errorf("no such rule named: \"%s\"", name)
		}
	}

	if len(rules) < 1 {
		return errors.New(fmt.Sprintf("no rules have been selected, names=%+v", names))
	}

	if len(rules) >= 2 {
		sort.SliceStable(rules, func(i, j int) bool {
			return rules[i].Salience > rules[j].Salience
		})
	}

	var eMsg []string
	for _, rule := range rules {
		rr := rule
		v, e, bx := rr.Execute()
		if bx {
			g.addResult(rr.RuleName, v)
		}
		if e != nil {
			eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
		}
	}

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}
	return nil
}

/**
user can choose specified name rules to run with sort
b bool:control whether continue to execute last rules ,when a rule execute error; if b == true ,the func is same to ExecuteSelectedRules
*/
func (g *Gengine) ExecuteSelectedRulesWithControl(rb *builder.RuleBuilder, b bool, names []string) error {
	if len(rb.Kc.SortRules) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	var rules []*base.RuleEntity
	for _, name := range names {
		if ruleEntity, ok := rb.Kc.RuleEntities[name]; ok {
			rr := ruleEntity
			rules = append(rules, rr)
		} else {
			log.Errorf("no such rule named: \"%s\"", name)
		}
	}

	if len(rules) < 1 {
		return errors.New(fmt.Sprintf("no rule has been selected, names=%+v", names))
	}

	if len(rules) >= 2 {
		sort.SliceStable(rules, func(i, j int) bool {
			return rules[i].Salience > rules[j].Salience
		})
	}

	var eMsg []string
	for _, rule := range rules {
		rr := rule
		v, e, bx := rr.Execute()
		if bx {
			g.addResult(rr.RuleName, v)
		}
		if e != nil {
			if b {
				eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
			} else {
				return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
			}
		}
	}

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}
	return nil
}

/**
user can choose specified name rules to run with sort
b bool:control whether continue to execute last rules ,when a rule execute error; if b == true ,the func is same to ExecuteSelectedRules
*/
func (g *Gengine) ExecuteSelectedRulesWithControlAndStopTag(rb *builder.RuleBuilder, b bool, sTag *Stag, names []string) error {
	if len(rb.Kc.SortRules) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	var rules []*base.RuleEntity
	for _, name := range names {
		if ruleEntity, ok := rb.Kc.RuleEntities[name]; ok {
			rr := ruleEntity
			rules = append(rules, rr)
		} else {
			log.Errorf("no such rule named: \"%s\"", name)
		}
	}

	if len(rules) < 1 {
		return errors.New(fmt.Sprintf("no rule has been selected, names=%+v", names))
	}

	if len(rules) >= 2 {
		sort.SliceStable(rules, func(i, j int) bool {
			return rules[i].Salience > rules[j].Salience
		})
	}

	var eMsg []string
	for _, rule := range rules {
		rr := rule
		v, e, bx := rr.Execute()
		if bx {
			g.addResult(rr.RuleName, v)
		}
		if e != nil {
			if b {
				eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
			} else {
				return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
			}
		}

		if sTag.StopTag {
			break
		}
	}

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}
	return nil
}

/**
user can choose specified name rules to concurrent run
*/
func (g *Gengine) ExecuteSelectedRulesConcurrent(rb *builder.RuleBuilder, names []string) error {
	if len(rb.Kc.RuleEntities) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	var rules []*base.RuleEntity
	for _, name := range names {
		if ruleEntity, ok := rb.Kc.RuleEntities[name]; ok {
			rr := ruleEntity
			rules = append(rules, rr)
		} else {
			log.Errorf("no such rule named: \"%s\"", name)
		}
	}

	if len(rules) == 0 {
		return errors.New(fmt.Sprintf("no rule has been selected, names=%+v", names))
	}

	if len(rules) == 1 {
		v, e, bx := rules[0].Execute()
		if bx {
			g.addResult(rules[0].RuleName, v)
		}
		if e != nil {
			return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rules[0].RuleName, e))
		}
		return nil
	}

	var errLock sync.Mutex
	var eMsg []string

	// len(rule) >= 2
	var wg sync.WaitGroup
	wg.Add(len(rules))
	for _, r := range rules {
		rr := r
		go func() {
			v, e, bx := rr.Execute()
			if bx {
				g.addResult(rr.RuleName, v)
			}
			if e != nil {
				errLock.Lock()
				eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
				errLock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}
	return nil
}

/**
user can choose specified name rules to run with mix model
*/
func (g *Gengine) ExecuteSelectedRulesMixModel(rb *builder.RuleBuilder, names []string) error {
	if len(rb.Kc.RuleEntities) == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	var rules []*base.RuleEntity
	for _, name := range names {
		if ruleEntity, ok := rb.Kc.RuleEntities[name]; ok {
			rr := ruleEntity
			rules = append(rules, rr)
		} else {
			log.Errorf("no such rule named: \"%s\"", name)
		}
	}

	if len(rules) == 0 {
		return errors.New(fmt.Sprintf("no rule has been selected, names=%+v", names))
	}

	if len(rules) == 1 {
		v, e, bx := rules[0].Execute()
		if bx {
			g.addResult(rules[0].RuleName, v)
		}
		if e != nil {
			return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rules[0].RuleName, e))
		}
		return nil
	}

	sort.SliceStable(rules, func(i, j int) bool {
		return rules[i].Salience > rules[j].Salience
	})

	if len(rules) == 2 {
		for _, r := range rules {
			v, err, bx := r.Execute()
			if bx {
				g.addResult(r.RuleName, v)
			}
			if err != nil {
				return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, err))
			}
		}
		return nil
	}

	// rLen >= 3
	v, e, bx := rules[0].Execute()
	if bx {
		g.addResult(rules[0].RuleName, v)
	}
	if e != nil {
		return errors.New(fmt.Sprintf("the most high priority rule: \"%s\"  executed, error:\n %+v", rules[0].RuleName, e))
	}

	var errLock sync.Mutex
	var eMsg []string

	var wg sync.WaitGroup
	wg.Add(len(rules) - 1)
	for _, r := range rules[1:] {
		rr := r
		go func() {
			v, e, bx := rr.Execute()
			if bx {
				g.addResult(rr.RuleName, v)
			}
			if e != nil {
				errLock.Lock()
				eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
				errLock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}
	return nil
}

//inverse mix model
func (g *Gengine) ExecuteInverseMixModel(rb *builder.RuleBuilder) error {
	rules := rb.Kc.SortRules
	length := len(rules)
	if length == 0 {
		return errors.New("no rule has been injected into engine.")
	}

	if length <= 2 {
		for _, r := range rules {
			v, e, bx := r.Execute()
			if bx {
				g.addResult(r.RuleName, v)
			}
			if e != nil {
				return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, e))
			}
		}
		return nil
	}

	var errLock sync.Mutex
	var eMsg []string

	var wg sync.WaitGroup
	wg.Add(length - 1)
	for _, r := range rules[:length-1] {
		rr := r
		go func() {
			v, e, bx := rr.Execute()
			if bx {
				g.addResult(rr.RuleName, v)
			}
			if e != nil {
				errLock.Lock()
				eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
				errLock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}

	v, e, bx := rules[length-1].Execute()
	if bx {
		g.addResult(rules[length-1].RuleName, v)
	}
	return e
}

//inverse mix model with user selected
func (g *Gengine) ExecuteSelectedRulesInverseMixModel(rb *builder.RuleBuilder, names []string) error {
	var rules []*base.RuleEntity
	//choose user need!
	for _, name := range names {
		if re, ok := rb.Kc.RuleEntities[name]; ok {
			rules = append(rules, re)
		} else {
			log.Errorf("no such rule named: \"%s\"", name)
		}
	}

	length := len(rules)
	if length == 0 {
		return errors.New("no rule has been selected to execute.")
	}

	//resort
	sort.SliceStable(rules, func(i, j int) bool {
		return rules[i].Salience > rules[j].Salience
	})

	if length <= 2 {
		for _, r := range rules {
			v, e, bx := r.Execute()
			if bx {
				g.addResult(r.RuleName, v)
			}
			if e != nil {
				return errors.New(fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", r.RuleName, e))
			}
		}
		return nil
	}

	var errLock sync.Mutex
	var eMsg []string

	var wg sync.WaitGroup
	wg.Add(length - 1)
	for _, r := range rules[:length-1] {
		rr := r
		go func() {
			v, e, bx := rr.Execute()
			if bx {
				g.addResult(rr.RuleName, v)
			}
			if e != nil {
				errLock.Lock()
				eMsg = append(eMsg, fmt.Sprintf("rule: \"%s\" executed, error:\n %+v ", rr.RuleName, e))
				errLock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if len(eMsg) > 0 {
		return errors.New(fmt.Sprintf("%+v", eMsg))
	}

	v, e, bx := rules[length-1].Execute()
	if bx {
		g.addResult(rules[length-1].RuleName, v)
	}
	return e
}
