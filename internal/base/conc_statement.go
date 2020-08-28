package base

import (
	"gengine/context"
	"github.com/sirupsen/logrus"
	"sync"
)

type ConcStatement struct {
	Assignments      []*Assignment
	FunctionCalls    []*FunctionCall
	MethodCalls      []*MethodCall
//	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (cs *ConcStatement) Initialize(dc *context.DataContext) {
	cs.dataCtx = dc

	if len(cs.Assignments) > 0 {
		for   _, assignment :=range cs.Assignments {
			assignment.Initialize(dc)
		}
	}

	if len(cs.FunctionCalls) > 0 {
		for _,fu := range cs.FunctionCalls {
			fu.Initialize(dc)
		}
	}

	if len(cs.MethodCalls) > 0{
		for  _, m := range cs.MethodCalls {
			m.Initialize(dc)
		}
	}

}

func (cs *ConcStatement) AcceptAssignment(assignment *Assignment) error {
	cs.Assignments = append(cs.Assignments, assignment)
	return nil
}

func (cs *ConcStatement)AcceptFunctionCall(funcCall *FunctionCall) error {
	cs.FunctionCalls = append(cs.FunctionCalls, funcCall)
	return nil
}

func (cs *ConcStatement)AcceptMethodCall(methodCall *MethodCall) error  {
	cs.MethodCalls = append(cs.MethodCalls, methodCall)
	return nil
}

func (cs *ConcStatement) Evaluate(Vars map[string]interface{}) (interface{}, error) {

	aLen := len(cs.Assignments)
	fLen := len(cs.FunctionCalls)
	mLen := len(cs.MethodCalls)
	l := aLen + fLen + mLen
	if l <= 0 {
		return nil, nil

	}else if l == 1 {
		if aLen > 0 {
			_, e := cs.Assignments[0].Evaluate(Vars)
			if e != nil {
				logrus.Errorf("conc block sort Assignment evaluate err: %+v ",e)
			}
			return nil, e
		}

		if fLen > 0 {
			_, e := cs.FunctionCalls[0].Evaluate(Vars)
			if e != nil {
				logrus.Errorf("conc block sort FunctionCall evaluate err: %+v ",e)
			}
			return nil, e
		}

		if mLen > 0 {
			_, e := cs.MethodCalls[0].Evaluate(Vars)
			if e != nil {
				logrus.Errorf("conc block sort FunctionCall evaluate err: %+v ",e)
			}
			return nil, e
		}

	}else {
		var wg sync.WaitGroup
		wg.Add(l)
		for _, assign :=range cs.Assignments {
			assignment := assign
			go func() {
				_, e := assignment.Evaluate(Vars)
				if e != nil {
					logrus.Errorf("concStatement Assignment err: %+v ", e)
				}
				wg.Done()
			}()
		}
		for _, fu :=range cs.FunctionCalls {
			fun := fu
			go func() {
				_, e := fun.Evaluate(Vars)
				if e != nil {
					logrus.Errorf("concStatement FunctionCall err: %+v ", e)
				}
				wg.Done()
			}()
		}

		for _, me :=range cs.MethodCalls {
			meth := me
			go func() {
				_, e := meth.Evaluate(Vars)
				if e != nil {
					logrus.Errorf("concStatement MethodCall err: %+v ", e)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
	return nil, nil
}