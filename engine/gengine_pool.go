package engine

import (
	"fmt"
	"gengine/builder"
	"gengine/context"
	"gengine/internal/core/errors"
	"github.com/google/martian/log"
	"sync"
)

// when you use NewGenginePool, you just think of it as the connection pool of mysql, the higher QPS you want to support, the more resource you need to give
type GenginePool struct {
	runningLock  sync.Mutex
	freeGengines []*gengineWrapper

	//just for check whether a rule exist
	ruleBuilder *builder.RuleBuilder

	rStr      string
	execModel int
	apis      map[string]interface{}

	additionLock     sync.Mutex
	additionGengines []*gengineWrapper
	additionNum      int64

	updateLock sync.Mutex
	version    int
	clear      bool //whether rules has been cleared ，if true it means there is no rules in gengine

	getEngineLock sync.RWMutex //just one can get this lock
}

type gengineWrapper struct {
	rulebuilder *builder.RuleBuilder
	gengine     *Gengine
	version     int
	addition    bool // when gengine resource is not enough and poollength >  minPool  and  poollength < maxPool, new gengine will be create, and it will be tagged addition=true; when poollength <  minPool it will be tagged addition=false
}

//poolLen  -> gengine pool length to init
//em       -> rule execute model: 1 sort model; 2 :concurrent model; 3 mix model
//rStr  -> rules string
//apiOuter -> which user want to add rule engine to use
// just init once!!!

// best practise：
// when the has cost-time operate in your rule or you want to support high concurrency(> 200000QPS) , please set poolMinLen bigger Appropriately
// when you use NewGenginePool,you just think of it as the connection pool of mysql, the higher QPS you want to support, the more resource you need to give
func NewGenginePool(poolMinLen, poolMaxLen int64, em int, rulesStr string, apiOuter map[string]interface{}) (*GenginePool, error) {

	if !(0 < poolMinLen && poolMinLen < poolMaxLen) {
		return nil, errors.New("pool length must be bigger than 0, and poolMaxLen must bigger than poolMinLen")
	}

	if em != 1 && em != 2 && em != 3 {
		return nil, errors.New(fmt.Sprintf("exec model must be 1 or 2 or 3), now it is %d", em))
	}


	v := 0
	fg := make([]*gengineWrapper, poolMinLen)
	for i := int64(0); i < poolMinLen; i++ {
		frb, e := makeRuleBuilder(rulesStr, apiOuter)
		if e != nil {
			return nil, e
		}
		fg[i] = &gengineWrapper{
			rulebuilder: frb,
			gengine:     NewGengine(),
			version:     v,
			addition:    false,
		}
	}

	ag := make([]*gengineWrapper, poolMaxLen-poolMinLen)
	for j := int64(0); j < poolMaxLen-poolMinLen; j++ {
		arb, e := makeRuleBuilder(rulesStr, apiOuter)
		if e != nil {
			return nil, e
		}
		ag[j] = &gengineWrapper{
			rulebuilder: arb,
			gengine:     NewGengine(),
			version:     v,
			addition:    true,
		}
	}

	srcRb, e := makeRuleBuilder(rulesStr, apiOuter)
	if e != nil {
		return nil, e
	}

	p := &GenginePool{
		ruleBuilder:      srcRb,
		rStr:             rulesStr,
		freeGengines:     fg,
		apis:             apiOuter,
		execModel:        em,
		version:          v,
		additionNum:      poolMaxLen - poolMinLen,
		additionGengines: ag,
		clear:            false,
	}
	return p, nil
}

//this could ensure make thread safety!
func makeRuleBuilder(ruleStr string, apiOuter map[string]interface{}) (*builder.RuleBuilder, error) {
	dataContext := context.NewDataContext()
	if apiOuter != nil {
		for k, v := range apiOuter {
			dataContext.Add(k, v)
		}
	}

	rb := builder.NewRuleBuilder(dataContext)
	if ruleStr != "" {
		if e := rb.BuildRuleFromString(ruleStr); e != nil {
			rb.Kc.ClearRules()
			return nil, errors.New(fmt.Sprintf("build rule from string err: %+v", e))
		}
	} else {
		log.Infof("the ruleStr is \"\"")
	}
	return rb, nil
}

