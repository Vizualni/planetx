package planetx

import "errors"

var (
	ErrSameNodeMultiplied = errors.New("graph can't have same node multiple times")
	ErrInvalidDirection   = errors.New("invalid direction")
	ErrCityPointsToItself = errors.New("city points to itself")
	ErrCityNameInvalid    = errors.New("city name is invalid")
	ErrInvalidInput       = errors.New("invalid planetx text input")
)
