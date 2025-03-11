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
