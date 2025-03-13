package products

import (
	"encoding/json"
	"errors"
	"exercicio/internal/domain"
	"os"
	"time"
)

type Storage struct {
	filePath string
}

func NewStorage(path string) *Storage {
	return &Storage{filePath: path}
}

func (s *Storage) ReadProducts() (map[int]domain.Product, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []map[string]interface{}
	myDecoder := json.NewDecoder(file)
	err = myDecoder.Decode(&products)
	if err != nil {
		return nil, err
	}

	productsMap := make(map[int]domain.Product)
	for _, product := range products {
		id, ok := product["id"].(float64)
		if !ok {
			return nil, errors.New("ID é inválido ou está ausente")
		}

		expirationStr, ok := product["expiration"].(string)
		if !ok || expirationStr == "" {
			return nil, errors.New("data de validade é inválida ou está ausente")
		}

		expiration, err := time.Parse("02/01/2006", expirationStr)
		if err != nil {
			return nil, errors.New("formato de data de validade inválido: " + err.Error())
		}

		name, ok := product["name"].(string)
		if !ok || name == "" {
			return nil, errors.New("nome é inválido ou está ausente")
		}

		quantity, ok := product["quantity"].(float64)
		if !ok {
			return nil, errors.New("quantidade é inválida ou está ausente")
		}

		codeValue, ok := product["code_value"].(string)
		if !ok || codeValue == "" {
			return nil, errors.New("code value é inválido ou está ausente")
		}

		isPublished, ok := product["is_published"].(bool)
		if !ok {
			return nil, errors.New("is Published é inválido ou está ausente")
		}

		price, ok := product["price"].(float64)
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

		productsMap[convertedProduct.ID] = convertedProduct
	}

	return productsMap, nil
}

func (s *Storage) WriteProducts(products map[int]domain.Product, path string) error {
	prodList := make([]domain.Product, 0, len(products))
	for _, product := range products {
		prodList = append(prodList, product)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(prodList); err != nil {
		return err
	}

	return nil
}
