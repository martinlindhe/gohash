package gohash

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"hash/adler32"
	"hash/crc32"

	"github.com/cxmcc/tiger"
	"github.com/dchest/blake256"
	"github.com/dchest/blake512"
	"github.com/htruong/go-md2"
	"github.com/jzelinskie/whirlpool"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

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
		"md2":        128,
		"md4":        128,
		"md5":        128,
		"ripemd160":  160,
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
		"shake128":   256,
		"shake256":   512,
		"tiger192":   192,
		"whirlpool":  512,
	}

	algoEquals = map[string]func(*[]byte, *[]byte) bool{
		"adler32":  adler32Equals,
		"blake224": blake224Equals,
		"blake256": blake256Equals,
		"blake384": blake384Equals,
		"blake512": blake512Equals,
		"crc32":    crc32Equals,
		"crc32c":   crc32cEquals,
		"crc32k":   crc32kEquals,
		// "gost":       gostEquals,
		"md2":        md2Equals,
		"md4":        md4Equals,
		"md5":        md5Equals,
		"ripemd160":  ripemd160Equals,
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
		"shake128":   shake128Equals,
		"shake256":   shake256Equals,
		"tiger192":   tiger192Equals,
		"whirlpool":  whirlpoolEquals,
	}
)

func adler32Equals(b *[]byte, expected *[]byte) bool {

	var expectedInt uint32
	_ = binary.Read(bytes.NewReader(*expected), binary.BigEndian, &expectedInt)
	return adler32.Checksum(*b) == expectedInt
}

func blake224Equals(b *[]byte, expected *[]byte) bool {

	w := blake256.New224()
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func blake256Equals(b *[]byte, expected *[]byte) bool {

	w := blake256.New()
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func blake384Equals(b *[]byte, expected *[]byte) bool {

	w := blake512.New384()
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func blake512Equals(b *[]byte, expected *[]byte) bool {

	w := blake512.New()
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func crc32Equals(b *[]byte, expected *[]byte) bool {

	var expectedInt uint32
	_ = binary.Read(bytes.NewReader(*expected), binary.BigEndian, &expectedInt)
	return crc32.ChecksumIEEE(*b) == expectedInt
}

func crc32cEquals(b *[]byte, expected *[]byte) bool {

	var expectedInt uint32
	_ = binary.Read(bytes.NewReader(*expected), binary.BigEndian, &expectedInt)

	tbl := crc32.MakeTable(crc32.Castagnoli)
	return crc32.Checksum(*b, tbl) == expectedInt
}

func crc32kEquals(b *[]byte, expected *[]byte) bool {

	var expectedInt uint32
	_ = binary.Read(bytes.NewReader(*expected), binary.BigEndian, &expectedInt)

	tbl := crc32.MakeTable(crc32.Koopman)
	return crc32.Checksum(*b, tbl) == expectedInt
}

/*
func gostEquals(b *[]byte, expected *[]byte) bool {
	// XXX cannot use due to packaging, see https://github.com/stargrave/gogost/issues/1

	w := gost341194.New(gost341194.SboxDefault)
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}
*/

func md2Equals(b *[]byte, expected *[]byte) bool {

	w := md2.New()
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func md4Equals(b *[]byte, expected *[]byte) bool {

	w := md4.New()
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func md5Equals(b *[]byte, expected *[]byte) bool {
	return byte16ArrayEquals(md5.Sum(*b), *expected)
}

func sha1Equals(b *[]byte, expected *[]byte) bool {
	return byte20ArrayEquals(sha1.Sum(*b), *expected)
}

func sha224Equals(b *[]byte, expected *[]byte) bool {
	return byte28ArrayEquals(sha256.Sum224(*b), *expected)
}
func sha256Equals(b *[]byte, expected *[]byte) bool {
	return byte32ArrayEquals(sha256.Sum256(*b), *expected)
}

func sha384Equals(b *[]byte, expected *[]byte) bool {
	return byte48ArrayEquals(sha512.Sum384(*b), *expected)
}

func sha512Equals(b *[]byte, expected *[]byte) bool {
	return byte64ArrayEquals(sha512.Sum512(*b), *expected)
}

func sha512_224Equals(b *[]byte, expected *[]byte) bool {
	return byte28ArrayEquals(sha512.Sum512_224(*b), *expected)
}

func sha512_256Equals(b *[]byte, expected *[]byte) bool {
	return byte32ArrayEquals(sha512.Sum512_256(*b), *expected)
}

func sha3_224Equals(b *[]byte, expected *[]byte) bool {
	return byte28ArrayEquals(sha3.Sum224(*b), *expected)
}

func sha3_256Equals(b *[]byte, expected *[]byte) bool {
	return byte32ArrayEquals(sha3.Sum256(*b), *expected)
}

func sha3_384Equals(b *[]byte, expected *[]byte) bool {
	return byte48ArrayEquals(sha3.Sum384(*b), *expected)
}

func sha3_512Equals(b *[]byte, expected *[]byte) bool {
	return byte64ArrayEquals(sha3.Sum512(*b), *expected)
}

func shake128Equals(b *[]byte, expected *[]byte) bool {

	h := make([]byte, 32)
	sha3.ShakeSum128(h, *b)
	return byteArrayEquals(h, *expected)
}

func shake256Equals(b *[]byte, expected *[]byte) bool {

	h := make([]byte, 64)
	sha3.ShakeSum256(h, *b)
	return byteArrayEquals(h, *expected)
}

func tiger192Equals(b *[]byte, expected *[]byte) bool {

	w := tiger.New()
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func whirlpoolEquals(b *[]byte, expected *[]byte) bool {

	w := whirlpool.New()
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}

func ripemd160Equals(b *[]byte, expected *[]byte) bool {

	w := ripemd160.New()
	w.Write(*b)
	return byteArrayEquals(w.Sum(nil), *expected)
}
