package products

import (
	"exercicio/internal/domain"
	"fmt"
	"log"
)

type memoryRepository struct {
	storage  *Storage
	products map[int]domain.Product
}

func NewMemoryRepository(storage *Storage) domain.Repository {
	productMap, err := storage.ReadProducts()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &memoryRepository{storage: storage, products: productMap}
}

func (r *memoryRepository) GetAll() (map[int]domain.Product, error) {
	return r.products, nil
}

func (r *memoryRepository) GetByID(productId int) (*domain.Product, error) {
	if product, exists := r.products[productId]; exists {
		return &product, nil
	}
	return nil, fmt.Errorf("%w: Produto com id %d não encontrado", domain.ErrResourceNotFound, productId)
}

func (r *memoryRepository) Create(product *domain.Product) error {
	if _, exists := r.products[product.ID]; exists {
		return fmt.Errorf("%w: Produto com id %d já existe", domain.ErrResourceAlreadyExists, product.ID)
	}
	r.products[product.ID] = *product
	return nil
}

func (r *memoryRepository) Update(product *domain.Product) error {
	if _, exists := r.products[product.ID]; !exists {
		return fmt.Errorf("%w: Produto com id %d não encontrado", domain.ErrResourceNotFound, product.ID)
	}
	r.products[product.ID] = *product
	return nil
}

func (r *memoryRepository) Delete(productId int) error {
	if _, exists := r.products[productId]; !exists {
		return fmt.Errorf("%w: Produto com id %d não encontrado", domain.ErrResourceNotFound, productId)
	}
	delete(r.products, productId)
	return nil
}
