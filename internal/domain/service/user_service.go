package service

import (
	"context"
	"database/sql"
	"errors"
	"simple-web/helper"
	"simple-web/internal/application/dto"
	"simple-web/internal/domain/entity"
	"simple-web/internal/domain/repository"
)

type UserService interface {
	FindAll(ctx context.Context) *dto.FindAllUsersResponseData
	Create(ctx context.Context, data *dto.UserCreateRequest) (*dto.FindOneUserResponseData, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository `di.inject:"userRepository"`
	DB             *sql.DB                   `di.inject:"database"`
}

func (us *UserServiceImpl) FindAll(ctx context.Context) *dto.FindAllUsersResponseData {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	response := dto.FindAllUsersResponseData{}

	users := us.UserRepository.FindAll(ctx, tx)
	for _, user := range users {
		response = append(response, dto.FindOneUserResponseData{
			ID:    int32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &response
}

func (us *UserServiceImpl) Create(ctx context.Context, data *dto.UserCreateRequest) (*dto.FindOneUserResponseData, error) {
	tx, err := us.DB.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)

	user, err := us.UserRepository.Insert(ctx, tx, &entity.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		return nil, errors.New("failed to register a new user")
	}

	response := &dto.FindOneUserResponseData{
		ID:    int32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}

	return response, nil
}
