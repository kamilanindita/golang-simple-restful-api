package main

import (
	"kamilanindita/golang-simple-restful-api/app"
	"kamilanindita/golang-simple-restful-api/controller"
	"kamilanindita/golang-simple-restful-api/exception"
	"kamilanindita/golang-simple-restful-api/helper"
	"kamilanindita/golang-simple-restful-api/middleware"
	"kamilanindita/golang-simple-restful-api/repository"
	"kamilanindita/golang-simple-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
