package exception

import (
	"net/http"
	"rsch/profile_service/helper"
	"rsch/profile_service/model/web"

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
			StatusMessage: http.StatusText(http.StatusUnauthorized),
			Data:          exception.Error,
		}
		helper.WriteToResponseBody(w, webResponse.StatusCode, webResponse)

		return true
	}

	return false
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	validatorException, ok := err.(validator.ValidationErrors)
	if ok {
		data := helper.MappingValidationErros(validatorException)

		webResponse := web.WebResponse{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: http.StatusText(http.StatusBadRequest),
			Data:          data,
		}

		helper.WriteToResponseBody(w, http.StatusBadRequest, webResponse)

		return true
	}

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

func internalServerError(w http.ResponseWriter, r *http.Request) {
	webResponse := web.WebResponse{
		StatusCode:    http.StatusInternalServerError,
		StatusMessage: http.StatusText(http.StatusInternalServerError),
	}

	helper.WriteToResponseBody(w, http.StatusInternalServerError, webResponse)
}
