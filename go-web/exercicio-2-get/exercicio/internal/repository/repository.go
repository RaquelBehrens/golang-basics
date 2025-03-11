package repository

import (
	"encoding/json"
	"errors"
	"exercicio/internal/domain"
	"os"
	"time"
)

type ProductRepository interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	Create(product domain.Product) error
}

type productRepository struct {
	products []domain.Product
}

func NewProductRepository(path string) (ProductRepository, error) {
	products, err := loadProductsFromFile(path)
	if err != nil {
		return nil, err
	}
	return &productRepository{products: products}, nil
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
		// Conversão segura de ID
		id, ok := productMap["id"].(float64)
		if !ok {
			return nil, errors.New("ID é inválido ou está ausente")
		}

		// Conversão segura de "expiration"
		expirationStr, ok := productMap["expiration"].(string)
		if !ok || expirationStr == "" {
			return nil, errors.New("data de validade é inválida ou está ausente")
		}

		expiration, err := time.Parse("02/01/2006", expirationStr)
		if err != nil {
			return nil, errors.New("formato de data de validade inválido: " + err.Error())
		}

		// Conversões seguras para outros campos
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
			ID:          int(id), // IDs no JSON tipicamente vêm como float64
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

func (r *productRepository) GetAll() ([]domain.Product, error) {
	return r.products, nil
}

func (r *productRepository) GetByID(id int) (*domain.Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("produto não encontrado")
}

func (r *productRepository) Create(product domain.Product) error {
	r.products = append(r.products, product)
	return nil
}
