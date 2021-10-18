package pkg

import (
	"database/sql"
	"time"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_kafka?parseTime=true")
	PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
