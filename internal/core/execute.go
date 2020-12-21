package core

import (
	"fmt"
	"gengine/internal/core/errors"
	"reflect"
	"strings"
)

func InvokeFunction(obj reflect.Value, methodName string, parameters []reflect.Value) (reflect.Value, error) {
	objVal := obj //reflect.ValueOf(obj)

	fun := objVal.MethodByName(methodName) //.Interface()
	/*var fun interface{}
	if !f.IsValid() {
		return nil, errors.New(fmt.Sprintf("NOT FOUND Function: %s", methodName))
	} else {
		fun = f.Interface()
	}*/
	//change type for base type params
	params := ParamsTypeChange(fun, parameters)
	/*args := make([]reflect.Value, 0)
	for _, param := range params {
		args = append(args, reflect.ValueOf(param))
	}*/
	rs := fun.Call(params)
	raw, e := GetRawTypeValue(rs)
	if e != nil {
		return reflect.ValueOf(nil), e
	}
	return raw, nil
}

/**
if want to support multi return ,change this method implements
*/
func GetRawTypeValue(rs []reflect.Value) (reflect.Value, error) {
	if len(rs) == 0 {
		return reflect.ValueOf(nil), nil
	}else {
		return rs[0], nil
	}
}
/*
func GetRawTypeValue(rs []reflect.Value) (interface{}, error) {
	if len(rs) == 0 {
		return nil, nil
	} else {
		switch rs[0].Kind() {
		case reflect.String:
			return rs[0].String(), nil
		case reflect.Bool:
			return rs[0].Bool(), nil
		case reflect.Int:
			return int(rs[0].Int()), nil
		case reflect.Int8:
			return int8(rs[0].Int()), nil
		case reflect.Int16:
			return int16(rs[0].Int()), nil
		case reflect.Int32:
			return int32(rs[0].Int()), nil
		case reflect.Int64:
			return rs[0].Int(), nil
		case reflect.Uint:
			return uint(rs[0].Uint()), nil
		case reflect.Uint8:
			return uint8(rs[0].Uint()), nil
		case reflect.Uint16:
			return uint16(rs[0].Uint()), nil
		case reflect.Uint32:
			return uint32(rs[0].Uint()), nil
		case reflect.Uint64:
			return rs[0].Uint(), nil
		case reflect.Float32:
			return float32(rs[0].Float()), nil
		case reflect.Float64:
			return rs[0].Float(), nil
		case reflect.Struct:
			return rs[0].Interface(), nil
		case reflect.Ptr:
			if rs[0].IsNil() {
				return rs[0].Interface(), nil
			}
			newPtr := reflect.New(rs[0].Elem().Type())
			newPtr.Elem().Set(rs[0].Elem())
			return newPtr.Interface(), nil
		case reflect.Slice, reflect.Map, reflect.Array:
			return rs[0].Interface(), nil
		case reflect.Interface:
			return rs[0].Interface(), nil
		case reflect.Chan:
			return rs[0].Interface(), nil
		case reflect.Complex64:
			return complex64(rs[0].Complex()), nil
		case reflect.Complex128:
			return rs[0].Complex(), nil
		case reflect.Func:
			return rs[0].Interface(), nil
		default:
			return nil, errors.New(fmt.Sprintf("Can't be handled type: %s", rs[0].Kind().String()))
		}
	}
}*/

func ValueToInterface(v reflect.Value) interface{} {
	switch v.Type().Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return v.Bool()
	case reflect.Int:
		return int(v.Int())
	case reflect.Int8:
		return int8(v.Int())
	case reflect.Int16:
		return int16(v.Int())
	case reflect.Int32:
		return int32(v.Int())
	case reflect.Int64:
		return v.Int()
	case reflect.Uint:
		return uint(v.Uint())
	case reflect.Uint8:
		return uint8(v.Uint())
	case reflect.Uint16:
		return uint16(v.Uint())
	case reflect.Uint32:
		return uint32(v.Uint())
	case reflect.Uint64:
		return v.Uint()
	case reflect.Float32:
		return float32(v.Float())
	case reflect.Float64:
		return v.Float()
	case reflect.Map:
		return v.Interface()
	case reflect.Array:
		return v.Interface()
	case reflect.Slice:
		return v.Interface()
	case reflect.Ptr:
		return v.Interface()
		//newPtr := reflect.New(v.Elem().Type())
		//newPtr.Elem().Set(v.Elem())
		//return newPtr.Interface()
	case reflect.Struct:
		if v.CanInterface() {
			return v.Interface()
		}
		return nil
	default:
		return nil
	}
}

func GetStructAttributeValue(obj reflect.Value, fieldName string) (reflect.Value, error) {
	stru := obj//reflect.ValueOf(obj)
	var attrVal reflect.Value
	if stru.Kind() == reflect.Ptr {
		attrVal = stru.Elem().FieldByName(fieldName)
	} else {
		attrVal = stru.FieldByName(fieldName)
	}
	return attrVal, nil
	//return ValueToInterface(attrVal), nil
}

