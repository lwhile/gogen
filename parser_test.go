package gogen

import (
	"testing"
)

var p = []byte(`{
    "key1":"value1",
	"key2":"value2"
}`)

func Test_jsonParser_Parse(t *testing.T) {
	pr := NewJSONParser(p)
	err := pr.Parse()
	if err != nil {
		t.Error(err)
	}
}

func Test_jsonParser_isMap(t *testing.T) {

}

func TestNewJSONParser(t *testing.T) {

}
