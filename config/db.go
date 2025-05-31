package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB adalah variabel global koneksi database
var DB *sql.DB

// ConnectDB untuk membuat koneksi ke MySQL
func ConnectDB() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Ambil variabel dari .env
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

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
