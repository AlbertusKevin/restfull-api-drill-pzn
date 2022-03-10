package main

import (
	"fmt"
	"net/http"
	"pzn-restful-api/helper"
	"pzn-restful-api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server{
	return &http.Server{
		Addr: "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicError(err)

	fmt.Println("Server running on port 3000")
}