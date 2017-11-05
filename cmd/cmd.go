package main

import (
	"flag"
	"io/ioutil"
	"os"

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
	name   = flag.String("name", "T", "the name of struct")
	pkg    = flag.String("pkg", "main", "the package for generated code")
	input  = flag.String("input", "", "location where input data")
	output = flag.String("output", "", "location where ouput data")
	web    = flag.Bool("web", false, "start web mode")
	port   = flag.String("port", "4928", "listen port")
)

const (
	defaultOutput = "./gogen_result.go"
	defaultInput  = "./input.json"
)

func helper() string {
	return ""
}

func initEnv() {
	flag.Parse()
	if *web {
		log.Info("Use web mode")
		return
	}
	// package name
	if *pkg == "" {
		*pkg = "main"
	}
	if *input == "" {
		*input = defaultInput
		log.Info("Use default input file:./input.json")
	}
	if *name == "" {
		log.Info("Use a default struct name: T")
	}
	if *output == "" {
		*output = defaultOutput
		log.Infof("Use default output file:%s", *output)
	}

}

func main() {
	// init flag
	initEnv()

	if !*web {
		fp, err := os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(fp)
		if err != nil {
			log.Fatal(err)
		}
		jsonParser := gogen.NewJSONParser(*pkg, *name, *output, data)
		if err = jsonParser.Parse(); err != nil {
			log.Fatal(err)
		}
		if err = jsonParser.Render(); err != nil {
			log.Fatal(err)
		}
		if err = jsonParser.Output(); err != nil {
			log.Fatal(err)
		}
		return
	}

	log.Infof("Starting web server and listen at port %s", *port)
	err := gogen.StartServer(*port)
	if err != nil {
		log.Fatal(err)
	}
}
