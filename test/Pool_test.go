package test

import (
	"fmt"
	"gengine/engine"
	"sync/atomic"
	"testing"
	"time"
)

type Reqest struct {
	Data float32
}

type Resp struct {
	Data uint
}


const (
	pool_rule = `
rule "测试规则名称1" "rule desc"
begin
	Req.Data = 10 + 7 + 8
   // print(Req.Data)
end`
	)

//test no rules in pool
func Test_pool_no_rules(t *testing.T) {

	t1 := time.Now()
	pool, e1 := engine.NewGenginePool(3, 5, 1, "", nil)
	if e1 != nil {
		println(fmt.Sprintf("e1: %+v", e1))
	}
	e2 := pool.ExecuteRules("", nil, "", nil)
	if e2 != nil {
		println(fmt.Sprintf("e1: %+v", e1))
	}

	println("cost time:", time.Since(t1), "ns")
}

func Test_once(t *testing.T){
	t1 := time.Now()
	apis := make(map[string]interface{})
	apis["print"] = fmt.Println
	pool, e1 := engine.NewGenginePool(5, 6, 1, pool_rule, apis)
	if e1 != nil {
		println(fmt.Sprintf("e1: %+v", e1))
	}
	println("build pool cost time:", time.Since(t1), "ns")
	reqest := &Reqest{}
	e2 := pool.ExecuteRules("Req", reqest, "", nil)
	if e2!=nil {
		panic(e2)
	}
}


func Test_pool_with_rules_for_goruntine(t *testing.T){

	max := int64(0)
	min := int64(1000000)
	cnt := int64(0)

	g1 := int64(0)
	g2 := int64(0)
	g3 := int64(0)
	g4 := int64(0)
	g5 := int64(0)

	t1 := time.Now()
	apis := make(map[string]interface{})
	apis["print"] = fmt.Println
	pool, e1 := engine.NewGenginePool(20, 26, 1, pool_rule, apis)
	if e1 != nil {
		println(fmt.Sprintf("e1: %+v", e1))
	}
	println("build pool cost time:", time.Since(t1), "ns")

	go func() {
		for {
			t2 := time.Now()
			reqest := &Reqest{Data: 1}
			e2 := pool.ExecuteRules("Req", reqest, "", nil)
			if e2 != nil {
				println(fmt.Sprintf("e1: %+v", e1))
			}
			duration := time.Since(t2)
			if int64(duration) > max {
				atomic.StoreInt64(&max, int64(duration))
			}
			if int64(duration) < min {
				atomic.StoreInt64(&min, int64(duration))
			}
			atomic.AddInt64(&cnt, 1)
			g1 ++
			//println("1 exec cost time:", , "ns\n")
		}
	}()


	/*go func() {
		for {
			t2 := time.Now()
			reqest := &Reqest{Data: 1}
			e2 := pool.ExecuteRules("Req", reqest, "", nil)
			if e2 != nil {
				println(fmt.Sprintf("e1: %+v", e1))
			}
			duration := time.Since(t2)
			if int64(duration) > max {
				atomic.StoreInt64(&max, int64(duration))
			}
			if int64(duration) < min {
				atomic.StoreInt64(&min, int64(duration))
			}
			atomic.AddInt64(&cnt, 1)
			g2 ++
			//println("2 exec cost time:", time.Since(t2), "ns\n")
		}
	}()

	go func() {
		for {
			t2 := time.Now()
			reqest := &Reqest{Data: 1}
			e2 := pool.ExecuteRules("Req", reqest, "", nil)
			if e2 != nil {
				println(fmt.Sprintf("e1: %+v", e1))
			}
			duration := time.Since(t2)
			if int64(duration) > max {
				atomic.StoreInt64(&max, int64(duration))
			}
			if int64(duration) < min {
				atomic.StoreInt64(&min, int64(duration))
			}
			atomic.AddInt64(&cnt, 1)
			g3 ++
			//println("3 exec cost time:", time.Since(t2), "ns\n")
		}
	}()

	go func() {
		for {
			t2 := time.Now()
			reqest := &Reqest{Data: 1}
			e2 := pool.ExecuteRules("Req", reqest, "", nil)
			if e2 != nil {
				println(fmt.Sprintf("e1: %+v", e1))
			}
			duration := time.Since(t2)
			if int64(duration) > max {
				atomic.StoreInt64(&max, int64(duration))
			}
			if int64(duration) < min {
				atomic.StoreInt64(&min, int64(duration))
			}
			g4 ++
			atomic.AddInt64(&cnt, 1)
			//println("4 exec cost time:", time.Since(t2), "ns\n")
		}
	}()

	go func() {
		for {
			t2 := time.Now()
			reqest := &Reqest{Data: 1}
			e2 := pool.ExecuteRules("Req", reqest, "", nil)
			if e2 != nil {
				println(fmt.Sprintf("e1: %+v", e1))
			}
			duration := time.Since(t2)
			if int64(duration) > max {
				atomic.StoreInt64(&max, int64(duration))
			}
			if int64(duration) < min {
				atomic.StoreInt64(&min, int64(duration))
			}
			atomic.AddInt64(&cnt, 1)
			g5 ++
			//println("5 exec cost time:", time.Since(t2), "ns\n")
		}
	}()
*/
	go func() {
		i := 0
		for {
			time.Sleep(1 *time.Second)
			i ++
			println("sort", i,", min: ", min , "ns, max: ", max ,"ns, cnt:" ,cnt, ", g1:", g1, ",g2:", g2, ",g3:", g3, ",g4:", g4, ",g5:", g5)
		}

	}()

	time.Sleep(10000 * time.Second)
}



