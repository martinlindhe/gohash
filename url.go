package gohash

import "fmt"

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
		if cnt%100000 == 0 {
			fmt.Println(string(toCheck))
		}
	}
}
