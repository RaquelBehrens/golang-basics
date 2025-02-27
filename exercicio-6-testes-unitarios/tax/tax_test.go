package tax_test

import (
	"testing"
	"unittests/tax"

	"github.com/stretchr/testify/require"
)

func TestTax_LessThanFiftyThousand(t *testing.T) {
	// arrange
	initialSalary := 40000.0
	expected := 0.0

	// act
	result := tax.Tax(initialSalary)

	// assert
	require.Equal(t, expected, result, "Verify the tax result.")
}

func TestTax_MoreThanFiftyThousand(t *testing.T) {
	// arrange
	initialSalary := 60000.0
	expected := 10200.0

	// act
	result := tax.Tax(initialSalary)

	// assert
	require.Equal(t, expected, result, "Verify the tax result.")
}

func TestTax_MoreThanOneHundredFiftyThousand(t *testing.T) {
	// arrange
	initialSalary := 160000.0
	expected := 43200.0

	// act
	result := tax.Tax(initialSalary)

	// assert
	require.Equal(t, expected, result, "Verify the tax result.")
}
