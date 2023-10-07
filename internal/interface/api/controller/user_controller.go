package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
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
	Validate    *validator.Validate `di.inject:"validate"`
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
	payload := new(dto.UserCreateRequest)

	err := json.NewDecoder(r.Body).Decode(payload)
	helper.PanicIfError(err)

	validationError := c.Validate.Struct(payload)
	helper.PanicIfError(validationError)

	user, err := c.UserService.Create(r.Context(), payload)
	if err != nil {
		helper.ToResponseBody(w, dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: false,
			Data:   err.Error(),
		})
		return
	}
	helper.ToResponseBody(w, user)
}
