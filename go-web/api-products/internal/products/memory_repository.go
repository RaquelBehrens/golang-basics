package products

import (
	"exercicio/internal/domain"
	"fmt"
)

type MemoryRepository struct {
	Storage  *Storage
	Products map[int]domain.Product
}

func NewMemoryRepository(storage *Storage, products map[int]domain.Product) domain.Repository {
	return &MemoryRepository{Storage: storage, Products: products}
}

func (r *MemoryRepository) GetAll() (map[int]domain.Product, error) {
	return r.Products, nil
}

func (r *MemoryRepository) GetByID(productId int) (*domain.Product, error) {
	if product, exists := r.Products[productId]; exists {
		return &product, nil
	}
	return nil, fmt.Errorf("%w: Produto com id %d não encontrado", domain.ErrResourceNotFound, productId)
}

func (r *MemoryRepository) Create(product *domain.Product) error {
	if _, exists := r.Products[product.ID]; exists {
		return fmt.Errorf("%w: Produto com id %d já existe", domain.ErrResourceAlreadyExists, product.ID)
	}
	r.Products[product.ID] = *product
	return nil
}

func (r *MemoryRepository) Update(product *domain.Product) error {
	if _, exists := r.Products[product.ID]; !exists {
		return fmt.Errorf("%w: Produto com id %d não encontrado", domain.ErrResourceNotFound, product.ID)
	}
	r.Products[product.ID] = *product
	return nil
}

func (r *MemoryRepository) Delete(productId int) error {
	if _, exists := r.Products[productId]; !exists {
		return fmt.Errorf("%w: Produto com id %d não encontrado", domain.ErrResourceNotFound, productId)
	}
	delete(r.Products, productId)
	return nil
}
