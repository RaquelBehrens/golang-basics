package prey

import (
	"testdoubles/positioner"
)

// NewPreyStub creates a new PreyStub
func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

// Stub is an implementation of the Prey interface
type PreyStub struct {
	// speed of the stub
	GetSpeedFunc func() (speed float64)
	// position of the stub
	GetPositionFunc func() (position *positioner.Position)
}

// GetSpeed returns the speed of the stub
func (s *PreyStub) GetSpeed() (speed float64) {
	return s.GetSpeedFunc()
}

func (s *PreyStub) GetPosition() (position *positioner.Position) {
	return s.GetPositionFunc()
}
