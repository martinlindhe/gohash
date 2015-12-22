package gohash

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// find hash of some string
func TestFindSha1Hash(t *testing.T) {

	keys := StrToDistinctByteSlice("tom")

	length := 5

	// sha1 of "motom":
	expectedHash := HexStringToBytes("9af7d87edaba03e23f6dbdaed29101ee1291c8a6")

	chk := FindMatchingHash(keys, MatchSha1, expectedHash, length)

	assert.Equal(t, "motom", string(chk))
}

// benchmarks given key length and print a prediction based on it
func BenchmarkSha1Speed(*testing.B) {

	keys := StrToDistinctByteSlice("580%(=QWI+qwi*Nn")

	length := 5

	expectedHash := HexStringToBytes("0000000000000000000000000000000000000000")

	iterations := Pow(len(keys), length)
	start := time.Now()
	FindMatchingHash(keys, MatchSha1, expectedHash, length)
	duration := time.Since(start)

	//	fmt.Printf("[benchmark length %d, %d iterations took %s]\n", length, iterations, duration)

	oneOperation := duration / time.Duration(iterations)

	for i := 6; i <= 10; i++ {
		it := Pow(len(keys), i)
		predictedTime := oneOperation * time.Duration(it)

		fmt.Printf("length = %2d, iterations = %-15d, predicted time = %s\n", i, it, predictedTime)
	}
}
