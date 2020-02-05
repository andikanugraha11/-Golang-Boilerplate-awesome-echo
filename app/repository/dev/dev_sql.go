package dev

import (
	"github.com/labstack/echo/v4"
	"database/sql"

	dRepo "github.com/andikanugraha11/golang-boilerplate-awesome-echo/app/repository"
)

func NewSQLDevRepo(Conn *sql.DB) dRepo.DevRepo {
	return &pqDevRepo{
		Conn: Conn,
	}
}

type pqDevRepo struct {
	Conn *sql.DB
}