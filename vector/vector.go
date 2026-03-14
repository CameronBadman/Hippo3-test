package vector

import "math/rand"

type Vector struct {
	Values []float32
}

func New(dims int) Vector {
	return Vector{Values: make([]float32, dims)}
}

func NewRandom(dims int) Vector {
	Values := make([]float32, dims)

	for i := range Values {
		Values[i] = rand.Float32()
	}

	return Vector{Values: Values}
}

func GetL2(v1 Vector, v2 Vector) float32 {
	var sum float32 = 0
	for i := range v1.Values {
		dx := (v1.Values[i] - v2.Values[i])
		// square
		dx = dx * dx
		sum += dx
	}

	return sum
}
