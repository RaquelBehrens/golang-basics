package salary_test

import (
	"testing"
	"unittests/salary"

	"github.com/stretchr/testify/require"
)

func TestSalary_CategoryA(t *testing.T) {
	// arrange
	minutosTrabalhados := 10500.0
	categoria := "A"
	expected := 787500.0

	// act
	result, err := salary.Salary(minutosTrabalhados, categoria)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, result, "Verify the category A salary.")
}

func TestSalary_CategoryB(t *testing.T) {
	// arrange
	minutosTrabalhados := 10500.0
	categoria := "B"
	expected := 315000.0

	// act
	result, err := salary.Salary(minutosTrabalhados, categoria)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, result, "Verify the category B salary.")
}

func TestSalary_CategoryC(t *testing.T) {
	// arrange
	minutosTrabalhados := 10500.0
	categoria := "C"
	expected := 175000.0

	// act
	result, err := salary.Salary(minutosTrabalhados, categoria)

	// assert
	require.NoError(t, err)
	require.Equal(t, expected, result, "Verify the category C salary.")
}

func TestSalary_CategoryNotFound(t *testing.T) {
	// arrange
	minutosTrabalhados := 10500.0
	categoria := "D"
	expected := 0.0

	// act
	result, err := salary.Salary(minutosTrabalhados, categoria)

	// assert
	require.Error(t, err)
	require.ErrorIs(t, err, salary.ErrCategoryNotFound)
	require.EqualError(t, err, salary.ErrCategoryNotFound.Error())
	require.Equal(t, expected, result, "Verify salary for unknown category.")
}
