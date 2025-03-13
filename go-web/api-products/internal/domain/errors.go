package domain

import "errors"

var ErrResourceNotFound = errors.New("resource not found")
var ErrResourceAlreadyExists = errors.New("resource already exists")
