package dto

type SuccessResponse struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Errors interface{} `json:"errors"`
}
