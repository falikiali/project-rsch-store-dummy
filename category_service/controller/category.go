package controller

import (
	"net/http"
	"rsch/category_service/helper"
	"rsch/category_service/model/domain"
	"rsch/category_service/model/web"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Category struct {
	CategoryService domain.CategoryService
}

func NewCategory(categoryService domain.CategoryService) domain.CategoryController {
	return &Category{
		CategoryService: categoryService,
	}
}

func (controller *Category) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := web.CreateCategoryRequest{}
	helper.ReadFromRequestBody(r, &request)

	data := controller.CategoryService.Create(r.Context(), request)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatudMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Category) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := web.UpdateCategoryRequest{}
	helper.ReadFromRequestBody(r, &request)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	request.Id = int16(id)

	data := controller.CategoryService.Update(r.Context(), request)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatudMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Category) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatudMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}
