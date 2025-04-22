package hunter_test

import (
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSharks(t *testing.T) {
	t.Run("case 1: successfully created white shark", func(t *testing.T) {
		// given
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 100
			return
		}
		sim := simulator.NewCatchSimulatorDefault(100, ps)

		// when
		output := hunter.CreateWhiteShark(sim)

		// then
		assert.NotNil(t, output)
	})
}

func TestGetSpeed(t *testing.T) {
	t.Run("case 1: can catch", func(t *testing.T) {
		// given
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 100
			return
		}
		sim := simulator.NewCatchSimulatorDefault(100, ps)
		shark := hunter.NewWhiteShark(10, &positioner.Position{X: 0, Y: 0, Z: 0}, sim)
		prey := prey.NewTuna(5, &positioner.Position{X: 100, Y: 0, Z: 0})

		// when
		output := shark.Hunt(prey)

		// then
		assert.NoError(t, output)
	})

	t.Run("case 2: can not catch", func(t *testing.T) {
		// given
		ps := positioner.NewPositionerStub()
		ps.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			linearDistance = 100
			return
		}
		sim := simulator.NewCatchSimulatorDefault(100, ps)
		shark := hunter.NewWhiteShark(5, &positioner.Position{X: 0, Y: 0, Z: 0}, sim)
		prey := prey.NewTuna(10, &positioner.Position{X: 100, Y: 0, Z: 0})

		// when
		output := shark.Hunt(prey)

		// then
		assert.EqualError(t, output, "can not hunt the prey: shark can not catch the prey")
	})
}
