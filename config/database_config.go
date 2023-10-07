package config

import (
	"os"
)

var (
	DatbaseUrl      = os.Getenv("DATABASE")
	MaxIdleConn     = os.Getenv("DATABASE_MAX_IDLE")
	MaxOpenConn     = os.Getenv("DATABASE_MAX_OPEN")
	MaxLifetimeConn = os.Getenv("DATABASE_MAX_LIFETIME")
	MaxIdleTimeConn = os.Getenv("DATABASE_MAX_IDLE_TIME")
)
