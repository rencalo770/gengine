package base

import (
	"gengine/context"
	"gengine/core/errors"
	"reflect"
	"strings"
)

//support map or array
type MapVar struct {
	Name string  // map name
	Intkey int64  // array index
	Strkey string // map key
	Varkey string // array index or map key
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (m *MapVar) Initialize(kc *KnowledgeContext, dc *context.DataContext) {
	m.knowledgeContext = kc
	m.dataCtx = dc
}

func (m *MapVar) Evaluate(Vars map[string]interface{}) (interface{}, error) {

	value, e := m.dataCtx.GetValue(Vars, m.Name)
	if e != nil {
		return nil,e
	}
	typeName := reflect.TypeOf(value).String()

	if len(m.Strkey) > 0 { //map
		return reflect.ValueOf(value).MapIndex(reflect.ValueOf(m.Strkey)).Interface(), nil
	}

	if len(m.Varkey) > 0 {

		k, e := m.dataCtx.GetValue(Vars, m.Varkey)
		if e != nil {
			return nil, e
		}

		if strings.HasPrefix(typeName, "map") {
			return reflect.ValueOf(value).MapIndex(reflect.ValueOf(k)).Interface(), nil
		}

		if strings.HasPrefix(typeName, "["){//array or slice
			i := int(reflect.ValueOf(k).Int())
			if i < 0 {
				return nil,errors.Errorf("MapVar index must be no-negative integer!")
			}

			return reflect.ValueOf(value).Index(i).Interface(),nil
		}
	}

	//m.Intkey
	if strings.HasPrefix(typeName, "map") {
		return reflect.ValueOf(value).MapIndex(reflect.ValueOf(m.Intkey)).Interface(), nil
	}

	if strings.HasPrefix(typeName, "["){//array or slice
		i := int(m.Intkey)
		if i < 0 {
			return nil,errors.Errorf("MapVar is Array, Intkey must be no-negative integer!")
		}
		return reflect.ValueOf(value).Index(i).Interface(),nil
	}

	return nil,errors.Errorf("MapVar Evaluate error!")
}

func (m *MapVar)AcceptVariable(name string) error{
	if len(m.Name) == 0 {
		m.Name = name
		return nil
	}

	if len(m.Varkey) == 0 {
		m.Varkey = name
		return nil
	}
	return errors.Errorf("MapVar's Varkey set three times!")
}

func (m *MapVar)AcceptInteger(i64 int64)  error{
	if i64 < 0 {
		return errors.Errorf("MapVar's index must be non-negative integer!")
	}

	m.Intkey = i64
	return nil
}

func (m *MapVar)AcceptString(str string) error  {
	if len(m.Strkey) == 0 {
		m.Strkey = str
		return nil
	}
	return errors.Errorf("MapVar's Strkey set three times!")
}