package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
	"reflect"
)

type IfStmt struct {
	Expression     *Expression
	StatementList  *Statements
	ElseIfStmtList []*ElseIfStmt
	ElseStmt       *ElseStmt
	dataCtx        *context.DataContext
}

func (i *IfStmt) Evaluate(Vars map[string]interface{}) (interface{}, error, bool) {

	it, err := i.Expression.Evaluate(Vars)
	if err != nil {
		return nil, err, false
	}

	if reflect.ValueOf(it).Bool() {
		if i.StatementList == nil {
			return nil, nil, false
		} else {
			return i.StatementList.Evaluate(Vars)
		}

	} else {

		if i.ElseIfStmtList != nil {
			for _, elseIfStmt := range i.ElseIfStmtList {
				v, err := elseIfStmt.Expression.Evaluate(Vars)
				if err != nil {
					return nil, err, false
				}

				if reflect.ValueOf(v).Bool() {
					return elseIfStmt.StatementList.Evaluate(Vars)
				}
			}
		}

		if i.ElseStmt != nil {
			return i.ElseStmt.Evaluate(Vars)
		} else {
			return nil, nil, false
		}
	}
}

func (i *IfStmt) Initialize(dc *context.DataContext) {
	i.dataCtx = dc

	if i.Expression != nil {
		i.Expression.Initialize(dc)
	}
	if i.StatementList != nil {
		i.StatementList.Initialize(dc)
	}

	if i.ElseIfStmtList != nil {
		for _, elseIfStmt := range i.ElseIfStmtList {
			elseIfStmt.Initialize(dc)
		}
	}

	if i.ElseStmt != nil {
		i.ElseStmt.Initialize(dc)
	}
}

func (i *IfStmt) AcceptExpression(expr *Expression) error {
	if i.Expression == nil {
		i.Expression = expr
		return nil
	}
	return errors.New("IfStmt Expression set twice!")
}

func (i *IfStmt) AcceptStatements(stmts *Statements) error {
	if i.StatementList == nil {
		i.StatementList = stmts
		return nil
	}
	return errors.New("ifStmt's statements set twice!")
}
