package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DBClients struct{}

func NewDBClients() *DBClients {
	return &DBClients{}
}

func (d *DBClients) ConnectDB() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln(err)
	}

	c := mysql.Config{
		DBName:               "todo",
		User:                 "gogo",
		Passwd:               "gogo",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	return db, err
}
