package simulator

// NewCatchSimulatorMock creates a new CatchSimulatorMock
func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}

// CatchSimulatorMock is a mock implementation of CatchSimulator
type CatchSimulatorMock struct {
	canCatchFunc func(hunter, prey *Subject) (ok bool)

	// Observer
	Calls struct {
		// CanCatch is the number of times the CanCatch method has been called
		CanCatch int
	}
}

// CanCatch returns true if the hunter can catch the prey
func (m *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (ok bool) {
	m.Calls.CanCatch++
	return m.canCatchFunc(hunter, prey)
}
