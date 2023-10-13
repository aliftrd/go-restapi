package repository

import (
	"context"
	"database/sql"
	"simple-web/internal/domain/entity"
	"simple-web/internal/helper"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []*entity.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) *entity.User
	FindByID(ctx context.Context, tx *sql.Tx, id int) *entity.User
	Insert(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
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

func (u *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) *entity.User {
	query := "SELECT id, name, email, password FROM users WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, query, id)
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

func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
