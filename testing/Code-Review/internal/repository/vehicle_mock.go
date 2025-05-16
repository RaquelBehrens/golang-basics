package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) FindAll() (v map[int]internal.Vehicle, err error) {
	args := m.Mock.Called()
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *RepositoryMock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Mock.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *RepositoryMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Mock.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *RepositoryMock) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	args := m.Mock.Called(brand)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *RepositoryMock) FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
	args := m.Mock.Called(fromWeight, toWeight)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
