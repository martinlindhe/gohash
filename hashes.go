package gohash

import "crypto/sha1"

type HashComparer func(val []byte, expected []byte) bool

// see if hash of computed value match expected value
func MatchSha1(data []byte, expected []byte) bool {
	return ByteArrayEquals(sha1.Sum(data), expected)
}

// calc all possible combinations of keys of given length
func FindMatchingHash(keys []byte, comparer HashComparer, expected []byte, length int) []byte {

	pos := make([]byte, length)

	numKeys := byte(len(keys))

	iterations := Pow(len(keys), length)

	for i := 0; i < iterations; i++ {
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
			if pos[roller] < numKeys {
				break
			}
			pos[roller] = 0
		}
	}
	return nil
}
