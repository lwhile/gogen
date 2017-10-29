package gogen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

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
	data []byte
	res  string
	fp   os.File

	p  interface{}
	st *Struct
	bf bytes.Buffer
}

// NewJSONParser return a parser about json
func NewJSONParser(b []byte) Parser {
	return &jsonParser{
		data: b,
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
	pr.st = &Struct{Name: "Test"}
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
		if mm, ok := isMap(v); !ok {
			field.Type = typeStr(v)
		} else {
			stInternal := &Struct{}
			stInternal.Fields = make([]Field, 0)
			stInternal.nesting = true
			stInternal.depth = st.depth + 1
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
		if _, err = pr.bf.Write([]byte(" " + STRUCT + LEFTBRACE + BR)); err != nil {
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
	return nil
}

func (pr *jsonParser) Output() error {
	return nil
}

func (pr *jsonParser) String() string {
	return string(pr.bf.Bytes())
}

func isMap(m interface{}) (map[string]interface{}, bool) {
	v, ok := m.(map[string]interface{})
	return v, ok
}

func typeStr(ife interface{}) (s string) {
	switch ife.(type) {
	case string:
		s = "string"
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
	default:
		s = "interface{}"
	}
	return
}
