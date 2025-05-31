package main

import (
	"rest-api-library/config"
	"rest-api-library/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// koneksi ke database MySQL
	config.ConnectDB()

	// inisialisasi echo instance
	e := echo.New()

	// daftarkan routes API
	routes.SetupRoutes(e)

	// jalankan server pada port 8080
	e.Logger.Fatal(e.Start(":8000"))
}
