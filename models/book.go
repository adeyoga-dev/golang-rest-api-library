package models

type Book struct {
	ID             int    `db:"id" json:"id"`
	Title          string `db:"title" json:"title"`
	Author         string `db:"author" json:"author"`
	Publisher      string `db:"publisher" json:"publisher"`
	Year           int    `db:"year" json:"year"`
	ISBN           string `db:"isbn" json:"isbn"`
	StockTotal     int    `db:"stock_total" json:"stock_total"`
	StockAvailable int    `db:"stock_available" json:"stock_available"`
}
