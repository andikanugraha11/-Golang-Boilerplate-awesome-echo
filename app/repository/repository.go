package repository

import (
	"github.com/labstack/echo/v4"
)

type DevRepo interface {
	Fetch(c echo.Context, num int64) (int, error)
}
