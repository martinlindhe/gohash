package gohash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"strings"
)

// ...
const (
	AllowedOnion = "abcdefghijklmnopqrstuvwxyz234567"
)

// Hasher ...
type Hasher struct {
	algo        string
	suffix      string
	expected    []byte
	minLength   int
	maxLength   int
	allowedKeys []byte
	reverse     bool
}

// NewHasher returns a new Hasher
func NewHasher() *Hasher {
	return &Hasher{}
}

// Algo sets the hash algorithm ("sha1", "sha512")
func (h *Hasher) Algo(algo string) {
	algo = strings.Replace(algo, "_", "-", -1)
	h.algo = algo
}

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

// Reverse sets wether to do sequential find in reverse order
func (h *Hasher) Reverse(b bool) {
	h.reverse = b
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

func (h *Hasher) verify() error {

	if len(h.allowedKeys) == 0 {
		return fmt.Errorf("allowedKeys unset")
	}

	if h.minLength == 0 {
		return fmt.Errorf("minLength unset")
	}

	if len(h.algo) == 0 {
		return fmt.Errorf("algo unset")
	}

	keyBitSize := len(h.expected) * 8
	expectedBitSize := len(h.expected) * 8

	// XXX use a map with algo id -> bitsize mapping
	if h.algo == "md5" && keyBitSize != 128 {
		return fmt.Errorf("expectedHash is wrong size, should be 128 bit, is %d", expectedBitSize)
	}

	if h.algo == "sha1" && keyBitSize != 160 {
		return fmt.Errorf("expectedHash is wrong size, should be 160 bit, is %d", expectedBitSize)
	}

	if h.algo == "sha224" && keyBitSize != 224 {
		return fmt.Errorf("expectedHash is wrong size, should be 224 bit, is %d", expectedBitSize)
	}

	if h.algo == "sha256" && keyBitSize != 256 {
		return fmt.Errorf("expectedHash is wrong size, should be 256 bit, is %d", expectedBitSize)
	}

	if h.algo == "sha384" && keyBitSize != 384 {
		return fmt.Errorf("expectedHash is wrong size, should be 384 bit, is %d", expectedBitSize)
	}

	if h.algo == "sha512" && keyBitSize != 512 {
		return fmt.Errorf("expectedHash is wrong size, should be 512 bit, is %d", expectedBitSize)
	}

	if h.algo == "sha512-224" && keyBitSize != 224 {
		return fmt.Errorf("expectedHash is wrong size, should be 224 bit, is %d", expectedBitSize)
	}

	if h.algo == "sha512-256" && keyBitSize != 256 {
		return fmt.Errorf("expectedHash is wrong size, should be 256 bit, is %d", expectedBitSize)
	}

	if h.algo != "md5" && h.algo != "sha1" &&
		h.algo != "sha224" && h.algo != "sha256" && h.algo != "sha384" &&
		h.algo != "sha512" && h.algo != "sha512-224" && h.algo != "sha512-256" {
		return fmt.Errorf("unknown algo %s", h.algo)
	}

	return nil
}

func (h *Hasher) equals(t []byte) bool {
	if h.algo == "md5" && byte16ArrayEquals(md5.Sum(t), h.expected) {
		return true
	}

	if h.algo == "sha1" && byte20ArrayEquals(sha1.Sum(t), h.expected) {
		return true
	}

	if h.algo == "sha224" && byte28ArrayEquals(sha256.Sum224(t), h.expected) {
		return true
	}

	if h.algo == "sha256" && byte32ArrayEquals(sha256.Sum256(t), h.expected) {
		return true
	}

	if h.algo == "sha384" && byte48ArrayEquals(sha512.Sum384(t), h.expected) {
		return true
	}

	if h.algo == "sha512" && byte64ArrayEquals(sha512.Sum512(t), h.expected) {
		return true
	}

	if h.algo == "sha512-224" && byte28ArrayEquals(sha512.Sum512_224(t), h.expected) {
		return true
	}

	if h.algo == "sha512-256" && byte32ArrayEquals(sha512.Sum512_256(t), h.expected) {
		return true
	}

	return false
}

// FindSequential calcs all possible combinations of keys of given length
func (h *Hasher) FindSequential() (string, error) {

	if err := h.verify(); err != nil {
		return "", err
	}

	tmp := make([]byte, h.minLength)

	firstAllowedKey := h.allowedKeys[0]
	lastAllowedKey := h.allowedKeys[len(h.allowedKeys)-1]

	// create initial mutation
	for x := 0; x < h.minLength; x++ {
		if h.reverse {
			tmp[x] = lastAllowedKey
		} else {
			tmp[x] = firstAllowedKey
		}
	}

	cnt := 0
	for {

		tmp2 := append(tmp, h.suffix...)

		if h.equals(tmp2) {
			return string(tmp2), nil
		}

		// update mutation
		for roller := h.minLength - 1; roller >= 0; roller-- {
			if h.reverse {
				if tmp[roller] == firstAllowedKey {
					tmp[roller] = lastAllowedKey
					continue
				} else {
					tmp[roller] = h.prevValueFor(tmp[roller])
					break
				}
			} else {
				if tmp[roller] == lastAllowedKey {
					tmp[roller] = firstAllowedKey
					continue
				} else {
					tmp[roller] = h.nextValueFor(tmp[roller])
					break
				}
			}
		}

		cnt++
		if cnt%1000000 == 0 {
			fmt.Println(string(tmp2), " (seq)")
		}
	}
}

// FindRandom uses random brute force to attempt to find by luck
func (h *Hasher) FindRandom() (string, error) {

	if h.reverse {
		return "", fmt.Errorf("reverse and random dont mix")
	}

	if err := h.verify(); err != nil {
		return "", err
	}

	tmp := make([]byte, h.minLength)

	firstAllowedKey := h.allowedKeys[0]
	allowedKeysLen := len(h.allowedKeys)

	// create initial mutation
	for x := 0; x < h.minLength; x++ {
		tmp[x] = firstAllowedKey
	}

	tmp = append(tmp, h.suffix...)

	cnt := 0
	for {
		// update mutation of first letters
		for roller := 0; roller < h.minLength; roller++ {
			tmp[roller] = h.allowedKeys[rand.Intn(allowedKeysLen)]
		}

		if h.equals(tmp) {
			return string(tmp), nil
		}

		cnt++
		if cnt%1000000 == 0 {
			fmt.Println(string(tmp), " (rnd)")
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

func (h *Hasher) prevValueFor(b byte) byte {

	prev := h.allowedKeys[0]
	for _, x := range h.allowedKeys {
		if x == b {
			return prev
		}
		prev = x
	}
	return prev
}
