package handlers

import (
	"net/http"
	"strconv"

	"rest-api-library/config"
	"rest-api-library/models"

	"github.com/labstack/echo/v4"
)

// helper untuk parsing ID dari param path
func parseIDParam(c echo.Context) (int, error) {
	idStr := c.Param("id")
	return strconv.Atoi(idStr)
}

// helper untuk response error JSON konsisten
func jsonError(c echo.Context, status int, message string) error {
	return c.JSON(status, map[string]string{"error": message})
}

// helper untuk ambil param query int dengan default
func getQueryParamInt(c echo.Context, key string, defaultValue int) int {
	valStr := c.QueryParam(key)
	val, err := strconv.Atoi(valStr)
	if err != nil || val < 1 {
		return defaultValue
	}
	return val
}

// Validasi input, return error message jika ada yang invalid, kosong string jika valid
func validateBookInput(b *models.Book) string {
	if b.Title == "" {
		return "Title is required"
	}
	if b.StockTotal < 0 {
		return "StockTotal cannot be negative"
	}
	if b.StockAvailable < 0 {
		return "StockAvailable cannot be negative"
	}

	return ""
}

// GetBooks ambil daftar buku dengan paginasi
func GetBooks(c echo.Context) error {
	page := getQueryParamInt(c, "page", 1)
	limit := getQueryParamInt(c, "limit", 10)
	offset := (page - 1) * limit

	query := `
		SELECT id, title, author, publisher, year, isbn, stock_total, stock_available
		FROM books
		LIMIT ? OFFSET ?`

	rows, err := config.DB.Query(query, limit, offset)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, "Failed to query books")
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Publisher, &b.Year, &b.ISBN, &b.StockTotal, &b.StockAvailable); err != nil {
			return jsonError(c, http.StatusInternalServerError, "Failed to scan book data")
		}
		books = append(books, b)
	}

	var total int
	err = config.DB.QueryRow("SELECT COUNT(*) FROM books").Scan(&total)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, "Failed to count books")
	}

	resp := map[string]interface{}{
		"page":  page,
		"limit": limit,
		"total": total,
		"books": books,
	}

	return c.JSON(http.StatusOK, resp)
}

// GetBookByID ambil buku berdasarkan ID
func GetBookByID(c echo.Context) error {
	id, err := parseIDParam(c)
	if err != nil {
		return jsonError(c, http.StatusBadRequest, "Invalid book ID")
	}

	var b models.Book
	err = config.DB.QueryRow(
		"SELECT id, title, author, publisher, year, isbn, stock_total, stock_available FROM books WHERE id = ?", id).
		Scan(&b.ID, &b.Title, &b.Author, &b.Publisher, &b.Year, &b.ISBN, &b.StockTotal, &b.StockAvailable)

	if err != nil {
		return jsonError(c, http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, b)
}

// CreateBook tambah buku baru dengan validasi input
func CreateBook(c echo.Context) error {
	var b models.Book
	if err := c.Bind(&b); err != nil {
		return jsonError(c, http.StatusBadRequest, "Invalid input")
	}

	if msg := validateBookInput(&b); msg != "" {
		return jsonError(c, http.StatusBadRequest, msg)
	}

	result, err := config.DB.Exec(
		"INSERT INTO books (title, author, publisher, year, isbn, stock_total, stock_available) VALUES (?, ?, ?, ?, ?, ?, ?)",
		b.Title, b.Author, b.Publisher, b.Year, b.ISBN, b.StockTotal, b.StockAvailable,
	)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, "Failed to insert book")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, "Failed to retrieve last insert ID")
	}
	b.ID = int(id)

	return c.JSON(http.StatusCreated, b)
}

// UpdateBook update data buku berdasarkan ID dengan validasi input
func UpdateBook(c echo.Context) error {
	id, err := parseIDParam(c)
	if err != nil {
		return jsonError(c, http.StatusBadRequest, "Invalid book ID")
	}

	var b models.Book
	if err := c.Bind(&b); err != nil {
		return jsonError(c, http.StatusBadRequest, "Invalid input")
	}

	if msg := validateBookInput(&b); msg != "" {
		return jsonError(c, http.StatusBadRequest, msg)
	}

	_, err = config.DB.Exec(
		"UPDATE books SET title=?, author=?, publisher=?, year=?, isbn=?, stock_total=?, stock_available=? WHERE id=?",
		b.Title, b.Author, b.Publisher, b.Year, b.ISBN, b.StockTotal, b.StockAvailable, id,
	)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, "Failed to update book")
	}

	b.ID = id
	return c.JSON(http.StatusOK, b)
}

// DeleteBook hapus buku berdasarkan ID
func DeleteBook(c echo.Context) error {
	id, err := parseIDParam(c)
	if err != nil {
		return jsonError(c, http.StatusBadRequest, "Invalid book ID")
	}

	_, err = config.DB.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, "Failed to delete book")
	}

	return c.NoContent(http.StatusNoContent)
}
