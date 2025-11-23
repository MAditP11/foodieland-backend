package app

import (
	"database/sql"
	"foodieland/helper"
	"time"
)

func NewDB() *sql.DB {
	db,err := sql.Open("mysql", "root:1234MaPt@tcp(localhost:3306)/foodieland?parseTime=true")
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}