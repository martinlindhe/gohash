package gohash

import (
	"crypto/sha1"
	"crypto/sha512"
)

// HashComparer ...
type HashComparer func(val []byte, expected []byte) bool

// MatchSha1 checks if two hashes match
func MatchSha1(data []byte, expected []byte) bool {
	return ByteArrayEquals(sha1.Sum(data), expected)
}

// MatchSha512 checks if two hashes match
func MatchSha512(data []byte, expected []byte) bool {
	return Byte64ArrayEquals(sha512.Sum512(data), expected)
}

// FindMatchingHash calcs all possible combinations of keys of given length
func FindMatchingHash(keys []byte, comparer HashComparer, expected []byte, length int) []byte {

	pos := make([]byte, length)

	numKeys := byte(len(keys))

	iterations := Pow(len(keys), length)

	for i := int64(0); i < iterations; i++ {
		tmp := make([]byte, length)

		// create current mutation:
		for x := 0; x < length; x++ {
			tmp[x] = keys[pos[x]]
		}

		if comparer(tmp, expected) {
			return tmp
		}

		// update mutation
		for roller := length - 1; roller >= 0; roller-- {
			pos[roller]++
			if pos[roller] < numKeys { // XXX ???
				break
			}
			pos[roller] = 0
		}
	}
	return nil
}
