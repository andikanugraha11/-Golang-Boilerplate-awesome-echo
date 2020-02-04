package helper

import (
	"github.com/spf13/viper"
	"log"
)

func EnvVariable(key string) string{
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
