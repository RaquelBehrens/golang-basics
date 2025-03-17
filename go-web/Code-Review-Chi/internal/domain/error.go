package domain

import "errors"

var ErrIdAlreadyExists = errors.New("vehicle with this id already exists")
var ErrVehicleNotFound = errors.New("vehicle not found")
var ErrInvalidRangeFormats = errors.New("invalidRangeFormats")
