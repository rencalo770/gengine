package core

import (
	"gengine/core/errors"
	"reflect"
)

func InvokeFunction(obj interface{}, methodName string, parameters []interface{}) (interface{}, error) {
	objVal := reflect.ValueOf(obj)
	fun := objVal.MethodByName(methodName).Interface()
	//change type for base type params
	funcVal, params := TypeChange(fun, parameters)
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
		switch rs[0].Kind().String(){
		case "string":
			return rs[0].String(),nil
		case "bool":
			return rs[0].Bool(), nil
		case "int":
			return int(rs[0].Int()),nil
		case "int8":
			return int8(rs[0].Int()),nil
		case "int16":
			return int16(rs[0].Int()),nil
		case "int32":
			return int32(rs[0].Int()),nil
		case "int64":
			return int64(rs[0].Int()),nil
		case "uint":
			return uint(rs[0].Uint()),nil
		case "uint8":
			return uint8(rs[0].Uint()),nil
		case "uint16":
			return uint16(rs[0].Uint()),nil
		case "uint32":
			return uint32(rs[0].Uint()),nil
		case "uint64":
			return rs[0].Uint(),nil
		case "float32":
			return float32(rs[0].Float()),nil
		case "float64":
			return rs[0].Float(),nil
		case "struct":
			return rs[0].Interface(),nil
		case "ptr":
			newPtr := reflect.New(rs[0].Elem().Type())
			newPtr.Elem().Set(rs[0].Elem())
			return newPtr.Interface(), nil
		case "slice":
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
设置属性值
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
		switch field.Type().Kind() {
		case reflect.String:
			field.SetString(reflect.ValueOf(value).String())
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(int64(reflect.ValueOf(value).Float()))
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(uint64(reflect.ValueOf(value).Float()))
			break
		case reflect.Float32, reflect.Float64:
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
func TypeChange(f interface{}, params []interface{})(interface{}, []interface{}){
	tf := reflect.TypeOf(f)
	if tf.Kind().String() == "ptr"{
		tf = tf.Elem()
	}
	plen := tf.NumIn()
	for i := 0; i < plen; i ++ {
		switch tf.In(i).Kind().String(){
		case "int":
			params[i] = int(reflect.ValueOf(params[i]).Int())
			break
		case "int8":
			params[i] = int8(reflect.ValueOf(params[i]).Int())
			break
		case "int16":
			params[i] = int16(reflect.ValueOf(params[i]).Int())
			break
		case "int32":
			params[i] = int32(reflect.ValueOf(params[i]).Int())
			break
		case "int64":
			params[i] = reflect.ValueOf(params[i]).Int()
			break
		case "uint":
			params[i] = uint(reflect.ValueOf(params[i]).Uint())
			break
		case "uint8":
			params[i] = uint8(reflect.ValueOf(params[i]).Uint())
			break
		case "uint16":
			params[i] = uint16(reflect.ValueOf(params[i]).Uint())
			break
		case "uint32":
			params[i] = uint32(reflect.ValueOf(params[i]).Uint())
			break
		case "uint64":
			params[i] = reflect.ValueOf(params[i]).Uint()
			break
		case "float32":
			params[i] = float32(reflect.ValueOf(params[i]).Float())
			break
		case "float64":
			params[i] = reflect.ValueOf(params[i]).Float()
			break
		default:
			continue
		}
	}
	return f,params
}
