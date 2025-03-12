package products

import (
	"exercicio/internal/domain"
)

func NewProductRepository(path string) domain.Repository {
	return NewMemoryRepository(path)
}
