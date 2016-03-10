package gohash

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestBruteforceUrlFromHashOne(t *testing.T) {

	hashString := "c4ecb61423d366a82fefbe64a28620027293006a7b9384b1b0389f50547c428b62850f1ec542406135dffbfbfefac1e86241b5620ed7d7d219ae77f6f7464231"
	expectedHash := HexStringToBytes(hashString)

	assert.Equal(t, "aaaaaaaaaaaaaaaf.onion", FindMatchingOnionURL(expectedHash))
}

func TestBruteforceUrlFromHash(t *testing.T) {

	// NOTE from cicada 3301 http://i.imgur.com/H4HbfDW.jpg

	// 128 chars = 64 bytes = 512 bit hash, so either sha2-512 or sha3 something
	// lets try sha2-512
	hashString := "36367763ab73783c7af284446c59466b4cd653239a311cb7116d4618dee09a8425893dc7500b464fdaf1672d7bef5e891c6e2274568926a49fb4f45132c2a8b4"
	expectedHash := HexStringToBytes(hashString)

	// NOTE: should end with .onion

	// XXX onion links look like "zohurs2pymumjgfu.onion"

	// Addresses in the .onion TLD are generally opaque, non-mnemonic,
	// 16-character alpha-semi-numeric hashes which are automatically
	// generated based on a public key when a hidden service is configured.
	// These 16-character hashes can be made up of any letter of the alphabet,
	// and decimal digits from 2 to 7, thus representing an 80-bit number in base32.

	chk := FindMatchingOnionURL(expectedHash)
	spew.Dump(chk)
}
