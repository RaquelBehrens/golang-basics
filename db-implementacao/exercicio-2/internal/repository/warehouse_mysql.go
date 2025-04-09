package repository

import (
	"app/internal"
	"database/sql"
)

// NewWarehousesMySQL returns a new instance of WarehousesMySQL
func NewWarehousesMySQL(db *sql.DB) *WarehousesMySQL {
	return &WarehousesMySQL{
		db: db,
	}
}

// WarehousesMySQL is a struct that represents a warehouse repository
type WarehousesMySQL struct {
	// db is the database connection
	db *sql.DB
}

// GetOne returns a warehouse by id
func (r *WarehousesMySQL) GetOne(id int) (p internal.Warehouse, err error) {
	// execute the query
	row := r.db.QueryRow(
		"SELECT `id`, `name`, `adress`, `telephone`, `capacity`"+
			"FROM `warehouses` WHERE `id` = ?",
		id,
	)
	if err = row.Err(); err != nil {
		return
	}

	// scan the row into the warehouse
	err = row.Scan(&p.ID, &p.Name, &p.Address, &p.Telephone, &p.Capacity)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrWarehouseNotFound
		}
		return
	}

	return
}

// Store stores a warehouse
func (r *WarehousesMySQL) Store(p *internal.Warehouse) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `warehouses` (`name`, `adress`, `telephone`, `capacity`) "+
			"VALUES (?, ?, ?, ?)",
		p.Name, p.Address, p.Telephone, p.Capacity,
	)
	if err != nil {
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

func (r *WarehousesMySQL) ReportProducts(id int) ([]internal.WarehouseReport, error) {
	query := `
		SELECT w.id, w.name, w.adress, w.telephone, w.capacity, COUNT(p.id) AS product_count
		FROM warehouses w
		LEFT JOIN products p ON w.id = p.warehouse_id
	`
	if id > 0 {
		query += " WHERE w.id = ?"
	}
	query += " GROUP BY w.id"

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []internal.WarehouseReport
	for rows.Next() {
		var report internal.WarehouseReport
		err = rows.Scan(&report.ID, &report.Name, &report.Address, &report.Telephone, &report.Capacity, &report.ProductCount)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *WarehousesMySQL) GetAll() (warehouses []internal.Warehouse, err error) {
	rows, err := r.db.Query(
		"SELECT `id`, `name`, `adress`, `telephone`, `capacity` FROM warehouses",
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var w internal.Warehouse
		err = rows.Scan(&w.ID, &w.Name, &w.Address, &w.Telephone, &w.Capacity)
		if err != nil {
			return
		}
		warehouses = append(warehouses, w)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
