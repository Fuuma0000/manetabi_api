package infrastructure

import (
	"database/sql"
)

type IUserInfrastructer interface {
}

type userInfrastructer struct {
	db *sql.DB
}

func NewUserInfrastructer(db *sql.DB) IUserInfrastructer {
	return &userInfrastructer{db}
}
