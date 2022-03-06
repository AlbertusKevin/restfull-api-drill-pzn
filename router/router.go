package router

import (
	"pzn-restful-api/controller"
	"pzn-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router{
	router := httprouter.New()

	router.GET("/api/categories",categoryController.FindAll)
	router.GET("/api/categories/:categoryId",categoryController.FindById)
	router.POST("/api/categories",categoryController.Create)
	router.DELETE("/api/categories/:categoryId",categoryController.Delete)
	router.PUT("/api/categories/:categoryId",categoryController.Update)

	router.PanicHandler = exception.ErrorHandler

	return router
}