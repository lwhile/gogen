package gogen

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/gogen", genStruct).Methods("POST")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	return r
}

// StartServer start a http server
func StartServer(port string) error {
	return http.ListenAndServe(":"+port, initRouter())
}
