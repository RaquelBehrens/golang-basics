package repository

import (
	"app/internal"
	"app/internal/domain"
	"fmt"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) Create(v internal.Vehicle) (err error) {
	_, exists := r.db[v.Id]
	if exists {
		return fmt.Errorf("%w: produto com id %d já existe", domain.ErrIdAlreadyExists, v.Id)
	}
	r.db[v.Id] = v
	return nil
}

func (r *VehicleMap) GetByColorYear(c string, y int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Color == c && value.FabricationYear == y {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) GetByBrandFabricatedBetween(b string, sy int, ey int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		fy := value.FabricationYear
		if value.Brand == b && sy <= fy && ey >= fy {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) GetAverageSpeedByBrand(b string) (s float64, err error) {
	var totalCars, sumSpeeds float64

	for _, value := range r.db {
		if value.Brand == b {
			sumSpeeds += value.MaxSpeed
			totalCars += 1
		}
	}

	s = sumSpeeds / totalCars
	return
}

func (r *VehicleMap) CreateBatch(v []internal.Vehicle) (err error) {
	for _, value := range v {
		id := value.Id
		_, exists := r.db[id]
		if exists {
			return fmt.Errorf("%w: produto com id %d já existe", domain.ErrIdAlreadyExists, id)
		}
		r.db[id] = value
	}
	return
}

func (r *VehicleMap) UpdateSpeed(v internal.Vehicle) (err error) {
	vehicle, exists := r.db[v.Id]
	if exists {
		vehicle.MaxSpeed = v.MaxSpeed
		r.db[v.Id] = vehicle
	} else {
		return fmt.Errorf("%w: veículo não encontrado", domain.ErrVehicleNotFound)
	}
	return
}

func (r *VehicleMap) GetByFuelType(t string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.FuelType == t {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) Delete(id int) (err error) {
	_, exists := r.db[id]

	if exists {
		delete(r.db, id)
	} else {
		return fmt.Errorf("%w: veículo não encontrado", domain.ErrVehicleNotFound)
	}

	return
}

func (r *VehicleMap) GetByTransmissionType(t string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Transmission == t {
			v[key] = value

		}
	}

	return
}

func (r *VehicleMap) UpdateFuel(v internal.Vehicle) (err error) {
	vehicle, exists := r.db[v.Id]
	if exists {
		vehicle.FuelType = v.FuelType
		r.db[v.Id] = vehicle
	} else {
		return fmt.Errorf("%w: veículo não encontrado", domain.ErrVehicleNotFound)
	}
	return
}

func (r *VehicleMap) GetAverageCapacityByBrand(b string) (c float64, err error) {
	var cap, total float64

	for _, value := range r.db {
		if value.Brand == b {
			cap += float64(value.Capacity)
			total += 1.0
		}
	}

	c = cap / total
	return
}

func (r *VehicleMap) GetByDimentions(minL, maxL, minW, maxW float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		validL := value.Length >= minL && value.Length <= maxL
		validW := value.Width >= minW && value.Width <= maxW
		if validL && validW {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) GetVehiclesByWeightRange(min, max float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		weight := value.Weight
		if weight >= min && weight <= max {
			v[key] = value
		}
	}

	return
}
