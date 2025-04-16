package main

import (
	"app/internal/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := &application.ConfigApplicationLoader{
		Db: &mysql.Config{
			User:   "sgbd",
			Passwd: "root",
			Net:    "tcp",
			Addr:   "localhost:3306",
			DBName: "fantasy_products",
		},
		FilePathCustomers: "../../docs/db/json/customers.json",
		FilePathProducts:  "../../docs/db/json/products.json",
		FilePathInvoices:  "../../docs/db/json/invoices.json",
		FilePathSales:     "../../docs/db/json/sales.json",
	}
	app := application.NewApplicationLoader(cfg)
	defer app.TearDown()
	// - set up
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
