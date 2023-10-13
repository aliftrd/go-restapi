package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simple-web/internal/application/dto"
	"simple-web/internal/domain/service"
	"simple-web/internal/helper"
	"strconv"
)

type UserController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
	FindByEmail(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, p httprouter.Params)
}

type UserControllerImpl struct {
	UserService service.UserService `di.inject:"userService"`
	Validate    *validator.Validate `di.inject:"validate"`
}

func (c *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users := c.UserService.FindAll(r.Context())
	response := dto.SuccessResponse{
		Code:   http.StatusOK,
		Status: true,
		Data:   users,
	}

	helper.ToResponseBody(w, response, response.Code)
}

func (c *UserControllerImpl) FindByEmail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := c.UserService.FindByEmail(r.Context(), p.ByName("email"))

	if user == nil {
		response := dto.ErrorResponse{
			Code:   http.StatusNotFound,
			Status: false,
			Errors: "User not found"}
		helper.ToResponseBody(w, response, response.Code)
		return
	}

	response := dto.SuccessResponse{
		Code:   http.StatusOK,
		Status: true,
		Data:   user,
	}

	helper.ToResponseBody(w, response, response.Code)
}

func (c *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	payload := new(dto.UserCreateRequest)

	err := json.NewDecoder(r.Body).Decode(payload)
	helper.PanicIfError(err)

	validationError := c.Validate.Struct(payload)
	helper.PanicIfError(validationError)

	user, err := c.UserService.Create(r.Context(), payload)
	if err != nil {
		response := dto.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: false,
			Errors: err.Error(),
		}
		helper.ToResponseBody(w, response, response.Code)
		return
	}

	helper.ToResponseBody(w, user, http.StatusCreated)
}

func (c *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	err := c.UserService.Delete(r.Context(), id)

	if err != nil {
		var response dto.ErrorResponse

		if err.Error() == "user not found" {
			response = dto.ErrorResponse{
				Code:   http.StatusNotFound,
				Status: false,
				Errors: err.Error(),
			}
		}

		if err.Error() == "failed to delete user" {
			response = dto.ErrorResponse{
				Code:   http.StatusInternalServerError,
				Status: false,
				Errors: err.Error(),
			}
		}

		helper.ToResponseBody(w, response, response.Code)
		return
	}

	response := dto.SuccessResponse{
		Code:   http.StatusOK,
		Status: true,
		Data:   "User deleted",
	}

	helper.ToResponseBody(w, response, response.Code)
}
