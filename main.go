package main

import (
	"fmt"
	"math"

	"hippo/vector"
)

func main() {
	vecSize := 8000
	numVecotrs := 10000
	vectors := make([]vector.Vector, 0, numVecotrs)

	for range numVecotrs {
		vectors = append(vectors, vector.NewRandom(vecSize))
	}

	example := vector.NewRandom(vecSize)
	closest := float32(math.Inf(1))
	res := 0
	for i := range vectors {
		dist := vector.GetL2(example, vectors[i])
		if dist < closest {
			res = i
			closest = dist
		}
	}
	fmt.Println(vectors[res])
}
