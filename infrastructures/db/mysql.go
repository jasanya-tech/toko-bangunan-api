package db

import (
	"database/sql"
	"fmt"

	"toko-bangunan/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

type MysqlImpl struct {
	DB *sql.DB
}

func NewMysqlConnection() *sql.DB {
	// log.Info().Msg("Initialize Mysql Connection")

	dbHost := config.Get().DB.Mysql.Host
	dbPort := config.Get().DB.Mysql.Port
	dbUser := config.Get().DB.Mysql.User
	dbPass := config.Get().DB.Mysql.Pass
	dbName := config.Get().DB.Mysql.Name

	fDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", fDB)
	if err != nil {
		log.Err(err).Msgf("error to load database %s", err)
	}
	log.Info().Str("name", dbName).Msg("success connect to db")
	return db
}
