package positioner

// NewPositionerStub returns a new NewPositionerStub instance
func NewPositionerStub() (positioner *PositionerStub) {
	positioner = &PositionerStub{}
	return
}

// PositionerStub is a struct that represents a stub positioner
type PositionerStub struct {
	GetLinearDistanceFunc func(from, to *Position) (linearDistance float64)
}

// GetLinearDistance returns the linear distance between 2 positions (in meters)
func (p *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	linearDistance = p.GetLinearDistanceFunc(from, to)
	return
}
