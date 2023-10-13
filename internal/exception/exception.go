package exception

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"simple-web/internal/application/dto"
	"simple-web/internal/helper"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := dto.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: false,
			Errors: exception.Error(),
		}

		helper.ToResponseBody(w, response, response.Code)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := dto.ErrorResponse{
		Code:   http.StatusInternalServerError,
		Status: false,
		Errors: err,
	}

	helper.ToResponseBody(w, response, response.Code)
}
