package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Database struct {
	Host     	string
	Port     	int
	User     	string
	Password 	string
	Name     	string
	MaxConn   	int
	TimeOut		int
}

func DBConnection() (*sql.DB, error) {
	database, _ := DatabaseConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		database.Host, database.Port, database.User, database.Password, database.Name,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	maxConn := database.MaxConn
	db.SetMaxOpenConns(maxConn)

	return db, nil
}

func DatabaseConfig() (*Database, error) {
	db := Database{}

	host, ok := viper.Get("DB_HOST").(string)
	if !ok {
		host = "localhost"
	}
	db.Host = host

	port, ok := viper.Get("DB_PORT").(int)
	if !ok {
		port = 5432
	}
	db.Port = port

	user, ok := viper.Get("DB_USER").(string)
	if !ok {
		user = "dummy"
	}
	db.User = user

	password, ok := viper.Get("DB_PASSWORD").(string)
	if !ok {
		password = "dummy"
	}
	db.Password = password

	name, ok := viper.Get("DB_NAME").(string)
	if !ok {
		name = "dummy"
	}
	db.Name = name

	maxConn, ok := viper.Get("DB_MAX_CONN").(int)
	if !ok {
		maxConn = -1
	}
	db.MaxConn = maxConn

	timeOut, ok := viper.Get("DB_TIME_OUT").(int)
	if !ok {
		timeOut = 0
	}
	db.TimeOut = timeOut

	return &db, nil
}
