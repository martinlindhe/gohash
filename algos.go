package gohash

var (
	algos = map[string]int{
		"md5":        128,
		"sha1":       160,
		"sha224":     224,
		"sha256":     256,
		"sha384":     384,
		"sha512":     512,
		"sha512-224": 224,
		"sha512-256": 256,
		"sha3-224":   224,
		"sha3-256":   256,
		"sha3-384":   384,
		"sha3-512":   512,
	}

	algoEquals = map[string]func(*[]byte, *[]byte) bool{
		"md5":        md5Equals,
		"sha1":       sha1Equals,
		"sha224":     sha224Equals,
		"sha256":     sha256Equals,
		"sha384":     sha384Equals,
		"sha512":     sha512Equals,
		"sha512-224": sha512_224Equals,
		"sha512-256": sha512_256Equals,
		"sha3-224":   sha3_224Equals,
		"sha3-256":   sha3_256Equals,
		"sha3-384":   sha3_384Equals,
		"sha3-512":   sha3_512Equals,
	}
)
