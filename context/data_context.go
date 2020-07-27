package context

import (
	"gengine/core"
	"gengine/core/errors"
	"gengine/define"
	"reflect"
	"strings"
	"sync"
)

type DataContext struct {
	lock  sync.Mutex
	base  sync.Map
}

func NewDataContext() *DataContext {
	dc := &DataContext{
		base: sync.Map{},
	}
	dc.loadInnerUDF()
	return dc
}

func (dc *DataContext)loadInnerUDF(){
	strconv := &define.StrconvWrapper{}
	dc.Add("strconv", strconv)
}

func (dc *DataContext)Add(key string, obj interface{})  {
	dc.base.Store(key, obj)
}

func (dc *DataContext)Get(key string) (interface{}, error){
	if value, ok := dc.base.Load(key);ok{
		return value, nil
	}else {
		return nil, errors.Errorf("not Found key :%s ", key)
	}
}

/**
execute the injected functions
function execute supply multi return values, but simplify ,just return one value
 */
func(dc *DataContext)ExecFunc(funcName string, parameters []interface{}) (interface{}, error) {
	if f, ok := dc.base.Load(funcName);ok{
		f, params := core.ParamsTypeChange(f, parameters)
		if f == nil {
			return nil, errors.Errorf("Can't find %s in DataContext[when use it, please set it before]!", funcName)
		}
		fun := reflect.ValueOf(f)
		args := make([]reflect.Value, 0)
		for _, param :=range params  {
			args = append(args, reflect.ValueOf(param))
		}
		res := fun.Call(args)
		raw, e := core.GetRawTypeValue(res)
		if e != nil {
			return nil,e
		}
		return raw, nil
	}else {
		return nil, errors.New("no such data found in DataContext!")
	}
}

/**
execute the struct's functions
function execute supply multi return values, but simplify ,just return one value
 */
