package gohash

import (
	"fmt"
	"math/rand"
)

// FindMatchingOnionURL uses incremental brute force to try all combinations
func FindMatchingOnionURL(expected []byte) string {

	keys := []byte("aaaaaaaaaaaaaaaa")
	length := len(keys)

	cnt := 0
	for {
		tmp := make([]byte, length)

		// create current mutation:
		for x := 0; x < length; x++ {
			tmp[x] = keys[x]
		}

		toCheck := append(tmp, []byte(".onion")...)

		if MatchSha512(toCheck, expected) {
			fmt.Println("Matched ", string(toCheck))
			return string(toCheck)
		}

		// update mutation
		for roller := length - 1; roller >= 0; roller-- {
			if keys[roller] == 'z' {
				keys[roller] = '2'
				break
			} else if keys[roller] == '7' {
				// roll over
				keys[roller] = 'a'
				continue
			} else {
				keys[roller]++
				break
			}
		}

		cnt++
		if cnt%1000000 == 0 {
			fmt.Println(string(toCheck), " (seq)")
		}
	}
}

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

	cnt := 0
	for {
		// update mutation of first 16 letters
		for roller := 0; roller < 16; roller++ {
			tmp[roller] = allowedChars[rand.Intn(allowedCharsLen)]
		}

		if MatchSha512(tmp, expected) {
			fmt.Println("Matched ", string(tmp))
			return string(tmp)
		}

		cnt++
		if cnt%1000000 == 0 {
			fmt.Println(string(tmp), " (rnd)")
		}
	}
}
