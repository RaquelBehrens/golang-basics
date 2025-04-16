package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// given
		ws := hunt.NewWhiteShark(true, false, 10)
		tn := hunt.NewTuna("Tuna", 5)

		// when
		err := ws.Hunt(tn)

		// then
		assert.Nil(t, err)
		assert.Equal(t, ws.Hungry, false)
		assert.Equal(t, ws.Tired, true)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// given
		ws := hunt.NewWhiteShark(false, false, 10)
		tn := hunt.NewTuna("Tuna", 5)

		// when
		err := ws.Hunt(tn)

		// then
		assert.NotNil(t, err)
		assert.Equal(t, err, hunt.ErrSharkIsNotHungry)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// given
		ws := hunt.NewWhiteShark(true, true, 10)
		tn := hunt.NewTuna("Tuna", 5)

		// when
		err := ws.Hunt(tn)

		// then
		assert.NotNil(t, err)
		assert.Equal(t, err, hunt.ErrSharkIsTired)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// given
		ws := hunt.NewWhiteShark(true, false, 3)
		tn := hunt.NewTuna("Tuna", 5)

		// when
		err := ws.Hunt(tn)

		// then
		assert.NotNil(t, err)
		assert.Equal(t, err, hunt.ErrSharkIsSlower)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// given
		ws := hunt.NewWhiteShark(true, false, 3)

		// when
		err := ws.Hunt(nil)

		// then
		assert.NotNil(t, err)
		assert.Equal(t, err, hunt.ErrTunaIsNil)
	})
}
