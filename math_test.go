package gohash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {

	assert.Equal(t, uint64(16), pow(2, 4))
	assert.Equal(t, uint64(65536), pow(2, 16))
	assert.Equal(t, uint64(1048576), pow(2, 20))
}
