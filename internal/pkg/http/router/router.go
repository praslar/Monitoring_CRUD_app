package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Config struct {
		Routes []Route
	}

	Route struct {
		Path    string
		Method  string
		Handler http.HandlerFunc
	}
)

func New(conf *Config) (http.Handler, error) {
	r := mux.NewRouter()

	for _, rt := range conf.Routes {
		var h http.Handler
		h = http.HandlerFunc(rt.Handler)
		r.Path(rt.Path).Methods(rt.Method).Handler(h)
	}

	return r, nil
}
