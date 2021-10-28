package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, exist := router.FindHandler(req.URL.Path, req.Method)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, req)

}

func (router *Router) FindHandler(path, method string) (http.HandlerFunc, bool) {
	handler, exist := router.rules[path][method]
	return handler, exist
}
