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
	base     map[string]interface{}
}

func NewDataContext() *DataContext {
	dc := &DataContext{
		base: make(map[string]interface{}),
	}
	dc.loadInnerUDF()
	return dc
}

func (dc *DataContext) loadInnerUDF() {
	strconv := &define.StrconvWrapper{}
	dc.Add("strconv", strconv)
}

func (dc *DataContext) Add(key string, obj interface{}) {
	dc.lockBase.Lock()
	dc.base[key] = obj
	dc.lockBase.Unlock()
}

func (dc *DataContext) Get(key string) (interface{}, error) {
	dc.lockBase.Lock()
	v := dc.base[key]
	dc.lockBase.Unlock()
	if v != nil {
		return v, nil
	} else {
		return nil, errors.New(fmt.Sprintf("NOT FOUND key :%s ", key))
	}
}

// Del delete the keyes
func (dc *DataContext) Del(keys ...string) {
	if len(keys) == 0 {
		return
	}
	dc.lockBase.Lock()
	for _, k := range keys {
		delete(dc.base, k)
	}
	dc.lockBase.Unlock()
}

/**
execute the injected functions
function execute supply multi return values, but simplify ,just return one value
*/
func (dc *DataContext) ExecFunc(funcName string, parameters []interface{}) (interface{}, error) {
	dc.lockBase.Lock()
	v := dc.base[funcName]
	dc.lockBase.Unlock()

	if v != nil {
		params := core.ParamsTypeChange(v, parameters)
		fun := reflect.ValueOf(v)
		args := make([]reflect.Value, 0)
		for _, param := range params {
			args = append(args, reflect.ValueOf(param))
		}
		res := fun.Call(args)
		raw, e := core.GetRawTypeValue(res)
		if e != nil {
			return nil, e
		}
		return raw, nil
	} else {
		return nil, errors.New(fmt.Sprintf("NOT FOUND function \"%s\"", funcName))
	}
}

/**
execute the struct's functions
function execute supply multi return values, but simplify ,just return one value
*/
func (dc *DataContext) ExecMethod(methodName string, args []interface{}) (interface{}, error) {
	structAndMethod := strings.Split(methodName, ".")
	//Dimit rule
	if len(structAndMethod) != 2 {
		return nil, errors.New(fmt.Sprintf("Not supported call, just support struct.method call, now length is %d", len(structAndMethod)))
	}

	dc.lockBase.Lock()
	v := dc.base[structAndMethod[0]]
	dc.lockBase.Unlock()

	if v != nil {
		res, err := core.InvokeFunction(v, structAndMethod[1], args)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, errors.New(fmt.Sprintf("Not found method: %s", methodName))
}

/**
get the value user set
*/
func (dc *DataContext) GetValue(Vars map[string]interface{}, variable string) (interface{}, error) {
	if strings.Contains(variable, ".") {
		//in dataContext
		structAndField := strings.Split(variable, ".")
		//Dimit rule
		if len(structAndField) != 2 {
			return nil, errors.New(fmt.Sprintf("Not supported Field, just support struct.field, now length is %d", len(structAndField)))
		}

		dc.lockBase.Lock()
		v := dc.base[structAndField[0]]
		dc.lockBase.Unlock()

		if v != nil {
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
		v := dc.base[variable]
		dc.lockBase.Unlock()

		if v != nil {
			return v, nil
		}
		//in RuleEntity
		dc.lockVars.Lock()
		res := Vars[variable]
		dc.lockVars.Unlock()
		return res, nil
	}
	return nil, errors.New(fmt.Sprintf("Did not found variable : %s ", variable))
}

func (dc *DataContext) SetValue(Vars map[string]interface{}, variable string, newValue interface{}) error {
	if strings.Contains(variable, ".") {
		structAndField := strings.Split(variable, ".")
		//Dimit rule
		if len(structAndField) != 2 {
			return errors.New(fmt.Sprintf("just support struct.field, now length is %d", len(structAndField)))
		}

		dc.lockBase.Lock()
		v := dc.base[structAndField[0]]
		dc.lockBase.Unlock()

		if v != nil {
			return core.SetAttributeValue(v, structAndField[1], newValue)
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
func (dc *DataContext) setSingleValue(obj interface{}, fieldName string, value interface{}) error {

	if reflect.ValueOf(obj).Kind() == reflect.Ptr {
		if reflect.ValueOf(value).Kind() == reflect.Ptr {
			//both ptr
			value = reflect.ValueOf(value).Elem().Interface()
		}

		objKind := reflect.ValueOf(obj).Elem().Kind()
		valueKind := reflect.ValueOf(value).Kind()
		if objKind == valueKind {
			reflect.ValueOf(obj).Elem().Set(reflect.ValueOf(value))
			return nil
		} else {
			valueKindStr := valueKind.String()
			switch objKind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if strings.HasPrefix(valueKindStr, "int") {
					reflect.ValueOf(obj).Elem().SetInt(reflect.ValueOf(value).Int())
					return nil
				}
				if strings.HasPrefix(valueKindStr, "float") {
					reflect.ValueOf(obj).Elem().SetInt(int64(reflect.ValueOf(value).Float()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "uint") {
					reflect.ValueOf(obj).Elem().SetInt(int64(reflect.ValueOf(value).Uint()))
					return nil
				}
				break
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if strings.HasPrefix(valueKindStr, "int") && reflect.ValueOf(value).Int() >= 0 {
					reflect.ValueOf(obj).Elem().SetUint(uint64(reflect.ValueOf(value).Int()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "float") && reflect.ValueOf(value).Float() >= 0 {
					reflect.ValueOf(obj).Elem().SetUint(uint64(reflect.ValueOf(value).Float()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "uint") {
					reflect.ValueOf(obj).Elem().SetUint(reflect.ValueOf(value).Uint())
					return nil
				}
				break
			case reflect.Float32, reflect.Float64:
				if strings.HasPrefix(valueKindStr, "int") {
					reflect.ValueOf(obj).Elem().SetFloat(float64(reflect.ValueOf(value).Int()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "float") {
					reflect.ValueOf(obj).Elem().SetFloat(reflect.ValueOf(value).Float())
					return nil
				}
				if strings.HasPrefix(valueKindStr, "uint") {
					reflect.ValueOf(obj).Elem().SetFloat(float64(reflect.ValueOf(value).Uint()))
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

func (dc *DataContext) SetMapVarValue(Vars map[string]interface{}, mapVarName, mapVarStrkey, mapVarVarkey string, mapVarIntkey int64, newValue interface{}) error {

	//value is map or slice or array
	value, e := dc.GetValue(Vars, mapVarName)
	if e != nil {
		return e
	}

	typeName := reflect.TypeOf(value).String()

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
			wantedKey, e := core.GetWantedValue(mapVarIntkey, keyType)
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

				wantedKey, e := core.GetWantedValue(mapVarIntkey, keyType)
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
