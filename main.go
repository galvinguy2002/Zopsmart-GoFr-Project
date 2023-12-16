package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Stock struct {
	ID     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Volume int     `json:"volume"`
}

func createStockDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "stocks.db")
	if err != nil {
		fmt.Println("Error opening the database")
	}
	createTableQuery := `CREATE TABLE stocks(id INTEGER PRIMARY KEY AUTOINCREMENT, symbol VARCHAR(255), price REAL, volume INTEGER);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Error creating database")
	}
}

func addStock(stock Stock) error {
	_, err := db.Exec("INSERT INTO stocks(symbol, price, volume) VALUES(?, ?, ?)", stock.Symbol, stock.Price, stock.Volume)
	return err
}

func viewStocks() ([]Stock, error) {
	rows, err := db.Query("SELECT * FROM stocks")
	if err != nil {
		fmt.Println("Error while executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var stocks []Stock
	for rows.Next() {
		var s Stock
		err := rows.Scan(&s.ID, &s.Symbol, &s.Price, &s.Volume)
		if err != nil {
			fmt.Println("Error while scanning row:", err)
			return nil, err
		}
		stocks = append(stocks, s)
	}

	return stocks, nil
}

func deleteStock(id int) error {
	_, err := db.Exec("DELETE FROM stocks WHERE id=?", id)
	return err
}

func updateStock(id int, stock Stock) error {
	_, err := db.Exec("UPDATE STOCKS SET symbol=?, price=?, volume=? WHERE id=?", stock.Symbol, stock.Price, stock.Volume, id)
	return err
}
func main() {

}
