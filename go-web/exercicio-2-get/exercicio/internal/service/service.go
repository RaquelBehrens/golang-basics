package service

import (
	"errors"
	"exercicio/internal/domain"
	"exercicio/internal/repository"
	"time"
)

type ProductService interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	Create(reqBody domain.RequestBodyProduct) (domain.Product, error)
	UpdateOrCreate(id int, reqBody domain.RequestBodyProduct) (domain.Product, error)
	Patch(id int, updates map[string]interface{}) (domain.Product, error)
	Delete(id int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{repo: r}
}

func (s *productService) GetAll() ([]domain.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetByID(id int) (*domain.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) Create(reqBody domain.RequestBodyProduct) (domain.Product, error) {
	// Validar campos obrigatórios
	if reqBody.Name == "" || reqBody.Quantity <= 0 || reqBody.CodeValue == "" || reqBody.Expiration == "" || reqBody.Price <= 0 {
		return domain.Product{}, errors.New("dados inválidos")
	}

	// Verificar se o code_value é único
	products, err := s.repo.GetAll()
	if err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.CodeValue == reqBody.CodeValue {
			return domain.Product{}, errors.New("code_value já existe")
		}
	}

	// Converter a data de validade para time.Time
	expiration, err := time.Parse("02/01/2006", reqBody.Expiration)
	if err != nil {
		return domain.Product{}, errors.New("data de validade inválida")
	}

	// Atribuir um ID único
	newID := 1
	for _, product := range products {
		if product.ID >= newID {
			newID = product.ID + 1
		}
	}

	product := domain.Product{
		ID:          newID,
		Name:        reqBody.Name,
		Quantity:    reqBody.Quantity,
		CodeValue:   reqBody.CodeValue,
		IsPublished: reqBody.IsPublished,
		Expiration:  expiration,
		Price:       reqBody.Price,
	}

	// Adicionar o produto à "database"
	if err := s.repo.Create(product); err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *productService) UpdateOrCreate(id int, reqBody domain.RequestBodyProduct) (domain.Product, error) {
	// Validar campos obrigatórios
	if reqBody.Name == "" || reqBody.Quantity <= 0 || reqBody.CodeValue == "" || reqBody.Expiration == "" || reqBody.Price <= 0 {
		return domain.Product{}, errors.New("dados inválidos")
	}

	// Verificar se o code_value é único
	products, err := s.repo.GetAll()
	if err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.CodeValue == reqBody.CodeValue {
			return domain.Product{}, errors.New("code_value já existe")
		}
	}

	// Converter a data de validade para time.Time
	expiration, err := time.Parse("02/01/2006", reqBody.Expiration)
	if err != nil {
		return domain.Product{}, errors.New("data de validade inválida")
	}

	product, err := s.repo.GetByID(id)
	if err != nil {
		// Atribuir um ID único
		newID := 1
		for _, product := range products {
			if product.ID >= newID {
				newID = product.ID + 1
			}
		}

		newProduct := &domain.Product{
			ID:          newID,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  expiration,
			Price:       reqBody.Price,
		}

		if err := s.repo.Create(*newProduct); err != nil {
			return domain.Product{}, err
		}
		return *newProduct, nil
	}

	product.Name = reqBody.Name
	product.Quantity = reqBody.Quantity
	product.CodeValue = reqBody.CodeValue
	product.IsPublished = reqBody.IsPublished
	product.Expiration = expiration
	product.Price = reqBody.Price

	// Adicionar o produto à "database"
	if err := s.repo.Update(id, *product); err != nil {
		return domain.Product{}, err
	}

	return *product, nil
}

func (s *productService) Patch(id int, updates map[string]interface{}) (domain.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return domain.Product{}, errors.New("produto não encontrado")
	}

	for key, value := range updates {
		switch key {
		case "name":
			if name, ok := value.(string); ok && name != "" {
				product.Name = name
			} else {
				return domain.Product{}, errors.New("campo nome inválido")
			}
		case "quantity":
			if quantity, ok := value.(float64); ok && quantity >= 0 {
				product.Quantity = int(quantity)
			} else {
				return domain.Product{}, errors.New("campo quantity inválido")
			}
		case "codeValue":
			if codeValue, ok := value.(string); ok && codeValue != "" {
				product.CodeValue = codeValue
			} else {
				return domain.Product{}, errors.New("campo codeValue inválido")
			}
		case "isPublished":
			if isPublished, ok := value.(bool); ok {
				product.IsPublished = isPublished
			} else {
				return domain.Product{}, errors.New("campo isPublished inválido")
			}
		case "expiration":
			if expirationStr, ok := value.(string); ok && expirationStr != "" {
				expiration, err := time.Parse("02/01/2006", expirationStr)
				if err != nil {
					return domain.Product{}, errors.New("campo inválido")
				}
				product.Expiration = expiration
			} else {
				return domain.Product{}, errors.New("campo inválido")
			}
		case "price":
			if price, ok := value.(float64); ok && price >= 0 {
				product.Price = price
			} else {
				return domain.Product{}, errors.New("campo preço inválido")
			}
		}
	}

	if err := s.repo.Update(id, *product); err != nil {
		return domain.Product{}, err
	}

	return *product, nil

}

func (s *productService) Delete(id int) error {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("produto não encontrado")
	}
	return s.repo.Delete(*product)
}
