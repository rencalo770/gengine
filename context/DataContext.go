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