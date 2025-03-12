package products

import (
	"encoding/json"
	"errors"
	"exercicio/internal/domain"
	"fmt"
	"log"
	"os"
	"time"
)

type memoryRepository struct {
	products map[int]domain.Product
}

func NewMemoryRepository(path string) domain.Repository {
	loadedProducts, err := loadProductsFromFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	productMap := make(map[int]domain.Product)
	for _, product := range loadedProducts {
		productMap[product.ID] = product
	}
	return &memoryRepository{products: productMap}
}

func loadProductsFromFile(path string) ([]domain.Product, error) {
	var products []map[string]interface{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	myDecoder := json.NewDecoder(file)
	err = myDecoder.Decode(&products)
	if err != nil {
		return nil, err
	}

	var convertedProducts []domain.Product
	for _, productMap := range products {
		id, ok := productMap["id"].(float64)
		if !ok {
			return nil, errors.New("ID é inválido ou está ausente")
		}

		expirationStr, ok := productMap["expiration"].(string)
		if !ok || expirationStr == "" {
			return nil, errors.New("data de validade é inválida ou está ausente")
		}

		expiration, err := time.Parse("02/01/2006", expirationStr)
		if err != nil {
			return nil, errors.New("formato de data de validade inválido: " + err.Error())
		}

		name, ok := productMap["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("nome é inválido ou está ausente")
		}

		quantity, ok := productMap["quantity"].(float64)
		if !ok {
			return nil, errors.New("quantidade é inválida ou está ausente")
		}

		codeValue, ok := productMap["code_value"].(string)
		if !ok || codeValue == "" {
			return nil, errors.New("code value é inválido ou está ausente")
		}

		isPublished, ok := productMap["is_published"].(bool)
		if !ok {
			return nil, errors.New("is Published é inválido ou está ausente")
		}

		price, ok := productMap["price"].(float64)
		if !ok {
			return nil, errors.New("preço é inválido ou está ausente")
		}

		convertedProduct := domain.Product{
			ID:          int(id),
			Name:        name,
			Quantity:    int(quantity),
			CodeValue:   codeValue,
			IsPublished: isPublished,
			Expiration:  expiration,
			Price:       price,
		}

		convertedProducts = append(convertedProducts, convertedProduct)
	}

	return convertedProducts, nil
}

func (r *memoryRepository) GetAll() ([]domain.Product, error) {
	products := make([]domain.Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
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
