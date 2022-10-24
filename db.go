package belajar_golang_database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	var (
		driverName     = "mysql"
		dataSourceName = "root:root@tcp(localhost:3306)/belajar_golang_database?parseTime=true" // username:password@tcp(host:port)/database_nam
	)
	// Open
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	// Database Pooling
	db.SetMaxIdleConns(10)                  // pengaturan berapa jumlah koneksi minimal yang di buat
	db.SetMaxOpenConns(100)                 // pengaturan berapa jumlah koneksi maksimal koneksi yang di buat
	db.SetConnMaxIdleTime(5 * time.Minute)  // pengaturan berapa lama koneksi yang sudah tidak digunakan lagi akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) // pengaturan berapa lama koneksi yang boleh digunakan

	return db
}
