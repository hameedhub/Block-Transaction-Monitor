package router

import (
	"github.com/gorilla/mux"
	"net/http"
)


type Router interface {
	POST(path string, handlerFunc http.HandlerFunc)
	GET(path string, handlerFunc http.HandlerFunc)
}

type route struct {
	sm *mux.Router
}

func NewRouter(mux *mux.Router) Router {
	return &route{mux}
}
func (r route) POST(path string, handlerFunc http.HandlerFunc )  {
	r.sm.Handle(path, handlerFunc).Methods(http.MethodPost)
}
func (r route) GET(path string, handlerFunc http.HandlerFunc )  {
	r.sm.Handle(path, handlerFunc).Methods(http.MethodGet)
}
