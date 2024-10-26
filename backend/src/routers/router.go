package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router is the base structure of the routers of the API
type Router struct {
	Method  string
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

// Routers ...
var Routers []Router

// InitilizarRouters ...
func InitilizarRouters() {
	r := mux.NewRouter()

	for _, v := range Routers {
		r.HandleFunc(v.Path, v.Handler).Methods(v.Method)
	}

	http.Handle("/", r)
}
