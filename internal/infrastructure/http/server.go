package http

import (
	"github.com/goioc/di"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simple-web/helper"
	"simple-web/internal/interface/api/controller"
	"simple-web/pkg/exception"
)

func StartServer() {
	router := httprouter.New()

	userController := di.GetInstance("userController").(controller.UserController)
	router.GET("/users", userController.FindAll)
	router.POST("/users", userController.Create)

	router.PanicHandler = exception.ErrorHandler

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
