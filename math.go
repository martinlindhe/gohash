package gohash

import "math"

func Pow(x int, y int) int {
	// NOTE: go 1.5 internal type of math.Pow only has a float version
	return int(math.Pow(float64(x), float64(y)))
}
