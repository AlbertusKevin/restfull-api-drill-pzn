package setup

import (
	"database/sql"
	"net/http"
	"pzn-restful-api/controller"
	"pzn-restful-api/helper"
	"pzn-restful-api/middleware"
	"pzn-restful-api/repository/category_repo"
	"pzn-restful-api/router"
	"pzn-restful-api/service/category_service"
	"time"

	"github.com/go-playground/validator"
)

func SetupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/pzn_golang_restapi_test")
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func SetupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := category_repo.NewCategoryRepository()
	categoryService := category_service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := router.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

func TruncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}