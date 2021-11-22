package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(l *log.Logger) {
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER_1"),
		Passwd: os.Getenv("MYSQL_PASS_1"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "twitter_clone",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		l.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		l.Fatal(err)
	}

	DB = db

	l.Println("Connected to database")
}

func GetDB() *sql.DB {
	return DB
}
