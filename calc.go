package gohash

import (
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

// TODO use this in hasher and dict, move all wrappers in algos.go here

// Calculator is used to calculate hash of input cleartext
type Calculator struct {
	data []byte
}

// NewCalculator creates a new Calculator
func NewCalculator(data []byte) *Calculator {

	return &Calculator{
		data: data,
	}
}

var (
	checksummers = map[string]func(*[]byte) *[]byte{
		"adler32":  adler32Sum,
		"blake224": blake224Sum,
		"blake256": blake256Sum,
		"blake384": blake384Sum,
		"blake512": blake512Sum,
		"crc32":    crc32Sum,
		"crc32c":   crc32cSum,
		"crc32k":   crc32kSum,
		// "gost":       gostEquals,
		"md2":          md2Sum,
		"md4":          md4Sum,
		"md5":          md5Sum,
		"ripemd160":    ripemd160Sum,
		"sha1":         sha1Sum,
		"sha224":       sha224Sum,
		"sha256":       sha256Sum,
		"sha384":       sha384Sum,
		"sha512":       sha512Sum,
		"sha512-224":   sha512_224Sum,
		"sha512-256":   sha512_256Sum,
		"sha3-224":     sha3_224Sum,
		"sha3-256":     sha3_256Sum,
		"sha3-384":     sha3_384Sum,
		"sha3-512":     sha3_512Sum,
		"shake128-256": shake128_256Sum,
		"shake256-512": shake256_512Sum,
		//"skein256-256": skein256_256Equals,
		//"skein512-256": skein512_256Equals,
		//"skein512-512": skein512_512Equals,
		"tiger192":  tiger192Sum,
		"whirlpool": whirlpoolSum,
	}
)

// Sum returns the checksum
func (c *Calculator) Sum(algo string) *[]byte {

	if checksum, ok := checksummers[algo]; ok {
		return checksum(&c.data)
	}

	return nil
}

func adler32Sum(b *[]byte) *[]byte {
	i := adler32.Checksum(*b)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return &bs
}

func blake224Sum(b *[]byte) *[]byte {
	w := blake256.New224()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func blake256Sum(b *[]byte) *[]byte {
	w := blake256.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func blake384Sum(b *[]byte) *[]byte {
	w := blake512.New384()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func blake512Sum(b *[]byte) *[]byte {
	w := blake512.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func crc32Sum(b *[]byte) *[]byte {
	i := crc32.ChecksumIEEE(*b)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return &bs
}

func crc32cSum(b *[]byte) *[]byte {
	tbl := crc32.MakeTable(crc32.Castagnoli)
	i := crc32.Checksum(*b, tbl)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return &bs
}

func crc32kSum(b *[]byte) *[]byte {
	tbl := crc32.MakeTable(crc32.Koopman)
	i := crc32.Checksum(*b, tbl)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return &bs
}

func md2Sum(b *[]byte) *[]byte {
	w := md2.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func md4Sum(b *[]byte) *[]byte {
	w := md4.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func md5Sum(b *[]byte) *[]byte {
	x := md5.Sum(*b)
	res := x[:]
	return &res
}

func ripemd160Sum(b *[]byte) *[]byte {
	w := ripemd160.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func sha1Sum(b *[]byte) *[]byte {
	x := sha1.Sum(*b)
	res := x[:]
	return &res
}

func sha224Sum(b *[]byte) *[]byte {
	x := sha256.Sum224(*b)
	res := x[:]
	return &res
}

func sha256Sum(b *[]byte) *[]byte {
	x := sha256.Sum256(*b)
	res := x[:]
	return &res
}

func sha384Sum(b *[]byte) *[]byte {
	x := sha512.Sum384(*b)
	res := x[:]
	return &res
}

func sha512Sum(b *[]byte) *[]byte {
	x := sha512.Sum512(*b)
	res := x[:]
	return &res
}

func sha512_224Sum(b *[]byte) *[]byte {
	x := sha512.Sum512_224(*b)
	res := x[:]
	return &res
}

func sha512_256Sum(b *[]byte) *[]byte {
	x := sha512.Sum512_256(*b)
	res := x[:]
	return &res
}

func sha3_224Sum(b *[]byte) *[]byte {
	x := sha3.Sum224(*b)
	res := x[:]
	return &res
}

func sha3_256Sum(b *[]byte) *[]byte {
	x := sha3.Sum256(*b)
	res := x[:]
	return &res
}

func sha3_384Sum(b *[]byte) *[]byte {
	x := sha3.Sum384(*b)
	res := x[:]
	return &res
}

func sha3_512Sum(b *[]byte) *[]byte {
	x := sha3.Sum512(*b)
	res := x[:]
	return &res
}

func shake128_256Sum(b *[]byte) *[]byte {
	res := make([]byte, 32)
	sha3.ShakeSum128(res, *b)
	return &res
}

func shake256_512Sum(b *[]byte) *[]byte {
	res := make([]byte, 64)
	sha3.ShakeSum256(res, *b)
	return &res
}

func tiger192Sum(b *[]byte) *[]byte {
	w := tiger.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func whirlpoolSum(b *[]byte) *[]byte {
	w := whirlpool.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}
