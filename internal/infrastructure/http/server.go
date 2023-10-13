package http

import (
	"github.com/goioc/di"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simple-web/internal/exception"
	"simple-web/internal/helper"
	"simple-web/internal/interface/api/controller"
)

func StartServer() {
	router := httprouter.New()

	userController := di.GetInstance("userController").(controller.UserController)
	router.GET("/users", userController.FindAll)
	router.GET("/users/:email", userController.FindByEmail)
	router.POST("/users", userController.Create)
	router.DELETE("/users/:id", userController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
