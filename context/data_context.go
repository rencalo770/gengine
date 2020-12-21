package context

import (
	"fmt"
	"gengine/internal/core"
	"gengine/internal/core/errors"
	"gengine/internal/define"
	"reflect"
	"strings"
	"sync"
)

type DataContext struct {
	lockVars sync.Mutex
	lockBase sync.Mutex
	base     map[string]reflect.Value
}

func NewDataContext() *DataContext {
	dc := &DataContext{
		base: make(map[string]reflect.Value),
	}
	dc.loadInnerUDF()
	return dc
}

func (dc *DataContext) loadInnerUDF() {
	strconv := &define.StrconvWrapper{}
	dc.Add("strconv", strconv)
	dc.Add("isNil", core.IsNil)
}

func (dc *DataContext) Add(key string, obj interface{}) {
	dc.lockBase.Lock()
	dc.base[key] = reflect.ValueOf(obj)
	dc.lockBase.Unlock()
}

func (dc *DataContext) Del(keys ...string) {
	if len(keys) == 0 {
		return
	}
	dc.lockBase.Lock()
	for _, key := range keys {
		delete(dc.base, key)
	}
	dc.lockBase.Unlock()
}

func (dc *DataContext) Get(key string) (reflect.Value, error) {
	dc.lockBase.Lock()
	v,ok := dc.base[key]
	dc.lockBase.Unlock()
	if ok {
		return v, nil
	} else {
		return reflect.ValueOf(nil), errors.New(fmt.Sprintf("NOT FOUND key :%s ", key))
	}
}

/**
execute the injected functions
function execute supply multi return values, but simplify ,just return one value
*/
func (dc *DataContext) ExecFunc(funcName string, parameters []reflect.Value) (reflect.Value, error) {
	dc.lockBase.Lock()
	v,ok := dc.base[funcName]
	dc.lockBase.Unlock()

	if ok {
		params := core.ParamsTypeChange(v, parameters)
		fun := v
		args := make([]reflect.Value, 0)
		for _, param := range params {
			args = append(args, param)
		}
		res := fun.Call(args)
		raw, e := core.GetRawTypeValue(res)
		if e != nil {
			return reflect.ValueOf(nil), e
		}
		return raw, nil
	} else {
		return reflect.ValueOf(nil), errors.New(fmt.Sprintf("NOT FOUND function \"%s\"", funcName))
	}
}

/**
execute the struct's functions
function execute supply multi return values, but simplify ,just return one value
*/
func (dc *DataContext) ExecMethod(methodName string, args []reflect.Value) (reflect.Value, error) {
	structAndMethod := strings.Split(methodName, ".")
	//Dimit rule
	if len(structAndMethod) != 2 {
		return reflect.ValueOf(nil), errors.New(fmt.Sprintf("Not supported call, just support struct.method call, now length is %d", len(structAndMethod)))
	}

	dc.lockBase.Lock()
	v,ok := dc.base[structAndMethod[0]]
	dc.lockBase.Unlock()

	if ok {
		res, err := core.InvokeFunction(v, structAndMethod[1], args)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		return res, nil
	}
	return reflect.ValueOf(nil), errors.New(fmt.Sprintf("Not found method: %s", methodName))
}

/**
get the value user set
*/
func (dc *DataContext) GetValue(Vars map[string]reflect.Value, variable string) (reflect.Value, error) {
	if strings.Contains(variable, ".") {
		//in dataContext
		structAndField := strings.Split(variable, ".")
		//Dimit rule
		if len(structAndField) != 2 {
			return reflect.ValueOf(nil), errors.New(fmt.Sprintf("Not supported Field, just support struct.field, now length is %d", len(structAndField)))
		}

		dc.lockBase.Lock()
		v,ok := dc.base[structAndField[0]]
		dc.lockBase.Unlock()

		if ok {
			return core.GetStructAttributeValue(v, structAndField[1])
		}

		//for return struct or struct ptr
		dc.lockVars.Lock()
		obj, ok := Vars[structAndField[0]]
		dc.lockVars.Unlock()
		if ok {
			return core.GetStructAttributeValue(obj, structAndField[1])
		}
	} else {
		//user set
		dc.lockBase.Lock()
		v,ok := dc.base[variable]
		dc.lockBase.Unlock()

		if ok {
			return v, nil
		}
		//in RuleEntity
		dc.lockVars.Lock()
		res,rok := Vars[variable]
		dc.lockVars.Unlock()
		if rok {
			return res, nil
		}

	}
	return reflect.ValueOf(nil), errors.New(fmt.Sprintf("Did not found variable : %s ", variable))
}

