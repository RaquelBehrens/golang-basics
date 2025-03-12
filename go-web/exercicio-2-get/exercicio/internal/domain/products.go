package domain

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"codeValue"`
	IsPublished bool      `json:"isPublished"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

type Repository interface {
	GetAll() (map[int]Product, error)
	GetByID(id int) (*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id int) error
}

type Service interface {
	GetAll() (map[int]Product, error)
	GetByID(id int) (*Product, error)
	Create(reqBody RequestBodyProduct) (Product, error)
	UpdateOrCreate(id int, reqBody RequestBodyProduct) (Product, error)
	Patch(id int, updates map[string]interface{}) (Product, error)
	Delete(id int) error
}

type RequestBodyProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"codeValue"`
	IsPublished bool    `json:"isPublished"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}
