package core

import (
	"fmt"
	errors2 "gengine/internal/core/errors"
	"reflect"
	"strings"
)

func Add(ax, bx interface{}) (interface{}, error){
	a := reflect.ValueOf(ax)
	akind := a.Kind().String()
	b := reflect.ValueOf(bx)
	bkind := b.Kind().String()

	if !(strings.HasPrefix(akind,"int")|| (strings.HasPrefix(akind,"uint")) || (strings.HasPrefix(akind,"float")) || akind == "string") && !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float")) || bkind == "sting") {
		return nil, errors2.Errorf("ADD(+) can't be used between %s and %s", akind, bkind)
	}
	if akind == "string" && bkind == "string" {
		//字符串相加
		return fmt.Sprintf("%s%s", a.String(), b.String()), nil
	}
	if akind == "string" && bkind !="string"  {
		return nil, errors2.Errorf("ADD(+) can't be used between %s and %s", akind, bkind)
	}

	if akind != "string" && bkind == "string" {
		return nil, errors2.Errorf("ADD(+) can't be used between %s and %s", akind, bkind)
	}

	var afloat float64
	if strings.HasPrefix(akind, "int") {
		afloat = float64(a.Int())
	}
	if strings.HasPrefix(akind, "uint") {
		afloat = float64(a.Uint())
	}
	if strings.HasPrefix(akind, "float") {
		afloat = a.Float()
	}

	var bfloat float64
	if strings.HasPrefix(bkind, "int") {
		bfloat = float64(b.Int())
	}
	if strings.HasPrefix(bkind, "uint") {
		bfloat = float64(b.Uint())
	}
	if strings.HasPrefix(bkind, "float") {
		bfloat = b.Float()
	}
	return afloat + bfloat, nil
}

func Sub(ax, bx interface{}) (interface{}, error){
	a := reflect.ValueOf(ax)
	b := reflect.ValueOf(bx)
	akind := a.Kind().String()
	bkind := b.Kind().String()

	if !(strings.HasPrefix(akind,"int")|| (strings.HasPrefix(akind,"uint")) || (strings.HasPrefix(akind,"float"))) && !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float"))) {
		return nil, errors2.Errorf("SUB(-) can't be used between %s and %s", akind, bkind)
	}

	var afloat float64
	if strings.HasPrefix(akind, "int") {
		afloat = float64(a.Int())
	}
	if strings.HasPrefix(akind, "uint") {
		afloat = float64(a.Uint())
	}
	if strings.HasPrefix(akind, "float") {
		afloat = a.Float()
	}

	var bfloat float64
	if strings.HasPrefix(bkind, "int") {
		bfloat = float64(b.Int())
	}
	if strings.HasPrefix(bkind, "uint") {
		bfloat = float64(b.Uint())
	}
	if strings.HasPrefix(bkind, "float") {
		bfloat = b.Float()
	}

	return afloat - bfloat, nil
}

func Mul(ax, bx interface{}) (interface{}, error){
	a := reflect.ValueOf(ax)
	b := reflect.ValueOf(bx)
	akind := a.Kind().String()
	bkind := b.Kind().String()

	if !(strings.HasPrefix(akind,"int")|| (strings.HasPrefix(akind,"uint")) || (strings.HasPrefix(akind,"float"))) && !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float"))) {
		return nil, errors2.Errorf("Mul(*) can't be used between %s and %s", akind, bkind)
	}

	var afloat float64
	if strings.HasPrefix(akind, "int") {
		afloat = float64(a.Int())
	}
	if strings.HasPrefix(akind, "uint") {
		afloat = float64(a.Uint())
	}
	if strings.HasPrefix(akind, "float") {
		afloat = a.Float()
	}

	var bfloat float64
	if strings.HasPrefix(bkind, "int") {
		bfloat = float64(b.Int())
	}
	if strings.HasPrefix(bkind, "uint") {
		bfloat = float64(b.Uint())
	}
	if strings.HasPrefix(bkind, "float") {
		bfloat = b.Float()
	}
	return afloat * bfloat, nil
}


func Div(ax, bx interface{}) (interface{}, error) {
	a := reflect.ValueOf(ax)
	b := reflect.ValueOf(bx)
	akind := a.Kind().String()
	bkind := b.Kind().String()

	if !(strings.HasPrefix(akind,"int")|| (strings.HasPrefix(akind,"uint")) || (strings.HasPrefix(akind,"float"))) && !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float"))) {
		return nil, errors2.Errorf("DIV(/) can't be used between %s and %s", akind, bkind)
	}

	var afloat float64
	if strings.HasPrefix(akind, "int") {
		afloat = float64(a.Int())
	}
	if strings.HasPrefix(akind, "uint") {
		afloat = float64(a.Uint())
	}
	if strings.HasPrefix(akind, "float") {
		afloat = a.Float()
	}

	var bfloat float64
	if strings.HasPrefix(bkind, "int") {
		bfloat = float64(b.Int())
	}
	if strings.HasPrefix(bkind, "uint") {
		bfloat = float64(b.Uint())
	}
	if strings.HasPrefix(bkind, "float") {
		bfloat = b.Float()
	}

	if bfloat == 0 {
		return nil, errors2.New("DIV(/) can't be used to Div ZERO(0)!")
	}
	return afloat / bfloat, nil
}