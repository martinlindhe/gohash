package gohash

import "math"

func Pow(x int, y int) int64 {
	// NOTE: go 1.5 internal type of math.Pow only has a float version
	return int64(math.Pow(float64(x), float64(y)))
}
