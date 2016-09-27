package gohash

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Hasher is used to find cleartext for checksum in `expected`, using algorithm in `algo`
type Hasher struct {
	algo        string
	prefix      string
	suffix      string
	expected    []byte
	minLength   int
	maxLength   int
	allowedKeys []byte
	reverse     bool

	// runtime stats
	try    uint64
	tick   uint64
	buffer []byte
}

// NewHasher returns a new Hasher
func NewHasher() *Hasher {
	return &Hasher{}
}

// Algo sets the hash algorithm ("sha1", "sha512")
func (h *Hasher) Algo(algo string) {
	algo = strings.Replace(algo, "_", "-", -1)
	algo = strings.ToLower(algo)
	h.algo = algo
}

// ExpectedHash sets the expected hash
func (h *Hasher) ExpectedHash(expected string) {
	tmp, _ := decodeHex([]byte(expected))
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

// Prefix sets a fixed prefix
func (h *Hasher) Prefix(s string) {
	h.prefix = s
	panic(fmt.Errorf("TODO impl Prefix for Hasher"))
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

// GetAllowedKeys returns the allowed keys
func (h *Hasher) GetAllowedKeys() string { return string(h.allowedKeys) }

// FindSequential calcs all possible combinations of keys of given length
func (h *Hasher) FindSequential() (string, error) {

	if err := h.verify(); err != nil {
		return "", err
	}

	h.buffer = make([]byte, h.minLength)

	firstAllowedKey := h.allowedKeys[0]
	lastAllowedKey := h.allowedKeys[len(h.allowedKeys)-1]

	// create initial mutation
	for x := 0; x < h.minLength; x++ {
		if h.reverse {
			h.buffer[x] = lastAllowedKey
		} else {
			h.buffer[x] = firstAllowedKey
		}
	}

	h.buffer = append(h.buffer, h.suffix...)

	go h.statusReport()

	buf := make([]byte, len(h.buffer))
	copy(buf, h.buffer)

	for {

		if h.equals() {
			return string(buf), nil
		}

		// update mutation
		for roller := h.minLength - 1; roller >= 0; roller-- {
			if h.reverse {
				if buf[roller] == firstAllowedKey {
					buf[roller] = lastAllowedKey
					continue
				} else {
					buf[roller] = h.prevValueFor(buf[roller])
					break
				}
			} else {
				if buf[roller] == lastAllowedKey {
					buf[roller] = firstAllowedKey
					continue
				} else {
					buf[roller] = h.nextValueFor(buf[roller])
					break
				}
			}
		}

		mutex.Lock()
		copy(h.buffer, buf)
		h.try++
		mutex.Unlock()
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

	h.buffer = make([]byte, h.minLength)

	firstAllowedKey := h.allowedKeys[0]
	allowedKeysLen := len(h.allowedKeys)

	// create initial mutation
	for x := 0; x < h.minLength; x++ {
		h.buffer[x] = firstAllowedKey
	}

	h.buffer = append(h.buffer, h.suffix...)

	buf := make([]byte, len(h.buffer))
	copy(buf, h.buffer)

	go h.statusReport()

	for {
		if h.equals() {
			return string(buf), nil
		}

		// update mutation of first letters
		for roller := 0; roller < h.minLength; roller++ {
			buf[roller] = h.allowedKeys[rand.Intn(allowedKeysLen)]
		}

		mutex.Lock()
		copy(h.buffer, buf)
		h.try++
		mutex.Unlock()
	}
}

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

	if requiredBitSize, ok := algos[h.algo]; ok {
		if keyBitSize != requiredBitSize {
			return fmt.Errorf("expectedHash is wrong size, should be %d bit, is %d",
				requiredBitSize, expectedBitSize)
		}
	} else {
		return fmt.Errorf("unknown algo %s", h.algo)
	}

	return nil
}

func (h *Hasher) equals() bool {

	calc := NewCalculator(h.buffer)
	return byteArrayEquals(*calc.Sum(h.algo), h.expected)
}

func (h *Hasher) statusReport() {

	for {
		time.Sleep(1 * time.Second)

		mutex.Lock()
		h.tick++
		avg := h.try / h.tick
		fmt.Printf("%s ~%d/s %s\n", h.algo, avg, string(h.buffer))
		mutex.Unlock()
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
