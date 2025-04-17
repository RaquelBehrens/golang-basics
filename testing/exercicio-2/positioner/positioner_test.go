package positioner_test

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests for the WhiteShark implementation - Hunt method
func TestGetLinearDistance(t *testing.T) {
	t.Run("case 1: negative coordinates", func(t *testing.T) {
		// given
		impl := positioner.NewPositionerDefault()
		from := &positioner.Position{X: -1, Y: -1, Z: -1}
		to := &positioner.Position{X: -1, Y: -1, Z: -1}
		expectedResult := 0.0

		// when
		linearDistance := impl.GetLinearDistance(from, to)

		// then
		assert.Equal(t, expectedResult, linearDistance)
	})

	t.Run("case 2: positive coordinates", func(t *testing.T) {
		// given
		impl := positioner.NewPositionerDefault()
		from := &positioner.Position{X: 1, Y: 1, Z: 1}
		to := &positioner.Position{X: 1, Y: 1, Z: 1}
		expectedResult := 0.0

		// when
		linearDistance := impl.GetLinearDistance(from, to)

		// then
		assert.Equal(t, expectedResult, linearDistance)
	})

	t.Run("case 3: coordinates return linear distance without decimals", func(t *testing.T) {
		// given
		impl := positioner.NewPositionerDefault()
		from := &positioner.Position{X: 0, Y: 0, Z: 2}
		to := &positioner.Position{X: 0, Y: 0, Z: 4}
		expectedResult := 2.0

		// when
		linearDistance := impl.GetLinearDistance(from, to)

		// then
		assert.Equal(t, expectedResult, linearDistance)
	})
}
