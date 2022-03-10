//go:build wireinject
// +build wireinject
package main

import (
	"net/http"
	"pzn-restful-api/controller"
	"pzn-restful-api/database"
	"pzn-restful-api/middleware"
	"pzn-restful-api/repository/category_repo"
	"pzn-restful-api/router"
	"pzn-restful-api/service/category_service"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	category_repo.NewCategoryRepository,
	category_service.NewCategoryService,
	controller.NewCategoryController,
)

func InitializedServer() *http.Server{
	wire.Build(
		database.Connect, 
		validator.New, 
		categorySet,
		// return NewRouter adalah http.Router, salah satu implement dari interface http.Handler
		router.NewRouter,
		// karena middleware butuhnya interface http.Handler, tapi kita tidak ada provider yang returnya nya interface
		// kita inginnya router tersebut yang sebagai handler, yang merupakan implementasi dari http.Handler
		// maka kita bind, jika butuh interface handler, berikan router
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}