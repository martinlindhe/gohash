package gohash

var (
	algos = map[string]int{
		"adler32":  32,
		"blake224": 224,
		"blake256": 256,
		"blake384": 384,
		"blake512": 512,
		"crc32":    32,
		"crc32c":   32,
		"crc32k":   32,
		// "gost":       256,
		"md2":       128,
		"md4":       128,
		"md5":       128,
		"ripemd160": 160,
		// "sha0":         160, // XXX todo support
		"sha1":         160,
		"sha224":       224,
		"sha256":       256,
		"sha384":       384,
		"sha512":       512,
		"sha512-224":   224,
		"sha512-256":   256,
		"sha3-224":     224,
		"sha3-256":     256,
		"sha3-384":     384,
		"sha3-512":     512,
		"shake128-256": 256,
		"shake256-512": 512,
		//"skein256-256": 256,
		//"skein512-256": 256,
		//"skein512-512": 512,
		"tiger192":  192,
		"whirlpool": 512,
	}

	// "gost":       gostEquals,

	//"skein256-256": skein256_256Equals,
	//"skein512-256": skein512_256Equals,
	//"skein512-512": skein512_512Equals,
)

/*
func gostEquals(b *[]byte, expected *[]byte) bool {
	// XXX cannot use due to packaging, see https://github.com/stargrave/gogost/issues/1

	w := gost341194.New(gost341194.SboxDefault)
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}
*/

/*
func skein256_256Equals(b *[]byte, expected *[]byte) bool {

	w := skein.NewHash(32) // XXX
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func skein512_256Equals(b *[]byte, expected *[]byte) bool {

	w := skein.NewHash(32)
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func skein512_512Equals(b *[]byte, expected *[]byte) bool {

	w := skein.NewHash(64) // XXXX dont work... xxxx
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}
*/
