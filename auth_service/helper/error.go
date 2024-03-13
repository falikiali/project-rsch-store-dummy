package helper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func MappingValidationErros(exception validator.ValidationErrors) map[string]string {
	data := make(map[string]string)

	for _, fieldError := range exception {
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

	return data
}
