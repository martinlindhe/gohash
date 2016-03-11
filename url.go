package gohash

import (
	"crypto/sha512"
	"fmt"
	"math/rand"
)

// FindMatchingOnionURLByRandom uses random brute force to attempt to find by luck
func FindMatchingOnionURLByRandom(expected []byte) string {

	allowedChars := []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'2', '3', '4', '5', '6', '7',
	}

	allowedCharsLen := len(allowedChars)

	keys := []byte("aaaaaaaaaaaaaaaa.onion")
	length := len(keys)

	tmp := make([]byte, length)

	// create initial mutation
	for x := 0; x < length; x++ {
		tmp[x] = keys[x]
	}

	fmt.Println("INITIAL STATE:", string(tmp))

	cnt := 0
	for {
		// update mutation of first 16 letters
		for roller := 0; roller < 16; roller++ {
			tmp[roller] = allowedChars[rand.Intn(allowedCharsLen)]
		}

		if byte64ArrayEquals(sha512.Sum512(tmp), expected) {
			fmt.Println("Matched ", string(tmp))
			return string(tmp)
		}

		cnt++
		if cnt%1000000 == 0 {
			fmt.Println(string(tmp), " (rnd)")
		}
	}
}
