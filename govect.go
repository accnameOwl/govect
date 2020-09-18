package govect

import "math"

// Vector []float64
type Vector []float64

// Clone ...
func (v Vector) Clone() Vector {
	return NewWithValues(v)
}

// Set ...
func (v Vector) Set(values []float64) {
	copy(v, values)
}

// Scale ...
func (v Vector) Scale(value float64) {
	length := len(v)
	for i := 0; i < length; i++ {
		v[i] *= value
	}
}

// Magnitude ...
func (v Vector) Magnitude() float64 {
	result := 0.0
	for _, e := range v {
		result += e * e
	}
	return math.Sqrt(result)
}

// Zero ...
func (v Vector) Zero() {
	for i := range v {
		v[i] = 0.0
	}
}

// Do ...
func (v Vector) Do(applyFn func(float64) float64) {
	for i, e := range v {
		v[i] = applyFn(e)
	}
}

// DoWithIndex ...
func (v Vector) DoWithIndex(applyFn func(int, float64) float64) {
	for i, e := range v {
		v[i] = applyFn(i, e)
	}
}
