package icore

import (
	"bytes"
	"encoding/gob"
	"testing"
)

type Animal struct {
	Name string
}

func copyRuleBuilder(src , dst interface{}) error{
	buf := new(bytes.Buffer)
	if e := gob.NewEncoder(buf).Encode(src); e != nil {
		return e
	}
	return gob.NewDecoder(buf).Decode(dst)
}

func Clone(a, b interface{}) error {

	buff := new(bytes.Buffer)
	e := gob.NewEncoder(buff).Encode(a)
	e =gob.NewDecoder(buff).Decode(b)
	return e
}


func Test_copy(t *testing.T) {

	animal := &Animal{Name: "cat"}
	var B Animal

	e := Clone(animal, &B)
	if e!=nil {
		panic(e)
	}

	println(B.Name)
	B.Name = "dog"
	println(animal.Name)
	println(B.Name)
}

func Test_1_copy(t *testing.T) {

/*	a1 := []int16{1}
	copy(a1, a1[1:])
	a1 = a1[:1-1]
	println(fmt.Sprintf("%+v", a1))*/
var s []int
println(len(s))
s = append(s, 1)
println(len(s))


}



