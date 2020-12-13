package test

import (
	"fmt"
	"gengine/engine"
	"testing"
	"time"
)

func random() int {
	return time.Now().Nanosecond()
}

func Test_pool_return_statments(t *testing.T) {

	ruleName := "test_pool_return"
	rule := `rule "` + ruleName + `"  
			begin
				return random()
			end
			`

	apis := make(map[string]interface{})
	apis["print"] = fmt.Println
	apis["random"] = random
	pool, e1 := engine.NewGenginePool(1, 2, 1, rule, apis)
	if e1 != nil {
		println(fmt.Sprintf("e1: %+v", e1))
	}

	data := make(map[string]interface{})
	e2, rrm1 := pool.ExecuteRulesWithMultiInput(data)
	if e2 != nil {
		panic(e2)
	}

	i1 := rrm1[ruleName]
	ix1 := i1.(int)
	println("ix1--->", ix1)

	e3, rrm2 := pool.ExecuteRulesWithMultiInput(data)
	if e3 != nil {
		panic(e2)
	}

	i2 := rrm2[ruleName]
	ix2 := i2.(int)
	println("ix2--->", ix2)

	i3 := rrm1[ruleName]
	ix3 := i3.(int)
	println("ix3--->", ix3)
}
