package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func CreateNewDBForSurvey() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", viper.GetString("postgre.user"),
		viper.GetString("postgre.password"), viper.GetString("postgre.dbname"), viper.GetString("postgre.host"),
		viper.GetString("postgre.sslmode"))

	dtb, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error from `Open` function, package `sql`: %v", err)
	}
	return dtb, nil
}
