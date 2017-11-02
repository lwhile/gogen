package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/gogen", genStruct).Methods("POST")
	return r
}

// Start service
func Start(port string) error {
	return http.ListenAndServe(port, initRouter())
}
