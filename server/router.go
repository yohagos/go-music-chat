package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/test", handler)
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, r)
}