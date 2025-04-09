package internal

import "errors"

var (
	ErrWarehouseNotFound  = errors.New("repository: warehouse not found")
	ErrWarehouseNotUnique = errors.New("repository: warehouse not unique")
	ErrWarehouseRelation  = errors.New("repository: warehouse relation error")
)

type RepositoryWarehouses interface {
	GetOne(id int) (p Warehouse, err error)
	Store(p *Warehouse) (err error)
	ReportProducts(id int) (warehouses []WarehouseReport, err error)
	GetAll() (w []Warehouse, err error)
}
