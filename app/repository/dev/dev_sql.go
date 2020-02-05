package dev

import (
	"github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/model"
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

func (p pqDevRepo) Fetch(c echo.Context, num int64) ([]*model.Dev, error) {
	panic("implement me")
}

