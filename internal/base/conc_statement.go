package base

import (
	"fmt"
	"gengine/context"
	"gengine/internal/core/errors"
	"sync"
)

type ConcStatement struct {
	Assignments   []*Assignment
	FunctionCalls []*FunctionCall
	MethodCalls   []*MethodCall
	dataCtx       *context.DataContext
}

func (cs *ConcStatement) Initialize(dc *context.DataContext) {
	cs.dataCtx = dc

	if len(cs.Assignments) > 0 {
		for _, assignment := range cs.Assignments {
			assignment.Initialize(dc)
		}
	}

	if len(cs.FunctionCalls) > 0 {
		for _, fu := range cs.FunctionCalls {
			fu.Initialize(dc)
		}
	}

	if len(cs.MethodCalls) > 0 {
		for _, m := range cs.MethodCalls {
			m.Initialize(dc)
		}
	}

}

func (cs *ConcStatement) AcceptAssignment(assignment *Assignment) error {
	cs.Assignments = append(cs.Assignments, assignment)
	return nil
}

func (cs *ConcStatement) AcceptFunctionCall(funcCall *FunctionCall) error {
	cs.FunctionCalls = append(cs.FunctionCalls, funcCall)
	return nil
}

func (cs *ConcStatement) AcceptMethodCall(methodCall *MethodCall) error {
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

	} else if l == 1 {
		if aLen > 0 {
			return cs.Assignments[0].Evaluate(Vars)
		}

		if fLen > 0 {
			return cs.FunctionCalls[0].Evaluate(Vars)
		}

		if mLen > 0 {
			return cs.MethodCalls[0].Evaluate(Vars)
		}

	} else {

		var errLock sync.Mutex
		var eMsg []string

		var wg sync.WaitGroup
		wg.Add(l)
		for _, assign := range cs.Assignments {
			assignment := assign
			go func() {
				_, e := assignment.Evaluate(Vars)
				if e != nil {
					errLock.Lock()
					eMsg = append(eMsg, fmt.Sprintf("%+v", e))
					errLock.Unlock()
				}
				wg.Done()
			}()
		}
		for _, fu := range cs.FunctionCalls {
			fun := fu
			go func() {
				_, e := fun.Evaluate(Vars)
				if e != nil {
					errLock.Lock()
					eMsg = append(eMsg, fmt.Sprintf("%+v", e))
					errLock.Unlock()
				}
				wg.Done()
			}()
		}

		for _, me := range cs.MethodCalls {
			meth := me
			go func() {
				_, e := meth.Evaluate(Vars)
				if e != nil {
					errLock.Lock()
					eMsg = append(eMsg, fmt.Sprintf("%+v", e))
					errLock.Unlock()
				}
				wg.Done()
			}()
		}
		wg.Wait()

		if len(eMsg) > 0 {
			return nil, errors.New(fmt.Sprintf("%+v", eMsg))
		}
	}
	return nil, nil
}
