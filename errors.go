package govect

import (
	"errors"
)

// ErrVectorNotSameSize ...
var ErrVectorNotSameSize = errors.New("Vectors are not the same size")

// ErrVectorInvalidDimension ...
var ErrVectorInvalidDimension = errors.New("Vectors' dimensions are not of the expected size")
