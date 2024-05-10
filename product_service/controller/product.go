package controller

import (
	"io"
	"net/http"
	"path/filepath"
	"rsch/product_service/exception"
	"rsch/product_service/helper"
	"rsch/product_service/model/domain"
	"rsch/product_service/model/web"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Product struct {
	ProductService domain.ProductService
}

func NewProduct(productService domain.ProductService) domain.ProductController {
	return &Product{
		ProductService: productService,
	}
}

func (controller *Product) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	helper.PanicIfError(err)

	imgFromReq, imgHeader, err := r.FormFile("img")
	helper.PanicIfError(err)

	contentType := imgHeader.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/jpg" {
		panic(exception.NewBadRequestError("the only file formats allowed are jpeg, jpg or png"))
	}

	if imgHeader.Size >= 2*1024*1024 {
		panic(exception.NewBadRequestError("maksimum size is 2mb"))
	}

	imgExt := filepath.Ext(imgHeader.Filename)
	imgData, err := io.ReadAll(imgFromReq)
	helper.PanicIfError(err)

	productImage := domain.ProductImage{
		Id:   helper.GenerateRandomString(16) + imgExt,
		Data: imgData,
	}

	category, err := strconv.Atoi(r.MultipartForm.Value["category"][0])
	helper.PanicIfError(err)

	price, err := strconv.ParseInt(r.MultipartForm.Value["price"][0], 10, 64)
	helper.PanicIfError(err)

	sizes := r.MultipartForm.Value["size"]
	stocks := r.MultipartForm.Value["stock"]

	if len(sizes) != len(stocks) {
		panic(exception.NewBadRequestError("size or stock does not match"))
	}

	detailSizes := []domain.ProductSize{}
	for i, size := range sizes {
		stock, err := strconv.Atoi(stocks[i])
		helper.PanicIfError(err)

		detailSize := domain.ProductSize{
			Size:  size,
			Stock: stock,
		}

		detailSizes = append(detailSizes, detailSize)
	}

	product := domain.Product{
		Name:        r.MultipartForm.Value["name"][0],
		Description: r.MultipartForm.Value["desc"][0],
		Purpose:     r.MultipartForm.Value["purpose"][0],
		Category:    int16(category),
		Price:       price,
		DetailSize:  detailSizes,
	}

	dataResponse := controller.ProductService.Create(r.Context(), product, productImage)

	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          dataResponse,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Product) Image(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idImage := params.ByName("idImage")
	img := controller.ProductService.Image(r.Context(), idImage)

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(img)
}

func (controller *Product) FindProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var productResponse []web.ProductResponse
	var paginationResponse web.PaginationResponse
	var err error
	var purpose string
	var category, page int

	filters := make(map[string]interface{})

	queryPage := r.URL.Query().Get("page")
	queryCategory := r.URL.Query().Get("category")
	queryPurpose := r.URL.Query().Get("purpose")
	querySearchName := r.URL.Query().Get("search")

	if queryCategory != "" {
		category, err = strconv.Atoi(queryCategory)
		helper.PanicIfError(err)
	}

	if queryPage != "" {
		page, err = strconv.Atoi(queryPage)
		helper.PanicIfError(err)
	} else {
		page = 1
	}

	if queryPurpose != "" {
		switch queryPurpose {
		case "1":
			purpose = "Men"
		case "2":
			purpose = "Women"
		case "3":
			purpose = "Unisex"
		}
	}

	switch {
	case queryCategory != "" && queryPurpose != "":
		filters["id_category"] = category
		filters["purpose"] = purpose
	case queryPurpose != "":
		filters["purpose"] = purpose
	case queryCategory != "":
		filters["id_category"] = category
	}

	filters["name"] = "%" + querySearchName + "%"
	productResponse, paginationResponse = controller.ProductService.FindProducts(r.Context(), page, filters)

	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          productResponse,
		Pagination:    paginationResponse,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Product) FindProductById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idProduct := params.ByName("idProduct")
	productResponse := controller.ProductService.FindProductById(r.Context(), idProduct)

	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          productResponse,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}
