package repository

import (
	"app/internal"
	"database/sql"
	"fmt"
	"time"
)

// NewProductsMySQL returns a new instance of ProductsMySQL
func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{
		db: db,
	}
}

// ProductsMySQL is a struct that represents a product repository
type ProductsMySQL struct {
	// db is the database connection
	db *sql.DB
}

// GetOne returns a product by id
func (r *ProductsMySQL) GetOne(id int) (p internal.Product, err error) {
	// execute the query
	row := r.db.QueryRow(
		"SELECT `id`, `name`, `quantity`, `code_value`, `is_published`, `expiration`, `price`, `id_warehouse` "+
			"FROM `products` WHERE `id` = ?",
		id,
	)
	if err = row.Err(); err != nil {
		return
	}

	var expirationValue interface{}

	// scan the row into the product
	err = row.Scan(&p.ID, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &expirationValue, &p.Price, &p.IDWarehouse)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrProductNotFound
		}
		return
	}

	if expirationValue != nil {
		expirationBytes, ok := expirationValue.([]uint8)
		if ok {
			expirationTime, parseErr := time.Parse("2006-01-02", string(expirationBytes))
			if parseErr != nil {
				return internal.Product{}, parseErr
			}
			p.Expiration = expirationTime
		}
	}

	return
}

// Store stores a product
func (r *ProductsMySQL) Store(p *internal.Product) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `products` (`name`, `quantity`, `code_value`, `is_published`, `expiration`, `price`, `id_warehouse`) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?)",
		p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.IDWarehouse,
	)
	if err != nil {

		fmt.Println(err)
		return
	}

	// get the last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	p.ID = int(id)

	return
}

// Update updates a product
func (r *ProductsMySQL) Update(p *internal.Product) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"UPDATE `products` SET `name` = ?, `quantity` = ?, `code_value` = ?, `is_published` = ?, `expiration` = ?, `price` = ?, `id_warehouse` = ?"+
			"WHERE `id` = ?",
		p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.ID, p.IDWarehouse,
	)
	if err != nil {
		return
	}

	return
}

// Delete deletes a product by id
func (r *ProductsMySQL) Delete(id int) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"DELETE FROM `products` WHERE `id` = ?",
		id,
	)
	if err != nil {
		return
	}

	return
}

func (r *ProductsMySQL) GetAll() (products []internal.Product, err error) {
	rows, err := r.db.Query(
		"SELECT `id`, `name`, `quantity`, `code_value`, `is_published`, `expiration`, `price`, `id_warehouse` FROM products",
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var w internal.Product
		var expirationValue interface{}

		err = rows.Scan(&w.ID, &w.Name, &w.Quantity, &w.CodeValue, &w.IsPublished, &expirationValue, &w.Price, &w.IDWarehouse)
		if err != nil {
			return
		}

		if expirationValue != nil {
			expirationBytes, ok := expirationValue.([]uint8)
			if ok {
				expirationTime, parseErr := time.Parse("2006-01-02", string(expirationBytes))
				if parseErr != nil {
					return nil, parseErr
				}
				w.Expiration = expirationTime
			}
		}

		products = append(products, w)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
