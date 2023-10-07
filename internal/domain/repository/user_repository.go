package repository

import (
	"context"
	"database/sql"
	"simple-web/helper"
	"simple-web/internal/domain/entity"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []*entity.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) *entity.User
	Insert(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
}

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []*entity.User {
	query := "SELECT id, name, email, password FROM users"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := new(entity.User)
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) *entity.User {
	query := "SELECT id, name, email, password FROM users WHERE email = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, query, email)
	helper.PanicIfError(err)
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	user := new(entity.User)
	err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	helper.PanicIfError(err)
	return user
}

func (u *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)

	return user, nil
}
