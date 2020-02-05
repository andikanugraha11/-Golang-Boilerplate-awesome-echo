package repository

import (
	"github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/model"
	"github.com/labstack/echo/v4"
)

type DevRepo interface {
	Fetch(c echo.Context, num int64) ([]*model.Dev, error)
}
