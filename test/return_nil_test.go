package test

import (
	"fmt"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"testing"
)

//in golang, you can't println(nil)
const nil_test_rule = `
rule "1" "test return nil" 
begin
s = live.GetStringPtr()
s1 = live.GetString()
println("-----", s1)

b = live.GetBoolPtr()
i = live.GetIntPtr()
u = live.GetUintPtr()
f = live.GetFloatPtr()
e = live.GetEverPtr()

ap = live.GetArrayPtr()
a = live.GetArray()

sp = live.GetSlicePtr()
sl = live.GetSlice()

mp = live.GetMapPtr()
m = live.GetMap()
end
`

type Live struct {
}

type Ever struct {
}

func (l *Live) GetStringPtr() *string {
	return nil
}

//can't return nil
/*func (l *Live)GetString() string {
	return nil
}
*/

func (l *Live) GetString() string {
	var s string
	return s
}

func (l *Live) GetBoolPtr() *bool {
	return nil
}

//can't return nil
/*func (l *Live)GetBool() bool {
	return nil
}*/

func (l *Live) GetIntPtr() *int {
	return nil
}

//can't return nil
/*func (l *Live)GetInt() int {
	return nil
}
*/

func (l *Live) GetUintPtr() *uint {
	return nil
}

//can't return nil
/*func (l *Live)GetUint() uint {
	return nil
}*/

func (l *Live) GetFloatPtr() *float64 {
	return nil
}

//can't return nil
/*func (l *Live)GetFloat() float64 {
	return nil
}*/

func (l *Live) GetEverPtr() *Ever {
	return nil
}

//can't return nil
/*func (l *Live)GetEver() Ever {
	return nil
}*/

func (l *Live) GetArrayPtr() *[4]int64 {
	return nil
}

//can't return nil
/*func (l *Live) GetArray() [4]int64 {
	return nil
}*/

func (l *Live) GetArray() [4]int64 {
	var x [4]int64
	return x
}

func (l *Live) GetSlicePtr() *[]int64 {
	return nil
}

func (l *Live) GetSlice() []int64 {
	return nil
}

func (l *Live) GetMapPtr() *map[int64]string {
	return nil
}

func (l *Live) GetMap() map[int64]string {
	return nil
}

func Test_return_nil(t *testing.T) {

	dataContext := context.NewDataContext()
	//inject struct
	//dataContext.Add("User", user)
	//rename and inject
	dataContext.Add("println", fmt.Println)
	dataContext.Add("live", &Live{})

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	e1 := ruleBuilder.BuildRuleFromString(nil_test_rule)
	if e1 != nil {
		panic(e1)
	}

	eng := engine.NewGengine()
	// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
	e2 := eng.Execute(ruleBuilder, true)
	if e2 != nil {
		println("err-------")
	}

}
