package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Create(v internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	return
}

func (s *VehicleDefault) GetByColorYear(c string, y int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByColorYear(c, y)
	return
}

func (s *VehicleDefault) GetByBrandFabricatedBetween(b string, sy int, ey int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByBrandFabricatedBetween(b, sy, ey)
	return
}

func (s *VehicleDefault) GetAverageSpeedByBrand(b string) (speed float64, err error) {
	speed, err = s.rp.GetAverageSpeedByBrand(b)
	return
}

func (s *VehicleDefault) CreateBatch(v []internal.Vehicle) (err error) {
	err = s.rp.CreateBatch(v)
	return
}

func (s *VehicleDefault) UpdateSpeed(v internal.Vehicle) (err error) {
	err = s.rp.UpdateSpeed(v)
	return
}

func (s *VehicleDefault) GetByFuelType(t string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByFuelType(t)
	return
}

func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}

func (s *VehicleDefault) GetByTransmissionType(t string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByTransmissionType(t)
	return
}

func (s *VehicleDefault) UpdateFuel(v internal.Vehicle) (err error) {
	err = s.rp.UpdateFuel(v)
	return
}

func (s *VehicleDefault) GetAverageCapacityByBrand(b string) (c float64, err error) {
	c, err = s.rp.GetAverageCapacityByBrand(b)
	return
}

func (s *VehicleDefault) GetByDimentions(minL, maxL, minW, maxW float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByDimentions(minL, maxL, minW, maxW)
	return
}

func (s *VehicleDefault) GetVehiclesByWeightRange(min, max float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetVehiclesByWeightRange(min, max)
	return
}
