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

func TestDictionaryCiccada(t *testing.T) {

	dict, err := NewDictionary("data/onion-sites.txt")
	assert.Equal(t, nil, err)

	// unknown hash from ciccada 3301, 2014onion7 56.jpg
	dict.ExpectedHash("36367763ab73783c7af284446c59466b4cd653239a311cb7116d4618dee09a8425893dc7500b464fdaf1672d7bef5e891c6e2274568926a49fb4f45132c2a8b4")

	res, algo, err := dict.Find()
	assert.Equal(t, nil, err)
	assert.Equal(t, "", algo)
	assert.Equal(t, "", string(res)) // XXX no result known
}
