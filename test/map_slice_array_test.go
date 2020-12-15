package test

import (
	"fmt"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"reflect"
	"strings"

	//"strings"
	"testing"
)

//指针类型和非指针类型
func isNil(v interface{}) bool {

	//var x interface{}
	//x = nil
	println("--21----->", reflect.ValueOf(v).String())
	println("--22----->", reflect.ValueOf(v).Interface() == nil)
	return reflect.ValueOf(v).IsZero()
 //return 	reflect.ValueOf(v).Interface()
}


var myRules  = `
rule "1"
begin
a = contextInt["1"]
println(a)
//println(isNil(a))
//b = contextObj["1"]
//println(isNil(b))
end
`



func Test_map_slice_array(t *testing.T){

	type Request struct {

	}

	contextObj := make(map[string]Request)
	contextInt := make(map[string]int)

	dataContext := context.NewDataContext()
	dataContext.Add("contextInt", &contextInt)
	dataContext.Add("contextObj", &contextObj)
	dataContext.Add("println", fmt.Println)
	dataContext.Add("isNil", isNil)

	ruleBuilder := builder.NewRuleBuilder(dataContext)
	err := ruleBuilder.BuildRuleFromString(myRules)
	if err!=nil{
		panic(err)
	}

	eng := engine.NewGengine()
	err = eng.Execute(ruleBuilder, true)
	if err!=nil {
		panic(err)
	}

}


func Test_in(t *testing.T){

	type Request struct {

	}

	apis := make(map[string]interface{})

	mySlice := make(map[string]int)
	apis["mySlice"] = &mySlice

	myArray := [2]string{"1","2"}
	apis["myArray"] = &myArray

	myMap := make(map[string]int)
	myMap["hello"] = 0
	apis["myMap"] = &myMap

	myMapObj := make(map[string]*Request)
	apis["myMapObj"] = &myMapObj

	m := apis["myMap"]
	typeStr := reflect.TypeOf(m).String()
	s := typeStr[strings.Index(typeStr, "]")+1:]
	println(s)

	//mo := apis["myMapObj"]

	//println(reflect.ValueOf(mo).Elem().MapIndex(reflect.ValueOf("hello")).IsValid()) // 非指针类型, 如果为空, isZero报错，isNil报错  IsValid= false
	//println(reflect.ValueOf(mo).Elem().MapIndex(reflect.ValueOf("hello")).IsNil()) //指针类型, 如果为空, isZero报错, isNil报错 IsValid = false

	println(reflect.ValueOf(m).Elem().MapIndex(reflect.ValueOf("hello")).IsValid()) //指针类型, 如果为空, isZero报错, isNil报错 IsValid = false
	println(reflect.ValueOf(m).Elem().MapIndex(reflect.ValueOf("aaa")).IsValid())

	println(reflect.TypeOf(m).Elem().String())

	//value := reflect.MakeMapWithSize(reflect.TypeOf(m).Elem(), 1)
	//println(reflect.ValueOf(value).Type().String())
	//println(reflect.ValueOf(m).Elem().MapIndex(reflect.ValueOf("hello1")).Type().String())

	//refmap.MapIndex(reflect.Zero(refmap.Type().Key()))


	fmt.Println(reflect.TypeOf(mySlice).Elem())
}

func Test_value_type(t *testing.T)  {

	apis := make(map[string]interface{})

	myMap := make(map[string]int)
	myMap["hello"] = 0
	apis["myMap"] = myMap

	m := apis["myMap"]

	println(reflect.TypeOf(m).Key().String())

	//println(reflect.TypeOf(m).Elem().Key().String())
	//println(reflect.TypeOf(m).Elem().Elem().String())




}