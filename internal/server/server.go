package server

import (
	"net/http"
	"restapi/internal/configuration"
	"restapi/internal/util"
)

type server struct {
	configuration *configuration.Configuration
}

func NewServer(configuration *configuration.Configuration) *server {
	return &server{
		configuration: configuration,
	}
}

func (s *server) Start() {
	util.Logger().Info("Starting server at ", s.configuration.Port)
	err := http.ListenAndServe(s.configuration.Port, nil)
	if err != nil {
		util.Logger().Fatal(err)
	}
}
