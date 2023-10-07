package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"simple-web/helper"
	"simple-web/internal/application/dto"
	"simple-web/internal/domain/entity"
	"simple-web/internal/domain/repository"
)

type UserService interface {
	FindAll(ctx context.Context) *dto.FindAllUsersResponseData
	Create(ctx context.Context, request *dto.UserCreateRequest) (*dto.FindOneUserResponseData, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository `di.inject:"userRepository"`
	DB             *sql.DB                   `di.inject:"database"`
	Validate       *validator.Validate       `di.inject:"validate"`
}

func (s UserServiceImpl) FindAll(ctx context.Context) *dto.FindAllUsersResponseData {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	response := dto.FindAllUsersResponseData{}

	users := s.UserRepository.FindAll(ctx, tx)
	for _, user := range users {
		response = append(response, dto.FindOneUserResponseData{
			ID:    int32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &response
}

func (s UserServiceImpl) Create(ctx context.Context, request *dto.UserCreateRequest) (*dto.FindOneUserResponseData, error) {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.Insert(ctx, tx, &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return nil, err
	}

	return &dto.FindOneUserResponseData{
		ID:    int32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
