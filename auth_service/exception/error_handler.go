package exception

import (
	"fmt"
	"net/http"
	"regexp"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/web"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if badRequestError(w, r, err) {
		return
	}

	if notFoundError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	validatorException, ok := err.(validator.ValidationErrors)
	if ok {
		data := make(map[string]string)
		for _, fieldError := range validatorException {
			re := regexp.MustCompile(`[A-Z][^A-Z]*`)
			split := re.FindAllString(fieldError.StructField(), -1)
			fieldName := strings.Join(split, " ")
			key := strings.Join(split, "_")

			var value string
			switch fieldError.Tag() {
			case "required":
				value = fmt.Sprintf("%s can not be empty", fieldName)
			case "email":
				value = fmt.Sprintf("%s invalid", fieldName)
			case "min":
				value = fmt.Sprintf("%s minimum %s characters", fieldName, fieldError.Param())
			case "eqfield":
				value = fmt.Sprintf("%s must be the same as %s", fieldName, fieldError.Param())
			case "jwt":
				value = fmt.Sprintf("%s invalid", fieldName)
			}

			data[strings.ToLower(key)] = value
		}

		webResponse := web.WebResponse{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: "Bad request",
			Data:          data,
		}
		helper.WriteToResponseBody(w, http.StatusBadRequest, webResponse)
		return true
	}

	exception, ok := err.(BadRequestError)
	if ok {
		webResponse := web.WebResponse{
			StatusCode:    http.StatusBadRequest,
			StatusMessage: "Bad request",
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
			StatusMessage: "Not found",
			Data:          exception.Error,
		}
		helper.WriteToResponseBody(w, http.StatusNotFound, webResponse)
		return true
	}

	return false
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	webResponse := web.WebResponse{
		StatusCode:    http.StatusInternalServerError,
		StatusMessage: "Internal server error",
	}

	helper.WriteToResponseBody(w, http.StatusInternalServerError, webResponse)
}
