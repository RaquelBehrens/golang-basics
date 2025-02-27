package food_test

import (
	"testing"
	"unittests/food"

	"github.com/stretchr/testify/require"
)

func TestFood_Dog(t *testing.T) {
	// arrange
	expected := 100000

	// act
	foodDog, err := food.Food(food.Dog)
	maxValue := foodDog(10)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, maxValue, "Verify the dog food.")
}

func TestFood_Cat(t *testing.T) {
	// arrange
	expected := 50000

	// act
	foodDog, err := food.Food(food.Cat)
	maxValue := foodDog(10)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, maxValue, "Verify the cat food.")
}

func TestFood_Hamster(t *testing.T) {
	// arrange
	expected := 2500

	// act
	foodDog, err := food.Food(food.Hamster)
	maxValue := foodDog(10)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, maxValue, "Verify the hamster food.")
}

func TestFood_Tarantula(t *testing.T) {
	// arrange
	expected := 1500

	// act
	foodDog, err := food.Food(food.Tarantula)
	maxValue := foodDog(10)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, maxValue, "Verify the tarantula food.")
}

func TestFood_CategoryNotFound(t *testing.T) {
	// arrange

	// act
	_, err := food.Food("doug")

	// assert
	require.Error(t, err)
	require.ErrorIs(t, err, food.ErrCategoryNotFound)
	require.EqualError(t, err, food.ErrCategoryNotFound.Error())
}
