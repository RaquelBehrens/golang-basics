package statistics_test

import (
	"testing"
	"unittests/statistics"

	"github.com/stretchr/testify/require"
)

func TestOperation_Maximum(t *testing.T) {
	// arrange
	expected := 10

	// act
	maxFunc, err := statistics.Operation(statistics.Maximum)
	maxValue := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, maxValue, "Verify the highest grade.")
}

func TestOperation_Average(t *testing.T) {
	// arrange
	expected := 4

	// act
	averageFunc, err := statistics.Operation(statistics.Average)
	averageValue := averageFunc(2, 3, 3, 4, 10, 2, 4, 5)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, averageValue, "Verify the average grade.")
}

func TestOperation_Minimum(t *testing.T) {
	// arrange
	expected := 2

	// act
	minFunc, err := statistics.Operation(statistics.Minimum)
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, minValue, "Verify the lowest grade.")
}

func TestOperation_CategoryNotFound(t *testing.T) {
	// arrange

	// act
	operation, err := statistics.Operation("max")

	// assert
	require.Error(t, err)
	require.ErrorIs(t, err, statistics.ErrCategoryNotFound)
	require.EqualError(t, err, statistics.ErrCategoryNotFound.Error())
	require.Nil(t, operation)
}
