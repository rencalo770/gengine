package core

import (
	"fmt"
	errors2 "gengine/internal/core/errors"
	"reflect"
	"strings"
)

func Add(ax, bx interface{}) (interface{}, error) {
	a := reflect.ValueOf(ax)
	akind := a.Kind().String()
	b := reflect.ValueOf(bx)
	bkind := b.Kind().String()

	if !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float")) || akind == "string") && !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float")) || bkind == "sting") {
		return nil, errors2.Errorf("ADD(+) can't be used between %s and %s", akind, bkind)
	}
	if akind == "string" && bkind == "string" {
		//字符串相加
		return fmt.Sprintf("%s%s", a.String(), b.String()), nil
	}
	if akind == "string" && bkind != "string" {
		return nil, errors2.Errorf("ADD(+) can't be used between %s and %s", akind, bkind)
	}

	if akind != "string" && bkind == "string" {
		return nil, errors2.Errorf("ADD(+) can't be used between %s and %s", akind, bkind)
	}

	if strings.HasPrefix(akind, "int") {
		if strings.HasPrefix(bkind, "int") {
			return  a.Int() + b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Int() + int64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Int()) + b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "uint") {
		if strings.HasPrefix(bkind, "int") {
			return  int64(a.Uint()) + b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Uint() + b.Uint(), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Uint()) + b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "float") {
		if strings.HasPrefix(bkind, "int") {
			return  a.Float() + float64(b.Int()), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Float() + float64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return a.Float() + b.Float(), nil
		}
	}
	return nil, errors2.Errorf("ADD(+) can't be used between %s and %s", akind, bkind)
}

func Sub(ax, bx interface{}) (interface{}, error) {
	a := reflect.ValueOf(ax)
	b := reflect.ValueOf(bx)
	akind := a.Kind().String()
	bkind := b.Kind().String()

	if !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float"))) && !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float"))) {
		return nil, errors2.Errorf("SUB(-) can't be used between %s and %s", akind, bkind)
	}

	if strings.HasPrefix(akind, "int") {
		if strings.HasPrefix(bkind, "int") {
			return  a.Int() - b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Int() - int64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Int()) - b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "uint") {
		if strings.HasPrefix(bkind, "int") {
			return  int64(a.Uint()) - b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Uint() - b.Uint(), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Uint()) - b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "float") {
		if strings.HasPrefix(bkind, "int") {
			return  a.Float() - float64(b.Int()), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Float() - float64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return a.Float() - b.Float(), nil
		}
	}
	return nil, errors2.Errorf("Sub(-) can't be used between %s and %s", akind, bkind)
}

func Mul(ax, bx interface{}) (interface{}, error) {
	a := reflect.ValueOf(ax)
	b := reflect.ValueOf(bx)
	akind := a.Kind().String()
	bkind := b.Kind().String()

	if !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float"))) && !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float"))) {
		return nil, errors2.Errorf("Mul(*) can't be used between %s and %s", akind, bkind)
	}

	if strings.HasPrefix(akind, "int") {
		if strings.HasPrefix(bkind, "int") {
			return  a.Int() * b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Int() * int64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Int()) * b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "uint") {
		if strings.HasPrefix(bkind, "int") {
			return  int64(a.Uint()) * b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Uint() * b.Uint(), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Uint()) * b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "float") {
		if strings.HasPrefix(bkind, "int") {
			return  a.Float() * float64(b.Int()), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Float() * float64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return a.Float() * b.Float(), nil
		}
	}
	return nil, errors2.Errorf("Mul(*) can't be used between %s and %s", akind, bkind)
}

func Div(ax, bx interface{}) (interface{}, error) {
	a := reflect.ValueOf(ax)
	b := reflect.ValueOf(bx)
	akind := a.Kind().String()
	bkind := b.Kind().String()

	if !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float"))) && !(strings.HasPrefix(akind, "int") || (strings.HasPrefix(akind, "uint")) || (strings.HasPrefix(akind, "float"))) {
		return nil, errors2.Errorf("DIV(/) can't be used between %s and %s", akind, bkind)
	}

	if strings.HasPrefix(bkind, "int") {
		bi := b.Int()
		if bi == 0 {
			return nil, errors2.New("DIV(/) can't be used to Div ZERO(0)!")
		}
	}
	if strings.HasPrefix(bkind, "uint") {
		bu := b.Uint()
		if bu == 0 {
			return nil, errors2.New("DIV(/) can't be used to Div ZERO(0)!")
		}
	}
	if strings.HasPrefix(bkind, "float") {
		bf := b.Float()
		if bf == 0.0 {
			return nil, errors2.New("DIV(/) can't be used to Div ZERO(0)!")
		}
	}


	if strings.HasPrefix(akind, "int") {
		if strings.HasPrefix(bkind, "int") {
			return  a.Int() / b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Int() / int64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Int()) / b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "uint") {
		if strings.HasPrefix(bkind, "int") {
			return  int64(a.Uint()) / b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Uint() / b.Uint(), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Uint()) / b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "float") {
		if strings.HasPrefix(bkind, "int") {
			return  a.Float() / float64(b.Int()), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Float() / float64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return a.Float() / b.Float(), nil
		}
	}
	return nil, errors2.Errorf("Div(/) can't be used between %s and %s", akind, bkind)
}
