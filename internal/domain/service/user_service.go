package service

import (
	"context"
	"database/sql"
	"errors"
	"simple-web/internal/application/dto"
	"simple-web/internal/domain/entity"
	"simple-web/internal/domain/repository"
	"simple-web/internal/helper"
)

type UserService interface {
	FindAll(ctx context.Context) *dto.FindAllUsersResponseData
	FindByEmail(ctx context.Context, email string) *dto.FindOneUserResponseData
	FindByID(ctx context.Context, id int) *dto.FindOneUserResponseData
	Create(ctx context.Context, data *dto.UserCreateRequest) (*dto.FindOneUserResponseData, error)
	Delete(ctx context.Context, id int) error
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

func (us *UserServiceImpl) FindByEmail(ctx context.Context, email string) *dto.FindOneUserResponseData {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := us.UserRepository.FindByEmail(ctx, tx, email)
	if user == nil {
		return nil
	}

	return &dto.FindOneUserResponseData{
		ID:    int32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
}

func (us *UserServiceImpl) FindByID(ctx context.Context, id int) *dto.FindOneUserResponseData {
	tx, err := us.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := us.UserRepository.FindByID(ctx, tx, id)
	if user == nil {
		return nil
	}

	return &dto.FindOneUserResponseData{
		ID:    int32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
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

func (us *UserServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := us.DB.Begin()
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(tx)

	user := us.UserRepository.FindByID(ctx, tx, id)
	if user == nil {
		return errors.New("user not found")
	}

	err = us.UserRepository.Delete(ctx, tx, id)
	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
