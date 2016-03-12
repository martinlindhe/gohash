package gohash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionary(t *testing.T) {

	dict, err := NewDictionary("data/onion-sites.txt")
	assert.Equal(t, nil, err)

	// sha256 of a row in onion-sites.txt, lets find it
	dict.ExpectedHash("be80cf12710923248c15649dbe44012623708399e34add8f1d5cb89bf6f96299")

	res, algo, err := dict.Find()
	assert.Equal(t, nil, err)
	assert.Equal(t, "sha256", algo)
	assert.Equal(t, "3qr42dbkhrjp55kg.onion", string(res))
}
