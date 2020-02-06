package dev

import (
	"context"
	"database/sql"
	"github.com/andikanugraha11/golang-boilerplate-awesome-echo/app/model"
	dRepo "github.com/andikanugraha11/golang-boilerplate-awesome-echo/app/repository"
	"github.com/labstack/echo/v4"
)

func NewSQLDevRepo(Conn *sql.DB) dRepo.DevRepo {
	return &pqDevRepo{
		Conn: Conn,
	}
}

type pqDevRepo struct {
	Conn *sql.DB
}

func (p pqDevRepo) FetchCHI(c context.Context, num int64) ([]*model.Dev, error) {
	panic("implement me")
}

func (p pqDevRepo) fetch(c echo.Context, query string, args ... interface{}) ([]*model.Dev, error) {
	rows, err := p.Conn.QueryContext(c.Request().Context(),query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*model.Dev, 0)
	for rows.Next() {
		data := new(model.Dev)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Email,
			&data.Username,
			&data.Password,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (p pqDevRepo) Fetch(c echo.Context, num int64) ([]*model.Dev, error) {
	query := "Select id, name, email, username, password From users limit $1"

	return p.fetch(c, query, num)
}

