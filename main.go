package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"gofr.dev/pkg/gofr"
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

	app := gofr.New()
	createStockDatabase()

	app.GET("/", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello! This is the Stocks Management API", nil
	})

	app.GET("/view", func(ctx *gofr.Context) (interface{}, error) {
		stocks, err := viewStocks()
		if err != nil {
			fmt.Println("Couldn't view stocks")
		}
		return stocks, nil
	})

	app.GET("/d/:id", func(ctx *gofr.Context) (interface{}, error) {
		idParam := ctx.Param("id")
		if idParam == "" {
			return nil, fmt.Errorf("ID not provided")
		}

		id, err := strconv.Atoi(idParam)
		if err != nil {
			return nil, fmt.Errorf("invalid format")
		}

		deletedStock, err := viewStocks()
		if err != nil {
			return nil, err
		}

		err = deleteStock(id)
		if err != nil {
			fmt.Println("Couldn't delete stock:", err)
			return nil, err
		}

		return deletedStock, nil
	})

	app.PUT("/add", func(ctx *gofr.Context) (interface{}, error) {
		var stock Stock
		if err := json.NewDecoder(ctx.Request().Body).Decode(&stock); err != nil {
			return nil, err
		}
		err := addStock(stock)
		if err != nil {
			return nil, err
		}
		return stock, nil
	})

	app.Start()
}
