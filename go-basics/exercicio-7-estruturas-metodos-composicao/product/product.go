package product

import (
	"errors"
	"slices"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var products = []Product{
	{ID: 1, Name: "Notebook", Price: 4000.0, Description: "SSD 1TB", Category: "Eletrônicos"},
	{ID: 2, Name: "Smartphone", Price: 2000.0, Description: "128GB, 4G", Category: "Eletrônicos"},
	{ID: 3, Name: "TV", Price: 3000.0, Description: "4K HDR", Category: "Eletrônicos"},
}

func (p Product) Save() {
	products = append(products, p)
}

func GetAll() []Product {
	return slices.Clone(products)
}

func GetById(id int) (Product, error) {
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}
	return Product{}, ErrProductNotFound
}
