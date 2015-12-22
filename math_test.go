package gohash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {

	assert.Equal(t, 16, Pow(2, 4))
	assert.Equal(t, 65536, Pow(2, 16))
	assert.Equal(t, 1048576, Pow(2, 20))
}
