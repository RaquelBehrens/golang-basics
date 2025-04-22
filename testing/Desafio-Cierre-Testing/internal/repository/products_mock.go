package repository

import "app/internal"

func NewProductsMock() *ProductsMock {
	return &ProductsMock{}
}

type ProductsMock struct {
	FuncSearchProducts func(query internal.ProductQuery) (p map[int]internal.Product, err error)

	Spy struct {
		SearchProducts int
	}
}

func (m *ProductsMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	m.Spy.SearchProducts++
	p, err = m.FuncSearchProducts(query)
	return
}
