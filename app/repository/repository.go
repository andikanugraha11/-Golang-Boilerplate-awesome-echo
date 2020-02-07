package repository

import (
	"context"
	"github.com/andikanugraha11/golang-boilerplate-awesome-echo/app/model"
	"github.com/labstack/echo/v4"
)

type DevRepo interface {
	FetchCHI(c context.Context, num int64) ([]*model.Dev, error)
	Fetch(c echo.Context, num int64) ([]*model.Dev, error)
	UpdateById(c echo.Context, id int64, data ... interface{}) error
}
