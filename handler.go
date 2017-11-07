package gogen

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lwhile/log"
)

// Response to client
type Response struct {
	Result  string `json:"result"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func genStruct(w http.ResponseWriter, r *http.Request) {
	resp := Response{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.Status = false
		resp.Message = "服务器异常"
		renderJSON(w, &resp)
		return
	}

	jsonParser := NewJSONParser("main", "T", "./test.go", data)
	if err = jsonParser.Parse(); err != nil {
		log.Error(err)
		resp.Status = false
		resp.Message = "非JSON数据"
		renderJSON(w, &resp)
		return
	}
	if err = jsonParser.Render(); err != nil {
		log.Error(err)
		resp.Status = false
		resp.Message = "服务器异常"
		renderJSON(w, &resp)
		return
	}

	resp.Status = true
	resp.Message = "success"
	resp.Result = jsonParser.String()
	renderJSON(w, &resp)
}

func renderJSON(w http.ResponseWriter, resp *Response) {
	d, err := json.Marshal(resp)
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Write(d)
}
