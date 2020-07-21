package core

import (
	"gengine/core/errors"
	"reflect"
	"strings"
)

func InvokeFunction(obj interface{}, methodName string, parameters []interface{}) (interface{}, error) {
	objVal := reflect.ValueOf(obj)
	fun := objVal.MethodByName(methodName).Interface()
	//change type for base type params
	funcVal, params := ParamsTypeChange(fun, parameters)
	args := make([]reflect.Value, 0)
	for _, param :=range params  {
		args = append(args, reflect.ValueOf(param))
	}
	rs := reflect.ValueOf(funcVal).Call(args)
	raw, e := GetRawTypeValue(rs)
	if e != nil {
		return nil,e
	}
	return raw,nil
}

/**
if want to support multi return ,change this method implements
 */
func GetRawTypeValue(rs []reflect.Value) (interface{},error){
	if len(rs) == 0 {
		return nil, nil
	}else {
		switch rs[0].Kind(){
		case reflect.String:
			return rs[0].String(),nil
		case reflect.Bool:
			return rs[0].Bool(), nil
		case reflect.Int:
			return int(rs[0].Int()),nil
		case reflect.Int8:
			return int8(rs[0].Int()),nil
		case reflect.Int16:
			return int16(rs[0].Int()),nil
		case reflect.Int32:
			return int32(rs[0].Int()),nil
		case reflect.Int64:
			return rs[0].Int(), nil
		case reflect.Uint:
			return uint(rs[0].Uint()),nil
		case reflect.Uint8:
			return uint8(rs[0].Uint()),nil
		case reflect.Uint16:
			return uint16(rs[0].Uint()),nil
		case reflect.Uint32:
			return uint32(rs[0].Uint()),nil
		case reflect.Uint64:
			return rs[0].Uint(),nil
		case reflect.Float32:
			return float32(rs[0].Float()),nil
		case reflect.Float64:
			return rs[0].Float(),nil
		case reflect.Struct:
			return rs[0].Interface(),nil
		case reflect.Ptr:
			newPtr := reflect.New(rs[0].Elem().Type())
			newPtr.Elem().Set(rs[0].Elem())
			return newPtr.Interface(), nil
		case reflect.Slice, reflect.Map, reflect.Array:
			return rs[0].Interface(),nil
		default:
			return nil, errors.Errorf("Can't be handled type: %s", rs[0].Kind().String())
		}
	}
}

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
		newPtr := reflect.New(v.Elem().Type())
		newPtr.Elem().Set(v.Elem())
		return newPtr.Interface()
	case reflect.Struct:
		if v.CanInterface() {
			return v.Interface()
		}
		return nil
	default:
		return nil
	}
}

func GetStructAttributeValue(obj interface{}, fieldName string) (interface{}, error) {
	stru := reflect.ValueOf(obj)
	var attrVal reflect.Value
	if stru.Kind() == reflect.Ptr {
		attrVal = stru.Elem().FieldByName(fieldName)
	} else {
		attrVal = stru.FieldByName(fieldName)
	}
	return ValueToInterface(attrVal), nil
}

/**
set field value
 */
