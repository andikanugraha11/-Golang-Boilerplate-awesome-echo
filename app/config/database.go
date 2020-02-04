package config

import "github.com/spf13/viper"

type Database struct {
	Host     	string
	Port     	int
	User     	string
	Password 	string
	Name     	string
	MaxCon   	int
	TimeOut		int
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


	return &db, nil
}
