package base

import (
	"fmt"
	"gengine/context"
	"gengine/internal/core"
	"gengine/internal/core/errors"
	"reflect"
	"strings"
)

//support map or array
type MapVar struct {
	SourceCode
	Name    string // map name
	Intkey  int64  // array index
	Strkey  string // map key
	Varkey  string // array index or map key
	dataCtx *context.DataContext
}

func (m *MapVar) Initialize(dc *context.DataContext) {
	m.dataCtx = dc
}

func (m *MapVar) Evaluate(Vars map[string]interface{}) (interface{}, error) {

	value, e := m.dataCtx.GetValue(Vars, m.Name)
	if e != nil {
		return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", m.LineNum, m.Column, m.Code, e))
	}
	typeName := reflect.TypeOf(value).String()

	// map
	if strings.HasPrefix(typeName, "map") {

		typeStr := reflect.TypeOf(value).String()
		keyType := typeStr[strings.Index(typeStr, "[")+1 : strings.Index(typeStr, "]")]

		if len(m.Varkey) > 0 {
			key, e := m.dataCtx.GetValue(Vars, m.Varkey)
			if e != nil {
				return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", m.LineNum, m.Column, m.Code, e))
			}

			wantedKey, e := core.GetWantedValue(key, keyType)
			if e != nil {
				return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", m.LineNum, m.Column, m.Code, e))
			}
			return reflect.ValueOf(value).MapIndex(reflect.ValueOf(wantedKey)).Interface(), nil
		}

		if len(m.Strkey) > 0 {
			return reflect.ValueOf(value).MapIndex(reflect.ValueOf(m.Strkey)).Interface(), nil
		}

		//intKey
		wantedKey, e := core.GetWantedValue(m.Intkey, keyType)
		if e != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", m.LineNum, m.Column, m.Code, e))
		}

		return reflect.ValueOf(value).MapIndex(reflect.ValueOf(wantedKey)).Interface(), nil
	}

	//slice or array
	if strings.HasPrefix(typeName, "[]") || (strings.HasPrefix(typeName, "[") && strings.Index(typeName, "]") != 1) {
		if len(m.Varkey) > 0 {
			wantedKey, e := m.dataCtx.GetValue(Vars, m.Varkey)
			if e != nil {
				return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", m.LineNum, m.Column, m.Code, e))
			}
			return reflect.ValueOf(value).Index(int(reflect.ValueOf(wantedKey).Int())).Interface(), nil
		}

		if m.Intkey >= 0 {
			return reflect.ValueOf(value).Index(int(m.Intkey)).Interface(), nil
		} else {
			return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, Slice or Array index must be non-negative!", m.LineNum, m.Column, m.Code))
		}
	}

	//pointer map
	if strings.HasPrefix(typeName, "*map[") {
		typeStr := reflect.TypeOf(value).String()
		keyType := typeStr[strings.Index(typeStr, "[")+1 : strings.Index(typeStr, "]")]

		if len(m.Varkey) > 0 {
			key, e := m.dataCtx.GetValue(Vars, m.Varkey)
			if e != nil {
				return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", m.LineNum, m.Column, m.Code, e))
			}
			wantedKey, e := core.GetWantedValue(key, keyType)
			if e != nil {
				return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", m.LineNum, m.Column, m.Code, e))
			}
			return reflect.ValueOf(value).Elem().MapIndex(reflect.ValueOf(wantedKey)).Interface(), nil
		}

		if len(m.Strkey) > 0 {
			return reflect.ValueOf(value).Elem().MapIndex(reflect.ValueOf(m.Strkey)).Interface(), nil
		}

		wantedKey, e := core.GetWantedValue(m.Intkey, keyType)
		if e != nil {
			return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", m.LineNum, m.Column, m.Code, e))
		}
		return reflect.ValueOf(value).Elem().MapIndex(reflect.ValueOf(wantedKey)).Interface(), nil
	}

	//pointer slice or pointer array
	if strings.HasPrefix(typeName, "*[]") || (strings.HasPrefix(typeName, "*[") && strings.Index(typeName, "]") != 2) {

		if len(m.Varkey) > 0 {
			wantedKey, e := m.dataCtx.GetValue(Vars, m.Varkey)
			if e != nil {
				return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", m.LineNum, m.Column, m.Code, e))
			}
			return reflect.ValueOf(value).Elem().Index(int(reflect.ValueOf(wantedKey).Int())).Interface(), nil
		}

		if m.Intkey >= 0 {
			return reflect.ValueOf(value).Elem().Index(int(m.Intkey)).Interface(), nil
		} else {
			return nil, errors.New(fmt.Sprintf("line %d, column %d, code %s, Slice or Array index must be non-negative!", m.LineNum, m.Column, m.Code))
		}
	}

	return nil, errors.New(fmt.Sprintf("line %d, column %d, code: %s, Evaluate MapVarValue Only support directly-Pointer-Map, directly-Pointer-Slice and directly-Pointer-Array  or Map, Slice and Array in Pointer-Struct!", m.LineNum, m.Column, m.Code))
}

func (m *MapVar) AcceptVariable(name string) error {
	if len(m.Name) == 0 {
		m.Name = name
		return nil
	}

	if len(m.Varkey) == 0 {
		m.Varkey = name
		return nil
	}
	return errors.New("MapVar's Varkey set three times!")
}

func (m *MapVar) AcceptInteger(i64 int64) error {
	m.Intkey = i64
	return nil
}

func (m *MapVar) AcceptString(str string) error {
	if len(m.Strkey) == 0 {
		m.Strkey = str
		return nil
	}
	return errors.New("MapVar's Strkey set three times!")
}
