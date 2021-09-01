package endpoints

import (
	"encoding/json"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository"
	"restapi/internal/util"
)

func CreateUserEndpointHandler(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		util.Logger().Error(err)
		w.WriteHeader(400)
		return
	}

	if err := repository.CreateUser(u); err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func GetUsersEndpointHandler(w http.ResponseWriter, r *http.Request) {

	users, err := repository.GetUsers()

	if err != nil {
		w.WriteHeader(500)
		return
	}

	encodingError := json.NewEncoder(w).Encode(users)

	if encodingError != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func PerformUserLogin(w http.ResponseWriter, r *http.Request) {

}