package server

import (
	"fmt"
	"net/http"
	"restapi/internal/util"
	"sort"
)

var endpointsStorage = make(map[string]map[string]*handlerType)
var preProcessGlobalMiddlewaresStorage = make([]*ApiMiddleware, 0)
var postProcessMiddlewaresStorage = make([]*ApiMiddleware, 0)

type handlerType struct {
	handlerFunc                  func(http.ResponseWriter, *http.Request)
	preProcessHandlerMiddlewares []*ApiMiddleware
}

func RegisterEndpoint(path string, method string, handler func(http.ResponseWriter, *http.Request), middlewares ...*ApiMiddleware) {

	if endpointsStorage[path] == nil {
		endpointsStorage[path] = make(map[string]*handlerType)

		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			if endpointsStorage[path][r.Method] != nil {
				runHandlingChain(w, r, endpointsStorage[path][r.Method])
			} else {
				util.WriteReqStatus(w, r, 404)
			}
		})
	}

	if endpointsStorage[path][method] != nil {
		util.Logger().Fatal(fmt.Sprintf("Possible invalidations of API mappings: %s -> %s", method, path))
	} else {
		endpointsStorage[path][method] = &handlerType{
			handlerFunc:                  handler,
			preProcessHandlerMiddlewares: middlewares,
		}

		sort.Slice(endpointsStorage[path][method].preProcessHandlerMiddlewares, func(i, j int) bool {
			return endpointsStorage[path][method].preProcessHandlerMiddlewares[i].priority < endpointsStorage[path][method].preProcessHandlerMiddlewares[j].priority
		})
	}
}

func RegisterPreProcessMiddleware(middlewares ...*ApiMiddleware) {
	preProcessGlobalMiddlewaresStorage = append(preProcessGlobalMiddlewaresStorage, middlewares...)
	sort.Slice(preProcessGlobalMiddlewaresStorage, func(i, j int) bool {
		return preProcessGlobalMiddlewaresStorage[i].priority < preProcessGlobalMiddlewaresStorage[j].priority
	})
}

func RegisterPostProcessMiddleware(middlewares ...*ApiMiddleware) {
	postProcessMiddlewaresStorage = append(postProcessMiddlewaresStorage, middlewares...)
	sort.Slice(postProcessMiddlewaresStorage, func(i, j int) bool {
		return postProcessMiddlewaresStorage[i].priority < postProcessMiddlewaresStorage[j].priority
	})
}

func runHandlingChain(w http.ResponseWriter, r *http.Request, handler *handlerType) {

	for _, middleware := range preProcessGlobalMiddlewaresStorage {
		middleware.middlewareFunc(w, r)
	}

	for _, mw := range handler.preProcessHandlerMiddlewares {
		mw.middlewareFunc(w, r)
	}

	handler.handlerFunc(w, r)

	for _, middleware := range postProcessMiddlewaresStorage {
		middleware.middlewareFunc(w, r)
	}

}
