package middleware

import (
	"net/http"
	"rsch/auth_service/exception"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func BearerTokenMiddleware(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			panic(exception.NewUnauthorizedError("bearer token cannot be empty"))
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}

		h(w, r, params)
	}
}
