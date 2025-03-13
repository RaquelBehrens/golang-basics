package products

import (
	"exercicio/internal/domain"
)

func NewProductRepository(storage *Storage, products map[int]domain.Product) domain.Repository {
	return NewMemoryRepository(storage, products)
}
