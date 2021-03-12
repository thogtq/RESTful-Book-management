package driver

import (
	"database/sql"
	"log"
	"os"
	"time"

	//mysql driver for database/sql package
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//ConnectDB to connect to mysql database
func ConnectDB() *sql.DB {
	var ( //database connection varibles
		dbConnection   = os.Getenv("DB_CONNECTION")
		mysqlUsername  = os.Getenv("DB_USERNAME")
		mysqlPassword  = os.Getenv("DB_PASSWORD")
		mysqlHost      = os.Getenv("DB_HOST")
		mysqlPort      = os.Getenv("DB_PORT")
		mysqlDb        = os.Getenv("DB_DATABASE")
		dataSourceName = mysqlUsername + ":" + mysqlPassword + "@tcp(" + mysqlHost + mysqlPort + ")/" + mysqlDb
	)
	var err error
	db, err = sql.Open(dbConnection, dataSourceName)
	if err != nil {
		//An exception must not happens
		log.Printf("cannot connect to mysql database\n")
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	return db
}
