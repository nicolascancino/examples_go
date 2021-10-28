package main

import (
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{port, NewRouter()}
}

func (server *Server) healthcheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

func (server *Server) handle(path, method string, handle http.HandlerFunc) {
	server.router.rules[path] = make(map[string]http.HandlerFunc)
	server.router.rules[path][method] = handle
}

func (server *Server) listen() error {
	http.Handle("/", server.router)
	if err := http.ListenAndServe(server.port, nil); err != nil {
		return err
	}
	return nil
}

func (server *Server) AddMiddleware(f http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for _, m := range middleware {
		f = m(f)
	}
	return f
}
