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

func (dc *DataContext)Get(key string) interface{}{
	value, _ := dc.base.Load(key)
	return value
}

/**
execute the injected functions
function execute supply multi return values, but simplify ,just return one value
 */
func(dc *DataContext)ExecFunc(funcName string, parameters []interface{}) (interface{}, error) {
	if f, ok := dc.base.Load(funcName);ok{
		f, params := core.TypeChange(f, parameters)
		if f == nil {
			return nil, errors.Errorf("Can't find %s in DataContext[when use it, please set it before]!")
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
		return nil, errors.Errorf("no such data found in DataContext!")
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
		Vars[variable] = newValue
		return nil
	}
	return errors.New("setValue not found error.")
}

func (dc *DataContext) SetMapVarValue(Vars map[string]interface{}, mapVarName, mapVarStrkey, mapVarVarkey  string , mapVarIntkey int64, newValue interface{}) error{

	//value is map or array
	value, e := dc.GetValue(Vars, mapVarName)
	if e != nil {
		return errors.Errorf("found mapVar error")
	}

	typeName := reflect.TypeOf(value).String()

	//set the map's key-value
	if len(mapVarStrkey) > 0 {
		reflect.ValueOf(value).SetMapIndex(reflect.ValueOf(mapVarStrkey), reflect.ValueOf(newValue))
		return nil
	}

	if len(mapVarVarkey) > 0 {
		k, e := dc.GetValue(Vars, mapVarVarkey)
		if e != nil {
			return e
		}

		//set the map's key-value
		if strings.HasPrefix(typeName, "map") {
			reflect.ValueOf(value).SetMapIndex(reflect.ValueOf(k),reflect.ValueOf(newValue))
			return nil
		}

		if strings.HasPrefix(typeName, "[]"){ //slice
			i := int(reflect.ValueOf(k).Int())
			if i < 0 {
				return errors.Errorf("MapVar Array or Slice index must be no-negative integer!")
			}
			reflect.ValueOf(value).Index(i).Set(reflect.ValueOf(newValue))
			return nil
		}

		//bug XXXXXXXX
		if strings.Index(typeName, "[") == 0 && strings.Index(typeName, "]") != 1 {
			return errors.Errorf("Not support to set Array's value by index")
		}
	}


	//set the map's key-value
	if strings.HasPrefix(typeName, "map") {
		reflect.ValueOf(value).SetMapIndex(reflect.ValueOf(mapVarIntkey),reflect.ValueOf(newValue))
		return nil
	}

	if strings.HasPrefix(typeName, "[]"){ //slice
		i := int(reflect.ValueOf(mapVarIntkey).Int())
		if i < 0 {
			return errors.Errorf("MapVar Array or Slice index must be no-negative integer!")
		}
		reflect.ValueOf(value).Index(i).Set(reflect.ValueOf(newValue))
		return nil
	}

	if strings.Index(typeName, "[") == 0 && strings.Index(typeName, "]") != 1 {
		return errors.Errorf("Not support to set Array's value by index")
	}

	return errors.New("SetMapVarValue not found error.")
}

func (dc *DataContext)makeArray(value interface{})  {

}
