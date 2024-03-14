package exception

import (
	"net/http"
	"rsch/category_service/helper"
	"rsch/category_service/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if unauthorizedError(w, r, err) {
		return
	}

	if badRequestError(w, r, err) {
		return
	}

	if notFoundError(w, r, err) {
		return
	}

	internalServerError(w, r)
}

func unauthorizedError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		webResponse := web.WebResponse{
			StatusCode:    http.StatusUnauthorized,
			StatudMessage: http.StatusText(http.StatusUnauthorized),
			Data:          exception.Error,
		}

		helper.WriteToResponseBody(w, http.StatusUnauthorized, webResponse)
		return true
	}

	return false
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	excetionValidation, ok := err.(validator.ValidationErrors)
	if ok {
		data := helper.MappingValidationErrors(excetionValidation)

		webResponse := web.WebResponse{
			StatusCode:    http.StatusBadRequest,
			StatudMessage: http.StatusText(http.StatusBadRequest),
			Data:          data,
		}

		helper.WriteToResponseBody(w, http.StatusBadRequest, webResponse)
		return true
	}

	exception, ok := err.(BadRequestError)
	if ok {
		webResponse := web.WebResponse{
			StatusCode:    http.StatusBadRequest,
			StatudMessage: http.StatusText(http.StatusBadRequest),
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
			StatudMessage: http.StatusText(http.StatusNotFound),
			Data:          exception.Error,
		}

		helper.WriteToResponseBody(w, http.StatusNotFound, webResponse)
		return true
	}

	return false
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	webResponse := web.WebResponse{
		StatusCode:    http.StatusInternalServerError,
		StatudMessage: http.StatusText(http.StatusInternalServerError),
	}

	helper.WriteToResponseBody(w, http.StatusInternalServerError, webResponse)
}
