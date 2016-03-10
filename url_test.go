package gohash

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBruteforceUrlFromHashSequential(t *testing.T) {

	hashString := "c4ecb61423d366a82fefbe64a28620027293006a7b9384b1b0389f50547c428b62850f1ec542406135dffbfbfefac1e86241b5620ed7d7d219ae77f6f7464231"
	expectedHash := hexStringToBytes(hashString)

	assert.Equal(t, "aaaaaaaaaaaaaaaf.onion", FindMatchingOnionURL(expectedHash))
}

func TestBruteforceUrlFromHashRandomly(t *testing.T) {

	rand.Seed(123)

	hashString := "c470e64945e2c11f55349b68a7983b79a8b695f0c1011a5389104f93c3c93a837dcf033a71c09c7bd2b3eb4763a25389682a8af0c3f0e8a530b343bb8ce34d52"
	expectedHash := hexStringToBytes(hashString)

	// XXX dont find it but should!
	assert.Equal(t, "dnmntrrxx224jknh.onion", FindMatchingOnionURLByRandom(expectedHash))
}
