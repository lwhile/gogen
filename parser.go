package gogen

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"

	"github.com/lwhile/log"
)

// Parser interface
type Parser interface {
	Parse() error
	Render() error
	Output() error
	String() string
}

type jsonParser struct {
	name string

	// data from input file
	data []byte

	savePath string

	// used to save json.Unmarshal result
	p interface{}

	st *Struct

	// result buffer
	bf bytes.Buffer

	result []byte
}

// NewJSONParser return a parser about json
func NewJSONParser(name, path string, b []byte) Parser {
	return &jsonParser{
		data:     b,
		name:     name,
		savePath: path,
	}
}

func (pr *jsonParser) Parse() error {
	err := json.Unmarshal(pr.data, &pr.p)
	if err != nil {
		return err
	}
	m, ok := isMap(pr.p)
	if !ok {
		return fmt.Errorf("parse error:%v is not a map[string]interface{}", pr.p)
	}
	pr.st = &Struct{Name: pr.name}
	pr.st.depth = 1
	pr.parse(pr.st, m)
	pr.Render()
	return nil
}

func (pr *jsonParser) parse(st *Struct, m map[string]interface{}) (err error) {
	if m == nil {
		err = fmt.Errorf("map is nil")
		log.Error(err)
		return
	}
	for k, v := range m {
		field := Field{}
		field.Key = k
		// fmt.Println("field.Key:", field.Key)
		arrFlag := false
		if arr, ok := isArray(v); ok {
			switch arr[len(arr)-1].(type) {
			case map[string]interface{}:
				v = arr[0]
				arrFlag = true
			}
		}
		if mm, ok := isMap(v); !ok {
			// if _, ok := isArray(v); ok {
			// 	field.array = true
			// }
			// if v, ok := isArray(v); ok {
			// 	if _, ok := v.([]Struct); ok {

			// 	}
			// }
			field.Type = typeStr(v)
		} else {
			stInternal := &Struct{}
			stInternal.Fields = make([]Field, 0)
			stInternal.nesting = true
			stInternal.depth = st.depth + 1
			if arrFlag {
				stInternal.isArray = true
			}
			field.Type = stInternal
			pr.parse(stInternal, mm)
		}
		st.Fields = append(st.Fields, field)
		// fmt.Println("st:", st.Fields)
	}
	return
}

func (pr *jsonParser) Render() error {
	return pr.render(pr.st)
}

func (pr *jsonParser) render(st *Struct) (err error) {
	if st == nil {
		err = fmt.Errorf("error to serialize a nil struct to []byte")
		log.Error(err)
		return
	}

	if !st.nesting {
		if _, err = pr.bf.Write([]byte(st.firstStr())); err != nil {
			return
		}
	} else {
		var arrStr string
		if st.isArray {
			arrStr = "[]"
		}
		if _, err = pr.bf.Write([]byte(" " + arrStr + STRUCT + LEFTBRACE + BR)); err != nil {
			return
		}
	}

	for _, f := range st.Fields {
		if _, err = pr.bf.Write([]byte(st.spaceStr() + f.Key + SPACE)); err != nil {
			return
		}
		if v, ok := f.Type.(*Struct); ok {
			pr.render(v)
		} else {
			t, ok := f.Type.(string)
			if !ok {
				err = fmt.Errorf("%v not a string type", f.Type)
				return
			}
			if _, err = pr.bf.Write([]byte(SPACE + t + BR)); err != nil {
				return
			}
		}
		//pr.bf.Write([]byte(BR))
	}
	if _, err = pr.bf.Write([]byte(st.lastStr())); err != nil {
		return
	}
	pr.result = pr.bf.Bytes()
	pr.bf.Reset()
	return nil
}

func (pr *jsonParser) Output() (err error) {
	if pr.result == nil {
		err = fmt.Errorf("parse result is nil")
		log.Error(err)
		return
	}
	if err = ioutil.WriteFile(pr.savePath, pr.result, 0666); err != nil {
		return
	}
	return nil
}

func (pr *jsonParser) String() string {
	return string(pr.result)
}

func isMap(m interface{}) (map[string]interface{}, bool) {
	v, ok := m.(map[string]interface{})
	return v, ok
}

func isArray(a interface{}) ([]interface{}, bool) {
	v, ok := a.([]interface{})
	return v, ok
}

func typeStr(ife interface{}) (s string) {
	switch ife.(type) {
	case string:
		s = "string"
	case []string:
		s = "[]string"
	case int:
		s = "int"
	case int64:
		s = "int64"
	case int32:
		s = "int32"
	case float32:
		s = "float32"
	case float64:
		s = "float64"
	case []interface{}:
		v := ife.([]interface{})
		switch v[len(v)-1].(type) {
		case float64:
			s = "[]float64"
		case string:
			s = "[]string"
		case map[string]interface{}:
			s = "[]"
		}
	default:
		s = "interface{}"
	}
	return
}
