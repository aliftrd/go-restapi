package dto

// Find All User
type FindAllUsersResponseData []FindOneUserResponseData

// Find One User
type FindOneUserResponseData struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Create User
type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum,min=8"`
}
