package database

import (
	"database/sql"
	"simple-web/config"
	"simple-web/internal/helper"
	"strconv"
	"time"
)

var DB *sql.DB

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", config.DatbaseUrl)
	helper.PanicIfError(err)

	maxIdleConn, _ := strconv.Atoi(config.MaxIdleTimeConn)
	maxOpenConn, _ := strconv.Atoi(config.MaxOpenConn)
	maxLifetimeConn, _ := strconv.Atoi(config.MaxLifetimeConn)
	maxIdleTimeConn, _ := strconv.Atoi(config.MaxIdleTimeConn)

	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(maxIdleTimeConn) * time.Minute)

	return db
}
