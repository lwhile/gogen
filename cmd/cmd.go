package main

import (
	"flag"
	"os"

	"io/ioutil"

	"github.com/lwhile/gogen"
	"github.com/lwhile/log"
)

var p = []byte(`{
    "key1":"value1",
	"key2":"value2",
	"key3":{"key3_1":"value3_1"},
	"key4":[1,2,3]
}`)

var (
	name  = flag.String("name", "T", "the name of struct")
	pkg   = flag.String("pkg", "main", "the package for generated code")
	input = flag.String("input", "", "")
)

func main() {
	flag.Parse()

	if *input == "" {
		log.Fatal("must specific a input file")
	}
	if *name == "" {
		log.Info("struct is a default value: T")
	}
	fp, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := gogen.NewJSONParser(*name, *input, b)
	if err = jsonParser.Parse(); err != nil {
		log.Fatal(err)
	}
	if err = jsonParser.Render(); err != nil {
		log.Fatal(err)
	}
	if err = jsonParser.Output(); err != nil {
		log.Fatal(err)
	}
}
