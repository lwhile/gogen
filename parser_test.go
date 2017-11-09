package gogen

import (
	"strconv"
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
		},
		"key3C":[
			1,2,3
		]
	},
	"key4":[1,2,3],
	"key5":[
		{"key5A":123},
		{"key5A":234}
	]
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

func Test_jsonParser_isMap(t *testing.T) {

}

func TestNewJSONParser(t *testing.T) {

}

func Test_map2PairSlice(t *testing.T) {
	m := make(map[string]interface{})
	for i := 0; i < 10; i++ {
		m[strconv.Itoa(i)] = i
	}
	opairs := map2PairSlice(m)
	for index, elem := range opairs {
		if elem.key != strconv.Itoa(index) {
			t.Error()
		}
	}
}

func Test_upperFirstKey(t *testing.T) {
	s := upperFirstKey("abc")
	if s != "Abc" {
		t.Errorf("%s != %s", s, "Abc")
	}

	s = upperFirstKey("a")
	if s != "A" {
		t.Errorf("%s != %s", s, "A")
	}
}
