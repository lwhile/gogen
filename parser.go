package gogen

import (
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
			field.Type = stInternal
			pr.parse(stInternal, mm)
		}
		st.Fields = append(st.Fields, field)
		// fmt.Println("st:", st.Fields)
	}
	return
}

func (pr *jsonParser) Render() error {
	// var bf bytes.Buffer
	// bf.Write([]byte(TYPE + SPACE + pr.st.Name + SPACE + STRUCT + LEFTBRACE + BR))
	// for _, field := range pr.st.Fields {
	// 	t, ok := field.Type.(string)
	// 	if !ok {
	// 		fmt.Println(field.Type)
	// 		continue
	// 	}
	// 	bf.Write([]byte(FOURSPACE + field.Key + SPACE + t + BR))
	// }
	// bf.Write([]byte(RIGHTBRACE))
	// fmt.Println(string(bf.Bytes()))
	// fmt.Println(pr.st.Fields[2].Type)
	return pr.render(pr.st)
}

func (pr *jsonParser) render(st *Struct) (err error) {
	if st == nil {
		err = fmt.Errorf("st is nil")
		log.Error(err)
		return
	}
	for _, field := range st.Fields {
		if t, ok := field.Type.(string); ok {
			fmt.Println(field.Key, ": ", t)
		} else {
			if v, ok := field.Type.(*Struct); ok {
				fmt.Print(field.Key, ":\n\t")
				pr.render(v)
			} else {
				err = fmt.Errorf("%v is not a Struct type", field.Type)
				log.Error(err)
				return
			}
		}
	}
	return nil
}

func (pr *jsonParser) Output() error {
	return nil
}

func (pr *jsonParser) String() string {
	return pr.res
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
