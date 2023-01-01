package main

import (
	_ "github.com/go-sql-driver/mysql"
	"learn-golang-restful-api/helper"
	"learn-golang-restful-api/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	//db := app.NewDB()
	//validate := validator.New()
	//categoryRepository := repository.NewCategoryRepositoryImpl()
	//categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryControllerImpl(categoryService)
	//router := app.NewRouter(categoryController)
	//authMiddleware := middleware.NewAuthMiddleware(router)
	//server := NewServer(authMiddleware)

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
