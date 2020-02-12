package repository

import (
	"context"
	"github.com/andikanugraha11/golang-boilerplate-awesome-echo/app/model"
	"github.com/labstack/echo/v4"
)

type DevRepo interface {
	FetchCHI(c context.Context, num int64) ([]*model.Dev, error)
	Fetch(c echo.Context, num int64) ([]*model.Dev, error)
	UpdateById(id int64, data *model.Dev) error
}
