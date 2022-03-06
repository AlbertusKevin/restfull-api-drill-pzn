package controller

import (
	"net/http"
	"pzn-restful-api/helper"
	"pzn-restful-api/model/web"
	"pzn-restful-api/service/category_service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService category_service.CategoryService
}

func NewCategoryController(categoryService category_service.CategoryService) CategoryController{
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (categoryController *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse :=  categoryController.CategoryService.Create(r.Context(),categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "ok",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(w,webResponse)
}

func (categoryController *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)
	
	categoryId, err := strconv.Atoi(Params.ByName("categoryId"))
	helper.PanicError(err)

	categoryUpdateRequest.Id = categoryId

	categoryResponse :=  categoryController.CategoryService.Update(r.Context(),categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "ok",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(w,webResponse)
}

func (categoryController *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	categoryId, err := strconv.Atoi(Params.ByName("categoryId"))
	helper.PanicError(err)

	status, _ := categoryController.CategoryService.Delete(r.Context(), categoryId)
	
	webResponse := web.WebResponse{
		Code: 200,
		Status: "ok",
	}

	if !status{
		webResponse= web.WebResponse{
			Code: 400,
			Status: "failed",
			Data: "Data tidak ditemukan",
		}
	}

	helper.WriteToResponseBody(w,webResponse)
}

func (categoryController *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	categoryId, err := strconv.Atoi(Params.ByName("categoryId"))
	helper.PanicError(err)

	categoryResponse := categoryController.CategoryService.FindById(r.Context(), categoryId)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "ok",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(w,webResponse)
}

func (categoryController *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	categoryResponses := categoryController.CategoryService.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code: 200,
		Status: "ok",
		Data: categoryResponses,
	}

	helper.WriteToResponseBody(w,webResponse)
}