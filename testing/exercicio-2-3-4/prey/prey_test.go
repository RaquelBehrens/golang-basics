package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTuna(t *testing.T) {
	t.Run("case 1: successfully created", func(t *testing.T) {
		// given

		// when
		output := prey.CreateTuna()

		// then
		assert.NotNil(t, output)
	})
}

func TestGetSpeed(t *testing.T) {
	t.Run("case 1: speed is 0", func(t *testing.T) {
		// given
		impl := prey.NewTuna(0.0, nil)
		expectedResult := 0.0

		// when
		output := impl.GetSpeed()

		// then
		assert.Equal(t, expectedResult, output)
	})

	t.Run("case 2: speed is greater than 0", func(t *testing.T) {
		// given
		impl := prey.NewTuna(1.0, nil)
		expectedResult := 1.0

		// when
		output := impl.GetSpeed()

		// then
		assert.Equal(t, expectedResult, output)
	})
}

func TestGetPosition(t *testing.T) {
	t.Run("case 1: position is nil", func(t *testing.T) {
		// given
		impl := prey.NewTuna(0.0, nil)

		// when
		output := impl.GetPosition()

		// then
		assert.Nil(t, output)
	})

	t.Run("case 2: position is not nil", func(t *testing.T) {
		// given
		position := &positioner.Position{X: 0, Y: 0, Z: 2}
		impl := prey.NewTuna(1.0, position)

		// when
		output := impl.GetPosition()

		// then
		assert.NotNil(t, output)
	})
}
