package grades_test

import (
	"testing"
	"unittests/grades"

	"github.com/stretchr/testify/require"
)

func TestAverage_CorrectAverage(t *testing.T) {
	// arrange
	grade1 := 10.0
	grade2 := 2.0
	expected := 6.0

	// act
	result := grades.Average(grade1, grade2)

	// assert
	require.Equal(t, expected, result, "Verify the average.")
}
