package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	user       = "root"
	password   = "123456"
	dataBase   = "db_products"
	server     = "127.0.0.1"
	port       = "3306"
)

func Open() *sql.DB {
	dataSourceName := user + ":" + password + "@tcp(" + server + ":" + port + ")/" + dataBase + "?parseTime=true"
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Connection with database successfully")
	return db
}
