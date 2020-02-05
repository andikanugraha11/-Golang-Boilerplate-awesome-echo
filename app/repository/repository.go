package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/model"
)

type DevRepo interface {
	Fetch(c echo.Context, num int64) (*model., error)
}