func SetAttributeValue(obj interface{}, fieldName string, value interface{}) error {
	var field = reflect.ValueOf(nil)
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
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
		return errors.Errorf("struct has no this field: %s", fieldName)
	}

	if field.CanSet() {
		typeName := reflect.ValueOf(value).Type().String()
		switch field.Type().Kind() {
		case reflect.String:
			field.SetString(reflect.ValueOf(value).String())
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if strings.HasPrefix(typeName, "uint") {
				field.SetInt(int64(reflect.ValueOf(value).Uint()))
				return nil
			}
			if strings.HasPrefix(typeName, "float") {
				field.SetInt(int64(reflect.ValueOf(value).Float()))
				return nil
			}
			field.SetInt(reflect.ValueOf(value).Int())
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if strings.HasPrefix(typeName, "int") && reflect.ValueOf(value).Int() >= 0{
				field.SetUint(uint64(reflect.ValueOf(value).Int()))
				return nil
			}
			if strings.HasPrefix(typeName, "float") && reflect.ValueOf(value).Float() >=0 {
				field.SetUint(uint64(reflect.ValueOf(value).Float()))
				return nil
			}
			field.SetUint(reflect.ValueOf(value).Uint())
			break
		case reflect.Float32, reflect.Float64:
			if strings.HasPrefix(typeName, "int") {
				field.SetFloat(float64(reflect.ValueOf(value).Int()))
				return nil
			}
			if strings.HasPrefix(typeName, "uint") {
				field.SetFloat(float64(reflect.ValueOf(value).Uint()))
				return nil
			}
			field.SetFloat(reflect.ValueOf(value).Float())
			break
		case reflect.Bool:
			field.SetBool(reflect.ValueOf(value).Bool())
			break
		case reflect.Slice:
			field.Set(reflect.ValueOf(value))
			break
		default:
			return errors.Errorf("%s:%s", "Not support type", field.Type().Kind().String())
		}
	} else {
		return errors.Errorf("%s:%s","can not set field type", field.Type().Kind().String())
	}
	return nil
}

/*
number type exchange
*/
func ParamsTypeChange(f interface{}, params []interface{})(interface{}, []interface{}){
	tf := reflect.TypeOf(f)
	if tf.Kind().String() == "ptr"{
		tf = tf.Elem()
	}
	plen := tf.NumIn()
	for i := 0; i < plen; i ++ {
		switch tf.In(i).Kind(){
		case reflect.Int:
			params[i] = int(reflect.ValueOf(params[i]).Int())
			break
		case reflect.Int8:
			params[i] = int8(reflect.ValueOf(params[i]).Int())
			break
		case reflect.Int16:
			params[i] = int16(reflect.ValueOf(params[i]).Int())
			break
		case reflect.Int32:
			params[i] = int32(reflect.ValueOf(params[i]).Int())
			break
		case reflect.Int64:
			params[i] = reflect.ValueOf(params[i]).Int()
			break
		case reflect.Uint:
			params[i] = uint(reflect.ValueOf(params[i]).Uint())
			break
		case reflect.Uint8:
			params[i] = uint8(reflect.ValueOf(params[i]).Uint())
			break
		case reflect.Uint16:
			params[i] = uint16(reflect.ValueOf(params[i]).Uint())
			break
		case reflect.Uint32:
			params[i] = uint32(reflect.ValueOf(params[i]).Uint())
			break
		case reflect.Uint64:
			params[i] = reflect.ValueOf(params[i]).Uint()
			break
		case reflect.Float32:
			params[i] = float32(reflect.ValueOf(params[i]).Float())
			break
		case reflect.Float64:
			params[i] = reflect.ValueOf(params[i]).Float()
			break
		default:
			continue
		}
	}
	return f,params
}


func GetWantedValue(newValue interface{}, toKind string) (interface{}, error) {
	rawKind := reflect.ValueOf(newValue).Kind().String()
	if rawKind == toKind {
		return newValue, nil
	}

	switch toKind {
	case "int":
		return int(reflect.ValueOf(newValue).Int()), nil
	case "int8":
		return int8(reflect.ValueOf(newValue).Int()), nil
	case "int16":
		return int16(reflect.ValueOf(newValue).Int()), nil
	case "int32":
		return int32(reflect.ValueOf(newValue).Int()), nil
	case "int64":
		return reflect.ValueOf(newValue).Int(), nil

	case "uint":
		return uint(reflect.ValueOf(newValue).Uint()), nil
	case "uint8":
		return uint8(reflect.ValueOf(newValue).Uint()), nil
	case "uint16":
		return uint16(reflect.ValueOf(newValue).Uint()), nil
	case "uint32":
		return uint32(reflect.ValueOf(newValue).Uint()), nil
	case "uint64":
		return reflect.ValueOf(newValue).Uint(), nil

	case "float32":
		return float32(reflect.ValueOf(newValue).Float()), nil
	case "float64":
		return reflect.ValueOf(newValue).Float(), nil
	}

	return newValue, nil
}
