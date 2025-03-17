package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Create(v Vehicle) (err error)
	GetByColorYear(c string, y int) (v map[int]Vehicle, err error)
	GetByBrandFabricatedBetween(b string, sy int, ey int) (v map[int]Vehicle, err error)
	GetAverageSpeedByBrand(b string) (s float64, err error)
	CreateBatch(v []Vehicle) (err error)
	UpdateSpeed(v Vehicle) (err error)
	GetByFuelType(t string) (v map[int]Vehicle, err error)
	Delete(id int) (err error)
	GetByTransmissionType(t string) (v map[int]Vehicle, err error)
	UpdateFuel(v Vehicle) (err error)
	GetAverageCapacityByBrand(b string) (c float64, err error)
	GetByDimentions(minL, maxL, minW, maxW float64) (v map[int]Vehicle, err error)
	GetVehiclesByWeightRange(min, max float64) (v map[int]Vehicle, err error)
}