func (dc *DataContext)ExecMethod(methodName string, args []interface{} ) (interface{}, error) {
	structAndMethod := strings.Split(methodName, ".")
	//Dimit rule
	if len(structAndMethod) != 2 {
		return nil,errors.Errorf("Not supported call, just support struct.method call, now length is %d", len(structAndMethod))
	}

	if struc, ok := dc.base.Load(structAndMethod[0]); ok {
		res, err := core.InvokeFunction(struc, structAndMethod[1], args)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, errors.Errorf("Not found method: %s",methodName)
}

/**
get the value user set
 */
func (dc *DataContext)GetValue(Vars map[string]interface{}, variable string)(interface{}, error){
	if strings.Contains(variable, ".") {
		//in dataContext
		structAndField := strings.Split(variable, ".")
		//Dimit rule
		if len(structAndField) != 2 {
			return nil,errors.Errorf("Not supported Field, just support struct.field , now length is %d", len(structAndField))
		}

		if obj,ok := dc.base.Load(structAndField[0]); ok{
			return core.GetStructAttributeValue(obj, structAndField[1])
		}

		//for return struct or struct ptr
		if obj,ok := Vars[structAndField[0]];ok{
			return core.GetStructAttributeValue(obj, structAndField[1])
		}
	}else {
		//user set
		if obj,ok := dc.base.Load(variable); ok{
			return obj,nil
		}
		//in RuleEntity
		return Vars[variable], nil
	}
	return nil, errors.Errorf("Did not found variable : %s ",variable)
}

func (dc *DataContext) SetValue(Vars map[string]interface{}  ,variable string, newValue interface{}) error {
	if strings.Contains(variable, ".") {
		structAndField := strings.Split(variable, ".")
		//Dimit rule
		if len(structAndField) != 2 {
			return errors.Errorf("Not supported Field, just support struct.field , now length is %d", len(structAndField))
		}

		if obj, ok := dc.base.Load(structAndField[0]);ok {
			return core.SetAttributeValue(obj, structAndField[1], newValue)
		}
	}else {
		//in RuleEntity
		dc.lock.Lock()
		Vars[variable] = newValue
		dc.lock.Unlock()
		return nil
	}
	return errors.New("setValue not found error.")
}

func (dc *DataContext) SetMapVarValue(Vars map[string]interface{}, mapVarName, mapVarStrkey, mapVarVarkey  string , mapVarIntkey int64, newValue interface{}) error{

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
		inTypeName =  strings.ReplaceAll(typeName, "*","")
		typeName = inTypeName
	}

	if strings.HasPrefix(typeName, "map") {
		//map
		typeNum = 1
		start := strings.Index(typeName, "[")
		end := strings.Index(typeName, "]")
		keyType =  typeName[start + 1 : end]
		valueType = strings.Trim(typeName[end+1:], " ")
	}else if strings.HasPrefix(typeName, "[]") {
		//slice
		typeNum = 2
		valueType = strings.ReplaceAll(strings.ReplaceAll(typeName, "[]", ""), " ", "")
	}else {
		//array
		typeNum = 3
		valueType = typeName[strings.Index(typeName, "]") + 1: ]
	}

	if ptr {
		//single map should be ptr
		if typeNum == 1{
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
				return  nil
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
			if len(mapVarVarkey) > 0  {
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
			}else {
				return errors.New("Slice or Array index  must be non-negative!")
			}
		}

		if typeNum == 3 {
			if strings.Contains(mapVarName, ".") {

				splits := strings.Split(mapVarName,".")
				stru, e := dc.Get(splits[0])
				if e != nil {
					return e
				}

				struType := reflect.TypeOf(stru).String()
				var arr reflect.Value
				if strings.HasPrefix(struType, "*") {
					arr = reflect.ValueOf(stru).Elem().FieldByName(splits[1])
				}else {
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

				if mapVarIntkey >=0 {
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					arr.Elem().Index(int(mapVarIntkey)).Set(reflect.ValueOf(wantedValue))
					return nil
				}else {
					return errors.New("Slice index must be bigger than 0!")
				}

			}else {

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

				if mapVarIntkey >=0 {
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					reflect.ValueOf(value).Elem().Index(int(mapVarIntkey)).Set(reflect.ValueOf(wantedValue))
					return nil
				}else {
					return errors.New("Slice index must be bigger than 0!")
				}
			}
		}

	}else {

		// map in pointer-struct
		if  typeNum == 1 {

			if strings.Contains(mapVarName, ".") {

				splits := strings.Split(mapVarName, ".")
				stru,err := dc.Get(splits[0])
				if err != nil {
					return errors.Errorf("Not Found struct :%s", stru)
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

			}else {
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

				if mapVarIntkey >=0 {
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					reflect.ValueOf(value).Index(int(mapVarIntkey)).Set(reflect.ValueOf(wantedValue))
					return nil
				}else {
					return errors.New("Slice index must be bigger than 0!")
				}

			}else {
				return errors.New("SetMapVarValue Only support directly-Pointer-Slice or Slice in Pointer-struct!")
			}
		}

		// array in pointer-struct
		if typeNum == 3 {
			if strings.Contains(mapVarName, ".") {

				splits := strings.Split(mapVarName,".")
				stru, e := dc.Get(splits[0])
				if e != nil {
					return e
				}

				struType := reflect.TypeOf(stru).String()
				var arr reflect.Value
				if strings.HasPrefix(struType, "*") {
					arr = reflect.ValueOf(stru).Elem().FieldByName(splits[1])
				}else {
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

				if mapVarIntkey >=0 {
					wantedValue, e := core.GetWantedValue(newValue, valueType)
					if e != nil {
						return e
					}

					arr.Index(int(mapVarIntkey)).Set(reflect.ValueOf(wantedValue))
					return nil
				}else {
					return errors.New("Slice index must be bigger than 0!")
				}

			}else {
				return errors.New("SetMapVarValue Only support directly-Pointer-Array or Array in Pointer-struct!")
			}
		}
	}
	return errors.New("SetMapVarValue Only support directly-Pointer-Map, directly-Pointer-Slice and directly-Pointer-Array or Map, Slice and Array in Pointer-Struct!")
}

func (dc *DataContext)makeArray(value interface{})  {

}
