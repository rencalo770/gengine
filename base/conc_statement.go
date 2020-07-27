package base

import (
	"gengine/context"
	"github.com/sirupsen/logrus"
	"sync"
)

type ConcStatement struct {
	Assignments      []*Assignment
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (cs *ConcStatement) Initialize(kc *KnowledgeContext, dc *context.DataContext) {
	cs.knowledgeContext = kc
	cs.dataCtx = dc

	if len(cs.Assignments) > 0 {
		for   _, assignment :=range cs.Assignments {
			assignment.Initialize(kc, dc)
		}
	}
}

func (cs *ConcStatement) AcceptAssignment(assignment *Assignment) error {
	cs.Assignments = append(cs.Assignments, assignment)
	return nil
}

func (cs *ConcStatement) Evaluate(Vars map[string]interface{}) (interface{}, error) {

	if len(cs.Assignments) > 1 {
		var wg sync.WaitGroup
		wg.Add(len(cs.Assignments))
		for  _,assign := range cs.Assignments {
			assignment := assign
			go func() {
				_, e := assignment.Evaluate(Vars)
				if e != nil {
					logrus.Errorf("concStatement Assignment err: %+v ", e)
				}
				wg.Done()
			}()
		}
		wg.Wait()
		return nil, nil
	} else {
		if len(cs.Assignments) == 0 {
			return nil,nil
		} else {
			_, e := cs.Assignments[0].Evaluate(Vars)
			if e != nil {
				return nil, e
			}
			return nil,nil
		}
	}
}