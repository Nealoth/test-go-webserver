package middlewares

import (
	"net/http"
	"restapi/internal/util"
)

func AuthMiddleware(w http.ResponseWriter, r *http.Request) {
	util.WriteReqStatus(w, r, 403)
}
