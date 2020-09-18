package govect

import "math"

// EPSILON ...
var EPSILON = math.Nextafter(1, 2) - 1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// New ...
func New(size int) Vector {
	return make(Vector, size)
}

// NewWithValues ...
func NewWithValues(values []float64) Vector {
	v := make(Vector, len(values))
	copy(v, values)
	return v
}

// Add ...
// Sums of two vectors, returns the resulting vector.
func Add(v1, v2 Vector) Vector {
	length := min(len(v1), len(v2))
	result := make(Vector, length)
	for i := 0; i < length; i++ {
		result[i] = v1[i] + v2[i]
	}
	return result
}

// Subtract ...
// Difference of two vectors, returns the resulting vector.
func Subtract(v1, v2 Vector) Vector {
	length := min(len(v1), len(v2))
	result := make(Vector, length)
	for i := 0; i < length; i++ {
		result[i] = v1[i] - v2[i]
	}
	return result
}

// Dot products of two vectors.
func Dot(v1, v2 Vector) (float64, error) {
	if len(v1) != len(v2) {
		return 0.0, ErrVectorNotSameSize
	}

	length := len(v1)
	result := 0.0
	for i := 0; i < length; i++ {
		result += v1[i] * v2[i]
	}

	return result, nil
}

// Cross products of two vectors.
// Vector dimensionality has to be 3.
func Cross(v1, v2 Vector) (Vector, error) {
	// Early error check to prevent redundant cloning
	if len(v1) != 3 || len(v2) != 3 {
		return nil, ErrVectorInvalidDimension
	}

	result := make(Vector, 3)
	result[0] = v1[1]*v2[2] - v1[2]*v2[1]
	result[1] = v1[2]*v2[0] - v1[0]*v2[2]
	result[2] = v1[0]*v2[1] - v1[1]*v2[0]

	return result, nil
}

// Unit ...
func Unit(v Vector) Vector {
	magRec := 1.0 / v.Magnitude()
	unit := v.Clone()
	for i := range unit {
		unit[i] *= magRec
	}
	return unit
}

// Hadamard ...
func Hadamard(v1, v2 Vector) (Vector, error) {
	if len(v1) != len(v2) {
		return nil, ErrVectorInvalidDimension
	}

	length := len(v1)
	result := make(Vector, length)
	for i := 0; i < length; i++ {
		result[i] = v1[i] * v2[i]
	}
	return result, nil
}
