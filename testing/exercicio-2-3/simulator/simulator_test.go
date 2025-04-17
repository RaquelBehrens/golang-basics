package simulator_test

import (
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCatch(t *testing.T) {
	t.Run("case 1: hunter gets pray", func(t *testing.T) {
		// given
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 100
			return
		}
		impl := simulator.NewCatchSimulatorDefault(100, ps)

		// when
		hunter := simulator.Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		prey := simulator.Subject{Speed: 5, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
		output := impl.CanCatch(&hunter, &prey)

		// then
		outputOk := true
		assert.Equal(t, outputOk, output)
	})

	t.Run("case 2: hunter is slower than pray", func(t *testing.T) {
		// given
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 100
			return
		}
		impl := simulator.NewCatchSimulatorDefault(100, ps)

		// when
		hunter := simulator.Subject{Speed: 5, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		prey := simulator.Subject{Speed: 10, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
		output := impl.CanCatch(&hunter, &prey)

		// then
		outputOk := false
		assert.Equal(t, outputOk, output)
	})

	t.Run("case 3: hunter is faster than pray, but there is no time", func(t *testing.T) {
		// given
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 100
			return
		}
		impl := simulator.NewCatchSimulatorDefault(10, ps)

		// when
		hunter := simulator.Subject{Speed: 10, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		prey := simulator.Subject{Speed: 5, Position: &positioner.Position{X: 100, Y: 0, Z: 0}}
		output := impl.CanCatch(&hunter, &prey)

		// then
		outputOk := false
		assert.Equal(t, outputOk, output)
	})
}
