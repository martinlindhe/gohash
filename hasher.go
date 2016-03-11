package gohash

import (
	"crypto/sha1"
	"crypto/sha512"
	"fmt"
	"math"
)

// Hasher ...
type Hasher struct {
	algo        string
	suffix      string
	expected    []byte
	minLength   int
	maxLength   int
	allowedKeys []byte
}

// NewHasher returns a new Hasher
func NewHasher() *Hasher {
	return &Hasher{}
}

// Algo sets the hash algorithm ("sha1", "sha512")
func (h *Hasher) Algo(algo string) { h.algo = algo }

// ExpectedHash sets the expected hash
func (h *Hasher) ExpectedHash(expected string) {
	tmp := hexStringToBytes(expected)
	h.expected = tmp[:]
}

// Length sets the length of key to find
func (h *Hasher) Length(len int) {
	h.minLength = len
	h.maxLength = len
}

// Suffix sets a fixed suffix
func (h *Hasher) Suffix(s string) { h.suffix = s }

// MinLength sets min length of key to find
func (h *Hasher) MinLength(len int) { h.minLength = len }

// MaxLength sets max length of key to find
func (h *Hasher) MaxLength(len int) { h.maxLength = len }

// AllowedKeys sets the allowed keys
func (h *Hasher) AllowedKeys(s string) {
	h.allowedKeys = strToDistinctByteSlice(s)
}

// GetAllowedKeys ...
func (h *Hasher) GetAllowedKeys() string { return string(h.allowedKeys) }

// SequentialFind calcs all possible combinations of keys of given length
func (h Hasher) SequentialFind() (string, error) {

	if len(h.allowedKeys) == 0 {
		return "", fmt.Errorf("allowedKeys unset")
	}

	if h.minLength == 0 {
		return "", fmt.Errorf("minLength unset")
	}

	combinations := math.Pow(float64(len(h.allowedKeys)), float64(h.minLength))

	fmt.Println("Brute-forcing", combinations, "combinations")

	tmp := make([]byte, h.minLength)

	firstAllowedKey := h.allowedKeys[0]
	lastAllowedKey := h.allowedKeys[len(h.allowedKeys)-1]

	// create initial mutation
	for x := 0; x < h.minLength; x++ {
		tmp[x] = firstAllowedKey
	}

	fmt.Println("INITIAL STATE:", string(tmp))

	cnt := 0
	for {

		// update mutation
		for roller := h.minLength - 1; roller >= 0; roller-- {
			if tmp[roller] == lastAllowedKey {
				// roll over
				tmp[roller] = firstAllowedKey
				continue
			} else {
				// XXX use a map with prepared lookup sequence for speed
				tmp[roller] = h.nextValueFor(tmp[roller])
				break
			}
		}

		tmp2 := append(tmp, h.suffix...)

		if h.algo == "sha1" && byte20ArrayEquals(sha1.Sum(tmp2), h.expected) {
			return string(tmp2), nil
		}

		if h.algo == "sha512" && byte64ArrayEquals(sha512.Sum512(tmp2), h.expected) {
			return string(tmp2), nil
		}

		cnt++
		if cnt%100000 == 0 {
			fmt.Println(string(tmp2))
		}

		if cnt > 100 {
			// return "xx", nil
		}
	}
}

func (h *Hasher) nextValueFor(b byte) byte {

	next := false
	for _, x := range h.allowedKeys {
		if next == true {
			return x
		}
		if x == b {
			next = true
		}
	}
	return '0'
}
