package dev

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/andikanugraha11/golang-boilerplate-awesome-echo/app/model"
	dRepo "github.com/andikanugraha11/golang-boilerplate-awesome-echo/app/repository"
	"github.com/labstack/echo/v4"
	"reflect"
	"strconv"
)

func NewSQLDevRepo(Conn *sql.DB) dRepo.DevRepo {
	return &pqDevRepo{
		Conn: Conn,
	}
}

type pqDevRepo struct {
	Conn *sql.DB
}

func (p pqDevRepo) update(query string, args ... interface{}) error{
	_, err := p.Conn.Exec(query, args...)
	return err
}

func (p pqDevRepo) UpdateById(id int64, data *model.Dev) error {
	var args []interface{}
	set := ""

	rv := reflect.ValueOf(model.Dev{
		Name:     data.Name,
		Email:    data.Email,
		Username: "",
		Password: "",
	})
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		if !f.IsValid() {
			continue
		}

		v := f.Interface()
		args = append(args, v)

		col := rv.Type().Field(i).Tag.Get("col")
		set += col + " = $" + strconv.Itoa(len(args)) + " AND "
	}

	query := "UPDATE users SET ... WHERE id="
	fmt.Println(query)
	fmt.Println(set)
	return p.update(query, args ,id)
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
	query := "SELECT id, name, email, username, password FROM users LIMIT $1"
	return p.fetch(c, query, num)
}

