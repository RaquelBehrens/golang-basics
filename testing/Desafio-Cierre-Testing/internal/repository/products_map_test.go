package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchProducts(t *testing.T) {
	t.Run("case 1: found product", func(t *testing.T) {
		// given
		db := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test1", Price: 1.0, SellerId: 1}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Test2", Price: 2.0, SellerId: 2}},
		}
		rp := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: 1}

		// when
		product, err := rp.SearchProducts(query)

		// then
		expectedResult := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test1", Price: 1.0, SellerId: 1}},
		}
		assert.Equal(t, product, expectedResult)
		assert.ErrorIs(t, err, nil)
	})

	t.Run("case 2: did not found product", func(t *testing.T) {
		// given
		db := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test1", Price: 1.0, SellerId: 1}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Test2", Price: 2.0, SellerId: 2}},
		}
		rp := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: 3}

		// when
		product, err := rp.SearchProducts(query)

		// then
		expectedResult := map[int]internal.Product{}
		assert.Equal(t, product, expectedResult)
		assert.ErrorIs(t, err, nil)
	})
}
