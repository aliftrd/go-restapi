package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"simple-web/helper"
	"simple-web/internal/domain/entity"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []*entity.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.User, error)
	Insert(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
}

type UserRepositoryImpl struct {
}

func (r UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []*entity.User {
	query := "SELECT id, name, email, password FROM users"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	var users []*entity.User

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		users = append(users, &user)
	}

	return users
}

func (r UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = ? LIMIT 1"
	row, err := tx.QueryContext(ctx, query, email)
	helper.PanicIfError(err)

	user := entity.User{}

	if !row.Next() {
		return &user, errors.New("user not found")
	}

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	helper.PanicIfError(err)

	return &user, nil
}

func (r UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	if me, ok := err.(*mysql.MySQLError); ok {
		if me.Number == 1062 {
			return nil, errors.New("email already exists")
		}
	}
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.ID = int(id)

	return user, nil
}
