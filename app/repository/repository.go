package repository

import (
	"context"
	"github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/model"
)

type DevRepo interface {
	Fetch(c context.Context, num int64) ([]*model.Dev, error)
}
