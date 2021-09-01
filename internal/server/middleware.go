package server

import "net/http"

type ApiMiddleware struct {
	middlewareFunc func(http.ResponseWriter, *http.Request)
	priority       int
}

func CreateMiddleware(middlewareFunc func(http.ResponseWriter, *http.Request), priority int) *ApiMiddleware {
	return &ApiMiddleware{
		middlewareFunc: middlewareFunc,
		priority:       priority,
	}
}
