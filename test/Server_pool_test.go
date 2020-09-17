package test

import (
	"fmt"
	"gengine/engine"
	"testing"
	"time"
)

const rp1  = `
rule "1" "1"
begin
sleep()
println()
end
`

const rp2  = `
rule "1" "1"
begin
sleep()
println()
end
rule "2" "2"
begin
sleep()
println()
end
`

func SleepTime()  {
	//println("睡觉")
	//time.Sleep(100 * time.Second )
}


func Test_rp1(t *testing.T){

	apis := make(map[string]interface{})
	apis["println"] = fmt.Println
	apis["sleep"] = SleepTime
	pool, e1 := engine.NewGenginePool(1, 2, 1, rp1, apis)
	if e1 != nil {
		println(fmt.Sprintf("e1: %+v", e1))
	}

	go func() {
		for   {

			data := make(map[string]interface{})
			sid := []string{"1", "2"}
			e := pool.ExecuteSelectedRulesConcurrentWithMultiInput(data, sid)
			if e != nil {
				println("execute err", fmt.Sprintf("%+v",e))
			}
			println("执行...")
			time.Sleep(1000 *time.Millisecond)
		}
	}()


	go func() {
		for  {
			if pool.IsExist("1") {
				println("exist 1...")
			}

			if pool.IsExist("2") {
				println("exist 2...")
			}
			time.Sleep(1 * time.Second)

		}
	}()




	go func() {
		time.Sleep(3 * time.Second)
		e := pool.UpdatePooledRules(rp2)
		if e!=nil {
			println("execute err", fmt.Sprintf("%+v",e))
		}
	}()



	go func() {

		time.Sleep(5 * time.Second)
		println("清空规则....")
		pool.ClearPoolRules()

	}()

	go func() {

		time.Sleep(20 * time.Second)
		println("更新规则....")
		e := pool.UpdatePooledRules(rp2)
		if e!=nil {
			println("execute err", fmt.Sprintf("%+v",e))
		}
	}()


	time.Sleep(30 *time.Second)



}

















