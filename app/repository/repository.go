package repository

import (
	"context"
	"github.com/andikanugraha11/golang-boilerplate-awesome-echo/app/model"
	"github.com/labstack/echo/v4"
)

type DevRepo interface {
	FetchCHI(c context.Context, num int) ([]*model.Dev, error)
	Fetch(c echo.Context, num int) ([]*model.Dev, error)
	UpdateById(id int, data *model.Dev) error
}
