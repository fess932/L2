package main

import (
	"fmt"
	"l2/develop/dev11/pkg"
	"net/http"
)

func NewRouter() *GRouter {
	return &GRouter{
		routes: make(map[string]route),
	}
}

// GRouter is a simple router implementation.
type GRouter struct {
	handler     http.Handler
	middlewares []func(http.Handler) http.Handler
	routes      map[string]route
}
type route struct {
	get  func(http.ResponseWriter, *http.Request)
	post func(http.ResponseWriter, *http.Request)
}

// Use add middleware to router
func (gr *GRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	gr.middlewares = append(gr.middlewares, middlewares...)
}
func (gr *GRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// init handler at first request
	if gr.handler == nil {
		gr.handler = http.HandlerFunc(gr.routeHTTP)
		for _, v := range gr.middlewares {
			gr.handler = v(gr.handler)
		}
	}

	gr.handler.ServeHTTP(w, r)
}
func (gr *GRouter) routeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, ok := gr.routes[r.URL.Path]; !ok {
		pkg.JSONError(w, http.StatusNotFound, pkg.ErrNotFound)

		return
	}

	if r.Method == "GET" {
		if gr.routes[r.URL.Path].get != nil {
			gr.routes[r.URL.Path].get(w, r)

			return
		}
	}

	if r.Method == "POST" {
		if gr.routes[r.URL.Path].post != nil {
			gr.routes[r.URL.Path].post(w, r)

			return
		}
	}

	pkg.JSONError(w, http.StatusMethodNotAllowed,
		fmt.Errorf("%s, %w", r.Method, pkg.ErrMethodNotAllowed),
	)
}
func (gr *GRouter) Get(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	var rt = gr.routes[pattern]
	rt.get = handler
	gr.routes[pattern] = rt
}
func (gr *GRouter) Post(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	var rt = gr.routes[pattern]
	rt.post = handler
	gr.routes[pattern] = rt
}
