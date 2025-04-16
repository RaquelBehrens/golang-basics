package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

func NewProductsJSON(file *os.File) *ProductsJSON {
	return &ProductsJSON{file: file}
}

type ProductsJSON struct {
	file *os.File
}

type ProductJSON struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (cj *ProductsJSON) Load() (cs []internal.Product, err error) {
	var products []ProductJSON
	err = json.NewDecoder(cj.file).Decode(&products)
	if err != nil {
		return
	}

	for _, c := range products {
		cs = append(cs, internal.Product{
			Id: c.ID,
			ProductAttributes: internal.ProductAttributes{
				Description: c.Description,
				Price:       c.Price,
			},
		})
	}
	return
}
