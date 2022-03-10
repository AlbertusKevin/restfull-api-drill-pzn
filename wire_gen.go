// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"net/http"
	"pzn-restful-api/controller"
	"pzn-restful-api/database"
	"pzn-restful-api/middleware"
	"pzn-restful-api/repository/category_repo"
	"pzn-restful-api/router"
	"pzn-restful-api/service/category_service"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	categoryRepository := category_repo.NewCategoryRepository()
	db := database.Connect()
	validate := validator.New()
	categoryService := category_service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	httprouterRouter := router.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(httprouterRouter)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var categorySet = wire.NewSet(category_repo.NewCategoryRepository, category_service.NewCategoryService, controller.NewCategoryController)