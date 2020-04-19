package test

import (
	"reflect"
	"strings"
	"testing"
)

type XX struct {
	X map[string]string
}

func Test_vx(t *testing.T) {

	//Slice
	var x = []int{1,2,3}
	println(reflect.TypeOf(x).String())
	reflect.ValueOf(x).Index(1).Set(reflect.ValueOf(99999))
	println(x[1])

	//Slice
	var y = []int{999,888,1919}
	println(reflect.TypeOf(y).String())
	reflect.ValueOf(y).Index(1).Set(reflect.ValueOf(100001))
	println(y[1])


	//Array
	var x1 = [3]int{1,2,3}

	println(reflect.TypeOf(x1).String())
	reflect.ValueOf(&x1).Elem().Index(1).Set(reflect.ValueOf(888))
	//x1[2] = 888
	println(x1[1])

	var ii interface{}
	ii = &x1
	reflect.ValueOf(ii).Elem().Index(2).Set(reflect.ValueOf(8080801))
	println("----------*--", x1[2])

	//reflect.ValueOf(ii).Index(2).Elem().Set(reflect.ValueOf(8889999))
	//println(reflect.ValueOf(ii).Index(2).Int())

	//Slice
	var y1 = x1[1:]
	println(reflect.TypeOf(y1).String())

	var m = make(map[interface{}]string)
	m["hello"] = "wold"
	println(reflect.TypeOf(m).String())

	xx := XX{X: make(map[string]string)}
	xx.X["hello"] = "world"
	println(reflect.TypeOf(xx.X).String())


	//var mm = make(map[string]map[string]string)
	println(reflect.ValueOf(x).Index(1).Int())
	println(reflect.ValueOf(m).MapIndex(reflect.ValueOf("hello")).String())
}

func Test_Slice(t *testing.T){

	a:= "[xxx]iniii"

	println(strings.Index(a,"["))
}
