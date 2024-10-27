package routers

import (
	"net/http"
	"tcc-project/src/middlewares"

	"github.com/gorilla/mux"
)

// Router is the base structure of the routers of the API
type Router struct {
	Method       string
	Path         string
	Handler      func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

// Routers ...
var Routers []Router

// InitilizarRouters ...
func InitilizarRouters() {
	r := mux.NewRouter()

	for _, v := range Routers {
		apiRouters := r.PathPrefix("/api/").Subrouter()
		if v.AuthRequired {
			apiRouters.Use(middlewares.JWTMiddleWare)
		}
		apiRouters.HandleFunc(v.Path, v.Handler).Methods(v.Method)
	}

	http.Handle("/", r)
}
