package exception

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"simple-web/helper"
	"simple-web/internal/application/dto"
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

		webResponse := dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: false,
			Data:   exception.Error(),
		}

		helper.ToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := dto.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: false,
		Data:   err,
	}

	helper.ToResponseBody(w, response)
}
