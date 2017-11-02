package service

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lwhile/gogen"
)

// Response to client
type Response struct {
	Result  string `json:"result"`
	Status  bool   `json:"statis"`
	Message string `json:"message"`
}

func genStruct(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	resp := Response{}
	if err != nil {
		resp.Status = false
	}
	jsonParser := gogen.NewJSONParser("main", "T", "./test.go", data)
	if err = jsonParser.Parse(); err != nil {
		log.Fatal(err)
	}
	if err = jsonParser.Render(); err != nil {
		log.Fatal(err)
	}
	if err = jsonParser.Output(); err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(jsonParser.String()))
}
