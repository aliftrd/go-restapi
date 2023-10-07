package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simple-web/helper"
	"simple-web/internal/application/dto"
	"simple-web/internal/domain/service"
)

type UserController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
}

type UserControllerImpl struct {
	UserService service.UserService `di.inject:"userService"`
}

func (c *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userResponses := c.UserService.FindAll(r.Context())
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: true,
		Data:   userResponses,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (c *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseForm()

	userRequest := dto.UserCreateRequest{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}

	user, err := c.UserService.Create(r.Context(), &userRequest)

	if err.Error() == "email already exists" {
		helper.ToResponseBody(w, dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: false,
			Data:   "email already exists",
		})
		return
	}

	helper.PanicIfError(err)
	helper.ToResponseBody(w, user)
}
