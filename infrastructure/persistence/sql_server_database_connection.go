package persistence

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	util "mondu-challenge-alihamedani/infrastructure/utils"
)

func InitSqlDB(config util.Config) *sql.DB {
	var host = config.SqlServerHost
	var port = config.SqlServerPort
	var user = config.SqlServerUser
	var password = config.SqlServerPassword
	var database = config.Database

	var cnn = "server=" + host + ";user id=" + user + ";password=" + password + ";port=" + port + ";database=" + database

	db, err := sql.Open("mssql", cnn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
