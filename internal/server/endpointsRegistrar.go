package server

import (
	"net/http"
	"restapi/internal/endpoints"
	"restapi/internal/middlewares"
)

var (
	post = http.MethodPost
	get  = http.MethodGet

	authMiddleware = CreateMiddleware(middlewares.AuthMiddleware, 0)
	logRequestMiddleware = CreateMiddleware(middlewares.RequestLoggingMiddleware, 0)
	logResponseMiddleware = CreateMiddleware(middlewares.ResponseLoggingMiddleware, 9999)
)

func InitEndpoints() {
	RegisterEndpoint("/user", post, endpoints.CreateUserEndpointHandler)
	RegisterEndpoint("/user", get, endpoints.GetUsersEndpointHandler, authMiddleware)
	RegisterEndpoint("/login", post, endpoints.PerformUserLogin)
}

func InitGlobalMiddlewares() {
	RegisterPreProcessMiddleware(logRequestMiddleware)
	RegisterPostProcessMiddleware(logResponseMiddleware)
}
