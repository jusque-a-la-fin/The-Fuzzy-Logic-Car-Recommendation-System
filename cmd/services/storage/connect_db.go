package storage

import (
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func CreateNewRDB() (*redis.Client, error) {
	db, err := strconv.Atoi(viper.GetString("db"))
	if err != nil {
		return nil, fmt.Errorf("error geting db number the from config file: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("addr"),
		Password: viper.GetString("password"),
		DB:       db,
	})
	return rdb, nil
}
