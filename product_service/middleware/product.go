package middleware

import (
	"net/http"
	"os"
	"rsch/product_service/exception"
	"rsch/product_service/helper"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func ApiKeyMiddleware(h httprouter.Handle) httprouter.Handle {
	err := godotenv.Load()
	helper.PanicIfError(err)

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if r.Header.Get("X-API-Key") == "" {
			panic(exception.NewUnauthorizedError("api key cannot be empty"))
		}

		if r.Header.Get("X-API-Key") != os.Getenv("API_KEY") {
			panic(exception.NewUnauthorizedError("api key is invalid"))
		}

		h(w, r, params)
	}
}
