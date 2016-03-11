package gohash

import "math"

func pow(x int, y int) uint64 {
	// NOTE: go 1.5 internal type of math.Pow only has a float version
	return uint64(math.Pow(float64(x), float64(y)))
}
