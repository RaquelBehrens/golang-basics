package repository

import (
	"app/internal"
	"app/internal/handler"
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewRepositoryProductDB(db *sql.DB) (r *RepositoryProductDB) {
	r = &RepositoryProductDB{
		st: db,
	}
	return
}

type RepositoryProductDB struct {
	st *sql.DB
}

func (r *RepositoryProductDB) FindById(id int) (p internal.Product, err error) {
	row := r.st.QueryRow("SELECT p.`id`, p.`name`, p.`quantity`, p.`code_value`, p.`is_published`, p.`expiration`, p.`price` FROM products p WHERE p.id = ?", id)
	if err := row.Err(); err != nil {
		return internal.Product{}, err
	}

	var pJSON handler.ProductJSON
	if err := row.Scan(&pJSON.Id, &pJSON.Name, &pJSON.Quantity, &pJSON.CodeValue, &pJSON.IsPublished, &pJSON.Expiration, &pJSON.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.Product{}, internal.ErrRepositoryProductNotFound
		}
		return internal.Product{}, err
	}

	return convertToProduct(pJSON)
}

func (r *RepositoryProductDB) Save(p *internal.Product) error {
	result, err := r.st.Exec(
		"INSERT INTO products (`name`, `quantity`, `code_value`, `is_published`, `expiration`, `price`) VALUES (?, ?, ?, ?, ?, ?)",
		(*p).Name, (*p).Quantity, (*p).CodeValue, (*p).IsPublished, (*p).Expiration, (*p).Price,
	)
	if err != nil {
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	(*p).Id = int(lastInsertId)
	return nil
}

func (r *RepositoryProductDB) UpdateOrSave(p *internal.Product) error {
	row := r.st.QueryRow("SELECT p.`id`, p.`name`, p.`quantity`, p.`code_value`, p.`is_published`, p.`expiration`, p.`price` FROM products p WHERE p.id = ?", p.Id)
	if err := row.Err(); err != nil {
		return err
	}

	var pJSON handler.ProductJSON
	if err := row.Scan(&pJSON.Id, &pJSON.Name, &pJSON.Quantity, &pJSON.CodeValue, &pJSON.IsPublished, &pJSON.Expiration, &pJSON.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return r.Save(p)
		}
		return err
	}

	return r.Update(p)
}

func (r *RepositoryProductDB) Update(p *internal.Product) error {
	_, err := r.st.Exec(
		"UPDATE products SET `name`=?, `quantity`=?, `code_value`=?, `is_published`=?, `expiration`=?, `price`=? WHERE `id`=?",
		(*p).Name, (*p).Quantity, (*p).CodeValue, (*p).IsPublished, (*p).Expiration, (*p).Price, (*p).Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryProductDB) Delete(id int) error {
	_, err := r.st.Exec("DELETE FROM products WHERE `id`=?", id)
	if err != nil {
		return err
	}
	return nil
}

func convertToProduct(pJSON handler.ProductJSON) (product internal.Product, err error) {
	expiration_date, err := time.Parse("2006-01-02", pJSON.Expiration)
	if err != nil {
		return internal.Product{}, err
	}
	product = internal.Product{
		Id: pJSON.Id,
		ProductAttributes: internal.ProductAttributes{
			Name:        pJSON.Name,
			Quantity:    pJSON.Quantity,
			CodeValue:   pJSON.CodeValue,
			IsPublished: pJSON.IsPublished,
			Expiration:  expiration_date,
			Price:       pJSON.Price,
		},
	}
	return
}
