package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB adalah variabel global koneksi database
var DB *sql.DB

// ConnectDB untuk membuat koneksi ke MySQL
func ConnectDB() {
	username := "root"
	password := ""
	host := "localhost"
	port := "3306"
	database := "rest-api-library"

	// format DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)

	// membuka koneksi database
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi database: %v", err)
	}

	// tes koneksi ke database
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	fmt.Println("Berhasil terkoneksi ke database MySQL")
}