func (dc *DataContext) SetValue(Vars map[string]reflect.Value, variable string, newValue reflect.Value) error {
	if strings.Contains(variable, ".") {
		structAndField := strings.Split(variable, ".")
		//Dimit rule
		if len(structAndField) != 2 {
			return errors.New(fmt.Sprintf("just support struct.field, now length is %d", len(structAndField)))
		}

		dc.lockBase.Lock()
		v,ok := dc.base[structAndField[0]]
		dc.lockBase.Unlock()

		if ok {
			return core.SetAttributeValue(v, structAndField[1], newValue)
		}else {
			dc.lockVars.Lock()
			vv,vok := Vars[structAndField[0]]
			dc.lockVars.Unlock()
			if vok {
				return core.SetAttributeValue(vv, structAndField[1], newValue)
			}
		}
	} else {
		dc.lockBase.Lock()
		v, ok := dc.base[variable]
		dc.lockBase.Unlock()
		if ok {
			return dc.setSingleValue(v, variable, newValue)
		} else {
			//in RuleEntity
			dc.lockVars.Lock()
			Vars[variable] = newValue
			dc.lockVars.Unlock()
			return nil
		}
	}
	return errors.New(fmt.Sprintf("setValue not found \"%s\" error.", variable))
}

//set single value
func (dc *DataContext) setSingleValue(obj reflect.Value, fieldName string, value reflect.Value) error {

	if obj.Kind() == reflect.Ptr {
		if value.Kind() == reflect.Ptr {
			//both ptr
			//value = reflect.ValueOf(value).Elem().Interface()
		}

		objKind := obj.Elem().Kind()
		valueKind := value.Kind()
		if objKind == valueKind {
			obj.Elem().Set(value)
			return nil
		} else {
			valueKindStr := valueKind.String()
			switch objKind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if strings.HasPrefix(valueKindStr, "int") {
					obj.Elem().SetInt(value.Int())
					return nil
				}
				if strings.HasPrefix(valueKindStr, "float") {
					obj.Elem().SetInt(int64(value.Float()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "uint") {
					obj.Elem().SetInt(int64(value.Uint()))
					return nil
				}
				break
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if strings.HasPrefix(valueKindStr, "int") && value.Int() >= 0 {
					obj.Elem().SetUint(uint64(value.Int()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "float") && value.Float() >= 0 {
					obj.Elem().SetUint(uint64(value.Float()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "uint") {
					obj.Elem().SetUint(value.Uint())
					return nil
				}
				break
			case reflect.Float32, reflect.Float64:
				if strings.HasPrefix(valueKindStr, "int") {
					obj.Elem().SetFloat(float64(value.Int()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "float") {
					obj.Elem().SetFloat(value.Float())
					return nil
				}
				if strings.HasPrefix(valueKindStr, "uint") {
					obj.Elem().SetFloat(float64(value.Uint()))
					return nil
				}
				break
			}
			return errors.New(fmt.Sprintf("\"%s\" value type \"%s\" is different from \"%s\" ", fieldName, reflect.ValueOf(obj).Elem().Kind().String(), reflect.ValueOf(value).Kind()))
		}
	} else {
		return errors.New(fmt.Sprintf("\"%s\" value is unassignable!", fieldName))
	}
}

func (dc *DataContext) SetMapVarValue(Vars map[string]reflect.Value, mapVarName, mapVarStrkey, mapVarVarkey string, mapVarIntkey int64, newValue reflect.Value) error {

	//value is map or slice or array
	value, e := dc.GetValue(Vars, mapVarName)
	if e != nil {
		return e
	}
	typeName := value.Type().String()
	//typeName := reflect.TypeOf(value).String()

	ptr := false
	typeNum := -1
	keyType := ""
	valueType := ""
	inTypeName := ""
	if strings.HasPrefix(typeName, "*") {
		ptr = true
		inTypeName = strings.ReplaceAll(typeName, "*", "")
		typeName = inTypeName
	}

	if strings.HasPrefix(typeName, "map") {
		//map
		typeNum = 1
		start := strings.Index(typeName, "[")
		end := strings.Index(typeName, "]")
		keyType = typeName[start+1 : end]
		valueType = strings.Trim(typeName[end+1:], " ")
	} else if strings.HasPrefix(typeName, "[]") {
		//slice
		typeNum = 2
		valueType = strings.ReplaceAll(strings.ReplaceAll(typeName, "[]", ""), " ", "")
	} else {
		//array
		typeNum = 3
		valueType = typeName[strings.Index(typeName, "]")+1:]
	}

	if ptr {
		//single map should be ptr
		if typeNum == 1 {
			if len(mapVarVarkey) > 0 {
				key, e := dc.GetValue(Vars, mapVarVarkey)
				if e != nil {
					return e
				}
				wantedKey, e := core.GetWantedValue(key, keyType)
				if e != nil {
					return e
				}

				wantedValue, e := core.GetWantedValue(newValue, valueType)
				if e != nil {
					return e
				}
				reflect.ValueOf(value).Elem().SetMapIndex(reflect.ValueOf(wantedKey), reflect.ValueOf(wantedValue))
				return nil
			}

			if len(mapVarStrkey) > 0 {
				wantedValue, e := core.GetWantedValue(newValue, valueType)
				if e != nil {
					return e
				}
				reflect.ValueOf(value).Elem().SetMapIndex(reflect.ValueOf(mapVarStrkey), reflect.ValueOf(wantedValue))
				return nil
			}

			//int key
			wantedKey, e := core.GetWantedValue(reflect.ValueOf(mapVarIntkey), keyType)
			if e != nil {
				return e
			}
			wantedValue, e := core.GetWantedValue(newValue, valueType)
			if e != nil {
				return e
			}
			reflect.ValueOf(value).Elem().SetMapIndex(reflect.ValueOf(wantedKey), reflect.ValueOf(wantedValue))
			return nil
		}

		//slice
		if typeNum == 2 {
			if len(mapVarVarkey) > 0 {
				key, e := dc.GetValue(Vars, mapVarVarkey)
				if e != nil {
					return e
				}
				wantedValue, e := core.GetWantedValue(newValue, valueType)
				if e != nil {
					return e
				}
				reflect.ValueOf(value).Elem().Index(int(reflect.ValueOf(key).Int())).Set(reflect.ValueOf(wantedValue))
				return nil
			}

			if len(mapVarStrkey) > 0 {
				return errors.New(fmt.Sprintf("the index of array or slice should not be string, now is str \"%s\"", mapVarStrkey))
			}

			if mapVarIntkey >= 0 {
				wantedValue, e := core.GetWantedValue(newValue, valueType)
				if e != nil {
					return e
				}
				reflect.ValueOf(value).Elem().Index(int(mapVarIntkey)).Set(reflect.ValueOf(wantedValue))
				return nil
			} else {
				return errors.New("Slice or Array index  must be non-negative!")
			}
		}

		if typeNum == 3 {
			if strings.Contains(mapVarName, ".") {

				splits := strings.Split(mapVarName, ".")
				stru, e := dc.Get(splits[0])
				if e != nil {
					return e
				}

				struType := reflect.TypeOf(stru).String()
				var arr reflect.Value
				if strings.HasPrefix(struType, "*") {
					arr = reflect.ValueOf(stru).Elem().FieldByName(splits[1])
				} else {
					arr = reflect.ValueOf(stru).FieldByName(splits[1])
				}

				if len(mapVarVarkey) > 0 {
					key, e := dc.GetValue(Vars, mapVarVarkey)
					if e != nil {
						return e
					}
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					arr.Elem().Index(int(reflect.ValueOf(key).Int())).Set(reflect.ValueOf(wantedValue))
					return nil
				}

				if len(mapVarStrkey) > 0 {
					return errors.New(fmt.Sprintf("the index of array or slice should not be string, now is str\"%s\"", mapVarStrkey))
				}

				if mapVarIntkey >= 0 {
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					arr.Elem().Index(int(mapVarIntkey)).Set(reflect.ValueOf(wantedValue))
					return nil
				} else {
					return errors.New("Slice index must be bigger than 0!")
				}

			} else {

				if len(mapVarVarkey) > 0 {
					key, e := dc.GetValue(Vars, mapVarVarkey)
					if e != nil {
						return e
					}
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					reflect.ValueOf(value).Elem().Index(int(reflect.ValueOf(key).Int())).Set(reflect.ValueOf(wantedValue))
					return nil
				}

				if len(mapVarStrkey) > 0 {
					return errors.New(fmt.Sprintf("the index of array or slice should not be string, now is str \"%s\"", mapVarStrkey))
				}

				if mapVarIntkey >= 0 {
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					reflect.ValueOf(value).Elem().Index(int(mapVarIntkey)).Set(reflect.ValueOf(wantedValue))
					return nil
				} else {
					return errors.New("Slice index must be bigger than 0!")
				}
			}
		}

	} else {

		// map in pointer-struct
		if typeNum == 1 {

			if strings.Contains(mapVarName, ".") {

				splits := strings.Split(mapVarName, ".")
				stru, err := dc.Get(splits[0])
				if err != nil {
					return errors.New(fmt.Sprintf("Not Found struct :%s", stru))
				}

				if len(mapVarVarkey) > 0 {
					key, e := dc.GetValue(Vars, mapVarVarkey)
					if e != nil {
						return e
					}
					wantedKey, e := core.GetWantedValue(key, keyType)
					if e != nil {
						return e
					}

					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}
					reflect.ValueOf(value).SetMapIndex(reflect.ValueOf(wantedKey), reflect.ValueOf(wantedValue))
					return nil
				}

				if len(mapVarStrkey) > 0 {
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}
					reflect.ValueOf(value).SetMapIndex(reflect.ValueOf(mapVarStrkey), reflect.ValueOf(wantedValue))
					return nil
				}

				wantedKey, e := core.GetWantedValue(reflect.ValueOf(mapVarIntkey), keyType)
				if e != nil {
					return e
				}

				wantedValue, e := core.GetWantedValue(newValue, valueType)
				if e != nil {
					return e
				}
				reflect.ValueOf(value).SetMapIndex(reflect.ValueOf(wantedKey), reflect.ValueOf(wantedValue))
				return nil

			} else {
				return errors.New("SetMapVarValue Only support directly-Pointer-Map or Map in Pointer-struct!")
			}
		}

		//slice in pointer struct
		if typeNum == 2 {
			if strings.Contains(mapVarName, ".") {

				if len(mapVarVarkey) > 0 {
					key, e := dc.GetValue(Vars, mapVarVarkey)
					if e != nil {
						return e
					}

					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					reflect.ValueOf(value).Index(int(reflect.ValueOf(key).Int())).Set(reflect.ValueOf(wantedValue))
					return nil
				}

				if len(mapVarStrkey) > 0 {
					return errors.New(fmt.Sprintf("the index of array or slice should not be string, now is str \"%s\"", mapVarStrkey))
				}

				if mapVarIntkey >= 0 {
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					reflect.ValueOf(value).Index(int(mapVarIntkey)).Set(reflect.ValueOf(wantedValue))
					return nil
				} else {
					return errors.New("Slice index must be bigger than 0!")
				}

			} else {
				return errors.New("SetMapVarValue Only support directly-Pointer-Slice or Slice in Pointer-struct!")
			}
		}

		// array in pointer-struct
		if typeNum == 3 {
			if strings.Contains(mapVarName, ".") {

				splits := strings.Split(mapVarName, ".")
				stru, e := dc.Get(splits[0])
				if e != nil {
					return e
				}

				struType := reflect.TypeOf(stru).String()
				var arr reflect.Value
				if strings.HasPrefix(struType, "*") {
					arr = reflect.ValueOf(stru).Elem().FieldByName(splits[1])
				} else {
					return errors.New("struct with Array must be pointer-struct, or you can't change Array's value!")
				}

				if len(mapVarVarkey) > 0 {
					key, e := dc.GetValue(Vars, mapVarVarkey)
					if e != nil {
						return e
					}
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					arr.Index(int(reflect.ValueOf(key).Int())).Set(reflect.ValueOf(wantedValue))
					return nil
				}

				if len(mapVarStrkey) > 0 {
					return errors.New(fmt.Sprintf("the index of array or slice should not be string, now is str \"%s\"", mapVarStrkey))
				}

				if mapVarIntkey >= 0 {
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					arr.Index(int(mapVarIntkey)).Set(reflect.ValueOf(wantedValue))
					return nil
				} else {
					return errors.New("Slice index must be bigger than 0!")
				}

			} else {
				return errors.New("SetMapVarValue Only support directly-Pointer-Array or Array in Pointer-struct!")
			}
		}
	}
	return errors.New("SetMapVarValue Only support directly-Pointer-Map, directly-Pointer-Slice and directly-Pointer-Array or Map, Slice and Array in Pointer-Struct!")
}

func (dc *DataContext) makeArray(value interface{}) {

}
