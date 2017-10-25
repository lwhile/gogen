package gogen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// Parser interface
type Parser interface {
	Parse() error
	Render() error
	Output() string
}

type jsonParser struct {
	// p is json data
	b []byte
	// res is result that used as go source
	res string
	// target code file
	fp os.File

	r  interface{}
	st *Struct
}

// NewJSONParser return a parser about json
func NewJSONParser(b []byte) Parser {
	return &jsonParser{
		b: b,
	}
}

func (pr *jsonParser) Parse() error {
	err := json.Unmarshal(pr.b, &pr.r)
	if err != nil {
		return err
	}
	m, ok := pr.isMap(pr.r)
	if !ok {
		return fmt.Errorf("parse error:%v is not a map[string]interface{}", pr.r)
	}
	pr.st = &Struct{}
	for k, v := range m {
		field := Field{}
		field.Name = k
		switch v.(type) {
		case string:
			field.Type = "string"
		case int:
			field.Type = "int"
		}
		pr.st.Fields = append(pr.st.Fields, field)
		pr.st.Name = "Test"
	}
	pr.Render()
	return nil
}

func (pr *jsonParser) Render() error {
	var bf bytes.Buffer
	bf.Write([]byte(TYPE + SPACE + pr.st.Name + SPACE + STRUCT + LEFTBRACE + BR))
	for _, field := range pr.st.Fields {
		bf.Write([]byte(FOURSPACE + field.Name + SPACE + field.Type + BR))
	}
	bf.Write([]byte(RIGHTBRACE))
	fmt.Println(string(bf.Bytes()))
	return nil
}

func (pr *jsonParser) Output() string {
	return pr.res
}

func (p *jsonParser) isMap(m interface{}) (map[string]interface{}, bool) {
	v, ok := m.(map[string]interface{})
	return v, ok
}
