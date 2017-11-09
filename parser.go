package gogen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"io/ioutil"

	"strings"

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
	// data from input file
	data []byte

	// used to save json.Unmarshal result
	p interface{}

	st *Struct

	// result buffer
	bf bytes.Buffer

	structName string
	pkgName    string
	output     string
}

// NewJSONParser return a parser about json
func NewJSONParser(pkgName, structName, output string, b []byte) Parser {
	return &jsonParser{
		data:       b,
		pkgName:    pkgName,
		structName: structName,
		output:     output,
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
	pr.st = &Struct{Name: pr.structName}
	pr.st.depth = 1
	pr.parse(pr.st, m)
	log.Info("Parse struct success")
	return nil
}

func (pr *jsonParser) parse(st *Struct, m map[string]interface{}) (err error) {
	if m == nil {
		err = fmt.Errorf("map is nil")
		log.Error(err)
		return
	}

	pairs := map2PairSlice(m)
	for _, elem := range pairs {
		k, v := elem.key, elem.val
		field := Field{}
		field.Key = k
		// fmt.Println("field.Key:", field.Key)
		arrFlag := false
		if arr, ok := isArray(v); ok {
			if len(arr) > 0 {
				switch arr[len(arr)-1].(type) {
				case map[string]interface{}:
					v = arr[0]
					arrFlag = true
				}
			}
		}
		if mm, ok := isMap(v); !ok {
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
	if _, err := pr.bf.Write([]byte("package " + pr.pkgName + BR + BR)); err != nil {
		return err
	}
	err := pr.render(pr.st)
	if err != nil {
		return err
	}
	log.Info("Render json success")
	return nil
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
		if _, err = pr.bf.Write([]byte(st.spaceStr() + upperFirstKey(f.Key) + SPACE)); err != nil {
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
			if _, err = pr.bf.Write([]byte(SPACE + t)); err != nil {
				return
			}
		}
		if _, err = pr.bf.Write([]byte(fmt.Sprintf(" `json:\"%s\"`\n", f.Key))); err != nil {
			return
		}
		//pr.bf.Write([]byte(BR))
	}
	if _, err = pr.bf.Write([]byte(st.lastStr())); err != nil {
		return
	}
	return nil
}

func (pr *jsonParser) Output() (err error) {
	if err = ioutil.WriteFile(pr.output, pr.bf.Bytes(), 0666); err != nil {
		return
	}
	log.Info("Output result success")
	return nil
}

func (pr *jsonParser) String() string {
	return string(pr.bf.Bytes())
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
	case float64:
		s = "float64"
	case []interface{}:
		v := ife.([]interface{})
		if len(v) < 1 {
			s = "interface{}"
			return
		}
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

type opair struct {
	key string
	val interface{}
}

type opairs []opair

func (o opairs) Len() int {
	return len(o)
}

func (o opairs) Less(i, j int) bool {
	return o[i].key <= o[j].key
}

func (o opairs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func map2PairSlice(m map[string]interface{}) opairs {
	slice := make(opairs, 0, len(m))
	for k, v := range m {
		slice = append(slice, opair{k, v})
	}
	sort.Sort(slice)
	return slice
}

func upperFirstKey(key string) string {
	r := []rune(key)
	if len(r) == 0 {
		return ""
	}
	return strings.ToUpper(string(r[0])) + string(r[1:])
}