// if there is no enough gengine source, no request will take a lock
func (gp *GenginePool) getGengine() (*gengineWrapper, error) {

	for {
		gp.getEngineLock.Lock()
		//check if there has enough resource in pool
		numFree := len(gp.freeGengines)
		if numFree > 0 {
			gp.runningLock.Lock()
			gw := gp.freeGengines[0]
			gp.freeGengines = gp.freeGengines[1:]
			gp.runningLock.Unlock()
			gp.getEngineLock.Unlock()
			return gw, nil
		}

		//check if there has addition resource
		numAddition := len(gp.additionGengines)
		if numAddition > 0 {
			gp.additionLock.Lock()
			gw := gp.additionGengines[0]
			gp.additionGengines = gp.additionGengines[1:]
			gp.additionLock.Unlock()
			gp.getEngineLock.Unlock()
			return gw, nil
		}

		//we can create a new gengine
		/*		if gp.additionCount < gp.additionNum {
				gp.additionCount++
				dstRb, e := makeRuleBuilder(gp.rStr, gp.apis)
				if e != nil {
					gp.additionCount--
					gp.getEngineLock.Unlock()
					return nil, e
				} else {
					gw := &gengineWrapper{
						rulebuilder: dstRb,
						gengine:     NewGengine(),
						version:     gp.version,
						addition:    true,
					}
					gp.getEngineLock.Unlock()
					return gw, nil
				}
			}*/
		gp.getEngineLock.Unlock()
	}
}

// async return gengine resource to pool,and update the rules
func (gp *GenginePool) putGengineLocked(gw *gengineWrapper) {
	//addition resource
	go func() {
		if gw.version != gp.version {
			rb, e :=  makeRuleBuilder(gp.rStr, gp.apis)
			if e != nil {
				log.Errorf("async update rules err:%+v", e)
			}else{
				gw.version = gp.version
				gw.rulebuilder = rb
			}
		}
		if gw.addition {
			gp.additionLock.Lock()
			gp.additionGengines = append(gp.additionGengines, gw)
			gp.additionLock.Unlock()
		} else {
			gp.runningLock.Lock()
			gp.freeGengines = append(gp.freeGengines, gw)
			gp.runningLock.Unlock()
		}
	}()
}

//update the rules in all engine in the pool
//update success: return nil
//update failed: return error
// this is very different from connection pool,
//connection pool just need to init once
//while gengine pool need to update the rules whenever the user want to init
func (gp *GenginePool) UpdatePooledRules(ruleStr string) error {
	//check the rules
	gp.updateLock.Lock()
	rb, e := makeRuleBuilder(ruleStr, gp.apis)
	if e != nil {
		gp.updateLock.Unlock()
		return e
	}else {
		gp.ruleBuilder = rb
		gp.version ++
		gp.rStr= ruleStr
		gp.clear = false
		gp.updateLock.Unlock()
		return nil
	}
}

//clear all rules in engine in pool
func (gp *GenginePool) ClearPoolRules() {
	gp.updateLock.Lock()
	gp.ruleBuilder = nil
	gp.rStr = ""
	gp.clear = true
	gp.updateLock.Unlock()
}

func (gp *GenginePool) SetExecModel(execModel int) error {
	gp.updateLock.Lock()
	defer gp.updateLock.Unlock()
	if execModel != 1 && execModel != 2 && execModel != 3 {
		return errors.New(fmt.Sprintf("exec model must be 1 or 2 or 3), now it is %d", execModel))
	} else {
		gp.execModel = execModel
	}
	return nil
}

//check the rule whether exist
func (gp *GenginePool) IsExist(ruleName string) bool {
	if gp.rStr == "" || gp.clear || gp.ruleBuilder == nil{
		return false
	}
	_, ok := gp.ruleBuilder.Kc.RuleEntities[ruleName]
	return ok
}

func (gp *GenginePool) prepare(reqName string, req interface{}, respName string, resp interface{}) (*gengineWrapper, error) {
	//get gengine resource
	gw, e := gp.getGengine()
	if e != nil {
		return nil, e
	}

	if reqName != "" && req != nil {
		gw.rulebuilder.Dc.Add(reqName, req)
	}

	if respName != "" && resp != nil {
		gw.rulebuilder.Dc.Add(respName, resp)
	}
	return gw, nil
}

func (gp *GenginePool) prepareWithMultiInput(data map[string]interface{}) (*gengineWrapper, error) {
	//get gengine resource
	gw, e := gp.getGengine()
	if e != nil {
		return nil, e
	}

	for k, v := range data {
		//user should not inject "" string or nil value
		if k != "" && v != nil {
			gw.rulebuilder.Dc.Add(k, v)
		} else {
			log.Errorf("you should not input null string or nil value")
		}
	}

	return gw, nil
}

