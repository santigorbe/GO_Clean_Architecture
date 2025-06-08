package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	return r
}

func StartServer(r *mux.Router) {
	http.ListenAndServe(":8080", r)
}
