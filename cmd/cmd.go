package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

var p = []byte(`{
    "key1":"value1",
	"key2":"value2",
	"key3":{"key3_1":"value3_1"},
	"key4":[1,2,3]
}`)

var (
	name  = flag.String("name", "", "the name of struct")
	pkg   = flag.String("pkg", "main", "the package for generated code")
	input = flag.String("input", "", "")
)

func main() {
	var t interface{}
	json.Unmarshal(p, &t)
	fmt.Println(t)
	switch v := t.(type) {
	case map[string]interface{}:
		fmt.Println(v)
	}
}
