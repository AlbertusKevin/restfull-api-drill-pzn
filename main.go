package main

import (
	"fmt"
	"net/http"
	"pzn-restful-api/controller"
	"pzn-restful-api/database"
	"pzn-restful-api/helper"
	"pzn-restful-api/middleware"
	"pzn-restful-api/repository/category_repo"
	"pzn-restful-api/router"
	"pzn-restful-api/service/category_service"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	validate := validator.New()
	db := database.Connect()
	defer db.Close()

	categoryRepository := category_repo.NewCategoryRepository()
	categoryService := category_service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	
	router := router.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)

	fmt.Println("Server running on port 3000")
}