//execute rules as the user set execute model when init or update
//req, it is better to be ptr, or you will not get changed data
//resp, it is better to be ptr, or you will not get changed dat
func (gp *GenginePool) ExecuteRules(reqName string, req interface{}, respName string, resp interface{}) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepare(reqName, req, respName, resp)
	if e != nil {
		return e
	}
	//release resource
	defer gp.putGengineLocked(gw)

	if gp.execModel == 1 { //sort
		// when some rule execute error ,it will continue to execute last
		e := gw.gengine.Execute(gw.rulebuilder, true)
		return e
	}

	if gp.execModel == 2 { //concurrent
		gw.gengine.ExecuteConcurrent(gw.rulebuilder)
		return nil
	}

	if gp.execModel == 3 { //mix
		gw.gengine.ExecuteMixModel(gw.rulebuilder)
		return nil
	}

	return nil
}

/**
user can input more data to use in engine

it is no difference with ExecuteRules, you just can inject more data use this api
*/
func (gp *GenginePool) ExecuteRulesWithMultiInput(data map[string]interface{}) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer gp.putGengineLocked(gw)

	if gp.execModel == 1 { //sort
		// when some rule execute error ,it will continue to execute last
		e := gw.gengine.Execute(gw.rulebuilder, true)
		return e
	}

	if gp.execModel == 2 { //concurrent
		gw.gengine.ExecuteConcurrent(gw.rulebuilder)
		return nil
	}

	if gp.execModel == 3 { //mix
		gw.gengine.ExecuteMixModel(gw.rulebuilder)
		return nil
	}

	return nil

}

/**
you can use stag to control the gengine execute rules behavior out of pool
if you know what you are doing, it will improve your rules execute performance

if you want to know more about stag, to see the note above every method in Gengine.go
*/
//req, it is better to be ptr, or you will not get changed data
//resp, it is better to be ptr, or you will not get changed data
func (gp *GenginePool) ExecuteRulesWithStopTag(reqName string, req interface{}, respName string, resp interface{}, stag *Stag) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepare(reqName, req, respName, resp)
	if e != nil {
		return e
	}
	//release resource
	defer gp.putGengineLocked(gw)

	if gp.execModel == 1 { //sort
		// when some rule execute error ,it will continue to execute last
		e := gw.gengine.ExecuteWithStopTagDirect(gw.rulebuilder, true, stag)
		return e
	}

	if gp.execModel == 2 { //concurrent
		gw.gengine.ExecuteConcurrent(gw.rulebuilder)
		return nil
	}

	if gp.execModel == 3 { //mix
		gw.gengine.ExecuteMixModelWithStopTagDirect(gw.rulebuilder, stag)
		return nil
	}

	return nil
}

func (gp *GenginePool) ExecuteRulesWithMultiInputAndStopTag(data map[string]interface{}, stag *Stag) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer gp.putGengineLocked(gw)

	if gp.execModel == 1 { //sort
		// when some rule execute error ,it will continue to execute last
		e := gw.gengine.ExecuteWithStopTagDirect(gw.rulebuilder, true, stag)
		return e
	}

	if gp.execModel == 2 { //concurrent
		gw.gengine.ExecuteConcurrent(gw.rulebuilder)
		return nil
	}

	if gp.execModel == 3 { //mix
		gw.gengine.ExecuteMixModelWithStopTagDirect(gw.rulebuilder, stag)
		return nil
	}

	return nil
}

/**
see ExecuteSelectedRules in gengine.go
*/
func (gp *GenginePool) ExecuteSelectedRulesWithMultiInput(data map[string]interface{}, names []string) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer gp.putGengineLocked(gw)

	gw.gengine.ExecuteSelectedRules(gw.rulebuilder, names)
	return nil
}

/**
see ExecuteSelectedRulesConcurrent in gengine.go
*/
func (gp *GenginePool) ExecuteSelectedRulesConcurrentWithMultiInput(data map[string]interface{}, names []string) error {

	//rules has bean cleared
	if gp.clear {
		//no data to execute rule
		return nil
	}

	gw, e := gp.prepareWithMultiInput(data)
	if e != nil {
		return e
	}
	//release resource
	defer gp.putGengineLocked(gw)

	gw.gengine.ExecuteSelectedRulesConcurrent(gw.rulebuilder, names)
	return nil
}
