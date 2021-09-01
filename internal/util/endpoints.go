package util

import (
	"context"
	"net/http"
)

func WriteReqStatus(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	context.WithValue(r.Context(), "status", status)
}

func GetReqStatus(r *http.Request) interface{} {
	v := r.Context().Value("status")
	return v
}