/**
set field value
*/
func SetAttributeValue(obj reflect.Value, fieldName string, value reflect.Value) error {
	var field = reflect.ValueOf(nil)
	objType := obj.Type()//reflect.TypeOf(obj)
	objVal :=  obj//reflect.ValueOf(obj)
	if objType.Kind() == reflect.Ptr {
		//it points to struct
		if objType.Elem().Kind() == reflect.Struct {
			field = objVal.Elem().FieldByName(fieldName)
		}
	} else {
		//not a pointer.
		if objType.Kind() == reflect.Struct {
			field = objVal.FieldByName(fieldName)
		}
	}

	if field == reflect.ValueOf(nil) {
		return errors.New(fmt.Sprintf("struct has no this field: %s", fieldName))
	}

	if field.CanSet() {
		typeName := value.Type().String()
		switch field.Type().Kind() {
		case reflect.String:
			field.SetString(value.String())
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if strings.HasPrefix(typeName, "uint") {
				field.SetInt(int64(value.Uint()))
				return nil
			}
			if strings.HasPrefix(typeName, "float") {
				field.SetInt(int64(value.Float()))
				return nil
			}
			field.SetInt(value.Int())
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if strings.HasPrefix(typeName, "int") && value.Int() >= 0 {
				field.SetUint(uint64(value.Int()))
				return nil
			}
			if strings.HasPrefix(typeName, "float") && value.Float() >= 0 {
				field.SetUint(uint64(value.Float()))
				return nil
			}
			field.SetUint(value.Uint())
			break
		case reflect.Float32, reflect.Float64:
			if strings.HasPrefix(typeName, "int") {
				field.SetFloat(float64(value.Int()))
				return nil
			}
			if strings.HasPrefix(typeName, "uint") {
				field.SetFloat(float64(value.Uint()))
				return nil
			}
			field.SetFloat(value.Float())
			break
		case reflect.Bool:
			field.SetBool(value.Bool())
			break
		case reflect.Slice:
			field.Set(value)
			break
		case reflect.Map:
			field.Set(value)
		case reflect.Array:
			field.Set(value)
		case reflect.Struct:
			field.Set(value)
		case reflect.Interface:
			field.Set(value)
		case reflect.Chan:
			field.Set(value)
		case reflect.Complex64:
			field.SetComplex(value.Complex())
		case reflect.Complex128:
			field.SetComplex(value.Complex())
		case reflect.Func:
			field.Set(value)
		default:
			return errors.New(fmt.Sprintf("Not support type:%s", field.Type().Kind().String()))
		}
	} else {
		return errors.New(fmt.Sprintf("%s:must Be Assignable, it should be or be in addressable value!", field.Type().Kind().String()))
	}
	return nil
}

const (
	_int   = 1
	_uint  = 2
	_float = 3
)

/*
number type exchange
*/
func ParamsTypeChange(f reflect.Value, params []reflect.Value) []reflect.Value {
	tf := f.Type()
	if tf.Kind() == reflect.Ptr {
		tf = tf.Elem()
	}
	plen := tf.NumIn()
	for i := 0; i < plen; i++ {
		switch tf.In(i).Kind() {
		case reflect.Int:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int(params[i].Float()))
			}
			break
		case reflect.Int8:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int8(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int8(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int8(params[i].Float()))
			}
			break
		case reflect.Int16:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int16(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int16(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int16(params[i].Float()))
			}
			break
		case reflect.Int32:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int32(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int32(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int32(params[i].Float()))
			}
			break
		case reflect.Int64:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(params[i].Int())
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int64(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int64(params[i].Float()))
			}
			break
		case reflect.Uint:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(uint(params[i].Float()))
			}
			break
		case reflect.Uint8:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint8(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint8(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(uint8(params[i].Float()))
			}
			break
		case reflect.Uint16:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint16(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint16(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(uint16(params[i].Float()))
			}
			break
		case reflect.Uint32:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint32(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint32(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(uint32(params[i].Float()))
			}
			break
		case reflect.Uint64:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint64(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(params[i].Uint())
			} else {
				params[i] = reflect.ValueOf(uint64(params[i].Float()))
			}
			break
		case reflect.Float32:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(float32(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(float32(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(float32(params[i].Float()))
			}
			break
		case reflect.Float64:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(float64(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(float64(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(params[i].Float())
			}
			break
		case reflect.Ptr:
			break
		case reflect.Interface:
			if !reflect.ValueOf(params[i]).IsValid() {
				params[i] = reflect.New(tf.In(i))
			}
		default:
			continue
		}
	}
	return params
}

func getNumType(param reflect.Value) int {
	ts := param.Kind().String()
	if strings.HasPrefix(ts, "int") {
		return _int
	}

	if strings.HasPrefix(ts, "uint") {
		return _uint
	}

	if strings.HasPrefix(ts, "float") {
		return _float
	}

	panic(fmt.Sprintf("it is not number type, type is %s !", ts))
}

func GetWantedValue(newValue reflect.Value, toKind string) (reflect.Value, error) {
	rawKind := newValue.Kind().String()
	//rawKind := newValue.Kind().String()
	if rawKind == toKind {
		return newValue, nil
	}

	switch toKind {
	case "int":
		return reflect.ValueOf(newValue.Int()), nil
	case "int8":
		return reflect.ValueOf(int8(newValue.Int())), nil
	case "int16":
		return reflect.ValueOf(int16(newValue.Int())), nil
	case "int32":
		return reflect.ValueOf(int32(newValue.Int())), nil
	case "int64":
		return newValue, nil

	case "uint":
		return reflect.ValueOf(uint(newValue.Uint())), nil
	case "uint8":
		return reflect.ValueOf(uint8(newValue.Uint())), nil
	case "uint16":
		return reflect.ValueOf(uint16(newValue.Uint())), nil
	case "uint32":
		return reflect.ValueOf(uint32(newValue.Uint())), nil
	case "uint64":
		return newValue, nil

	case "float32":
		return reflect.ValueOf(float32(newValue.Float())), nil
	case "float64":
		return newValue, nil
	}

	return newValue, nil
}
