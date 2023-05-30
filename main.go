package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/salesof7/studies_go-hexagonal/adapters/db"
	"github.com/salesof7/studies_go-hexagonal/app"
)

func main() {
	Db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db.NewProductDb(Db)
	productService := app.NewProductService(productDbAdapter)
	product, _ := productService.Create("product exemplo", 30)
	productService.Enable(product)
}
