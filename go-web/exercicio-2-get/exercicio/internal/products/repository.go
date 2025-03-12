package products

import (
	"exercicio/internal/domain"
)

func NewProductRepository(storage *Storage) domain.Repository {
	return NewMemoryRepository(storage)
}
