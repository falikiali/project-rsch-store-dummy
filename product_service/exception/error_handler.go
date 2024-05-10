package exception

import (
	"net/http"
	"rsch/product_service/helper"
	"rsch/product_service/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if badRequestError(w, r, err) {
		return
	}

	if notFoundError(w, r, err) {
		return
	}

	if unauthorizedError(w, r, err) {
		return
	}

	internalServerError(w, r)
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		webResponse := web.WebResponse{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: http.StatusText(http.StatusBadRequest),
			Data:          exception.Error,
		}

		helper.WriteToResponseBody(w, http.StatusBadRequest, webResponse)
		return true
	}
	return false
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		webResponse := web.WebResponse{
			StatusCode:    http.StatusNotFound,
			StatusMessage: http.StatusText(http.StatusNotFound),
			Data:          exception.Error,
		}

		helper.WriteToResponseBody(w, http.StatusNotFound, webResponse)
		return true
	}

	return false
}

func unauthorizedError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		webResponse := web.WebResponse{
			StatusCode:    http.StatusUnauthorized,
			StatusMessage: http.StatusText(http.StatusUnauthorized),
			Data:          exception.Error,
		}

		helper.WriteToResponseBody(w, http.StatusUnauthorized, webResponse)
		return true
	}

	return false
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	webResponse := web.WebResponse{
		StatusCode:    http.StatusInternalServerError,
		StatusMessage: http.StatusText(http.StatusInternalServerError),
	}

	helper.WriteToResponseBody(w, http.StatusInternalServerError, webResponse)
}
