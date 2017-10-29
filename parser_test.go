package gogen

import (
	"fmt"
	"testing"
)

var p = []byte(`{
    "key1":"value1",
	"key2":123,
	"key3":{
		"key3A":"value3A",
		"key3B": {
			"key3BA":"value3BA",
			"key3BB": {
				"key3BBA":1
			}
		}
	}
}`)

var r = []byte(`type Test struct {
	s string 
	i int 
	T struct {
		t1 string 
		t2 int 
		T1 struct {
			t1A string 
		}
	}
	B float32
}`)

type Test struct {
	s string
	i int
	T struct {
		t1 string
		t2 int
		T1 struct {
			t1A string
		}
	}
	B float32
}

func Test_jsonParser_Parse(t *testing.T) {
	pr := NewJSONParser(p)
	err := pr.Parse()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(pr.String())
}

func Test_jsonParser_isMap(t *testing.T) {

}

func TestNewJSONParser(t *testing.T) {

}
