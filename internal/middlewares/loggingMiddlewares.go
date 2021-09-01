package middlewares

import (
	"net/http"
	"restapi/internal/util"
)

func RequestLoggingMiddleware(w http.ResponseWriter, r *http.Request) {
	util.WriteReqStatus(w, r, 400)
	util.Logger().Info("API <-- ", r.Method, " ", r.URL.Path)
}

func ResponseLoggingMiddleware(w http.ResponseWriter, r *http.Request) {
	util.GetReqStatus(r)
	util.Logger().Info("API --> ", r.Method, " ", r.URL.Path)
}
