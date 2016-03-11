package gohash

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcAdler32(t *testing.T) {

	// from https://en.wikipedia.org/wiki/Adler-32#Example
	calc := NewCalculator([]byte("Wikipedia"))
	assert.Equal(t, "11e60398", hex.EncodeToString(*calc.Sum("adler32")))
}

func TestCalcBlake(t *testing.T) {

	calc := NewCalculator([]byte("The quick brown fox jumps over the lazy dog"))

	assert.Equal(t, "c8e92d7088ef87c1530aee2ad44dc720cc10589cc2ec58f95a15e51b",
		hex.EncodeToString(*calc.Sum("blake224")))

	assert.Equal(t, "7576698ee9cad30173080678e5965916adbb11cb5245d386bf1ffda1cb26c9d7",
		hex.EncodeToString(*calc.Sum("blake256")))

	assert.Equal(t, "67c9e8ef665d11b5b57a1d99c96adffb3034d8768c0827d1c6e60b54871e8673651767a2c6c43d0ba2a9bb2500227406",
		hex.EncodeToString(*calc.Sum("blake384")))

	// from https://en.wikipedia.org/wiki/BLAKE_(hash_function)#BLAKE_hashes
	assert.Equal(t, "1f7e26f63b6ad25a0896fd978fd050a1766391d2fd0471a77afb975e5034b7ad2d9ccf8dfb47abbbe656e1b82fbc634ba42ce186e8dc5e1ce09a885d41f43451",
		hex.EncodeToString(*calc.Sum("blake512")))
}

func TestCalcCrc32(t *testing.T) {

	calc := NewCalculator([]byte("The quick brown fox jumps over the lazy dog"))

	assert.Equal(t, "414fa339",
		hex.EncodeToString(*calc.Sum("crc32"))) // aka "crc32b" in php

	assert.Equal(t, "22620404",
		hex.EncodeToString(*calc.Sum("crc32c")))

	assert.Equal(t, "e021db90",
		hex.EncodeToString(*calc.Sum("crc32k")))

	// NOTE: none of these crc32 implementations seem to correspond to the php "crc32" one
}

func TestCalcMd2(t *testing.T) {

	// from https://en.wikipedia.org/wiki/MD2_(cryptography)#MD2_hashes
	calc := NewCalculator([]byte("The quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "03d85a0d629d2c442e987525319fc471", hex.EncodeToString(*calc.Sum("md2")))
}

func TestCalcMd4(t *testing.T) {

	// from https://en.wikipedia.org/wiki/MD4#MD4_hashes
	calc := NewCalculator([]byte("The quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "1bee69a46ba811185c194762abaeae90", hex.EncodeToString(*calc.Sum("md4")))
}

func TestCalcMd5(t *testing.T) {

	// from https://en.wikipedia.org/wiki/MD5#MD5_hashes
	calc := NewCalculator([]byte("The quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "9e107d9d372bb6826bd81d3542a419d6", hex.EncodeToString(*calc.Sum("md5")))
}

func TestCalcRipemd160(t *testing.T) {

	// from https://en.wikipedia.org/wiki/RIPEMD#RIPEMD-160_hashes
	calc := NewCalculator([]byte("The quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "37f332f68db77bd9d7edd4969571ad671cf9dd3b", hex.EncodeToString(*calc.Sum("ripemd160")))
}

func TestCalcSha1(t *testing.T) {

	// from https://en.wikipedia.org/wiki/SHA-1#Example_hashes
	calc := NewCalculator([]byte("The quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12", hex.EncodeToString(*calc.Sum("sha1")))
}

func TestCalcSha2(t *testing.T) {

	// from https://en.wikipedia.org/wiki/SHA-2#Examples_of_SHA-2_variants
	calc := NewCalculator([]byte(""))
	assert.Equal(t, "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f",
		hex.EncodeToString(*calc.Sum("sha224")))

	assert.Equal(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		hex.EncodeToString(*calc.Sum("sha256")))

	assert.Equal(t, "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b",
		hex.EncodeToString(*calc.Sum("sha384")))

	assert.Equal(t, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
		hex.EncodeToString(*calc.Sum("sha512")))

	assert.Equal(t, "6ed0dd02806fa89e25de060c19d3ac86cabb87d6a0ddd05c333b84f4",
		hex.EncodeToString(*calc.Sum("sha512-224")))

	assert.Equal(t, "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a",
		hex.EncodeToString(*calc.Sum("sha512-256")))
}

func TestCalcSha3(t *testing.T) {

	// from https://en.wikipedia.org/wiki/SHA-3#Examples_of_SHA-3_variants

	calc := NewCalculator([]byte(""))
	assert.Equal(t, "6b4e03423667dbb73b6e15454f0eb1abd4597f9a1b078e3f5b5a6bc7",
		hex.EncodeToString(*calc.Sum("sha3-224")))

	assert.Equal(t, "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a",
		hex.EncodeToString(*calc.Sum("sha3-256")))

	assert.Equal(t, "0c63a75b845e4f7d01107d852e4c2485c51a50aaaa94fc61995e71bbee983a2ac3713831264adb47fb6bd1e058d5f004",
		hex.EncodeToString(*calc.Sum("sha3-384")))

	assert.Equal(t, "a69f73cca23a9ac5c8b567dc185a756e97c982164fe25859e0d1dcc1475c80a615b2123af1f5f94c11e3e9402c3ac558f500199d95b6d3e301758586281dcd26",
		hex.EncodeToString(*calc.Sum("sha3-512")))

	assert.Equal(t, "7f9c2ba4e88f827d616045507605853ed73b8093f6efbc88eb1a6eacfa66ef26",
		hex.EncodeToString(*calc.Sum("shake128-256")))

	assert.Equal(t, "46b9dd2b0ba88d13233b3feb743eeb243fcd52ea62b81b82b50c27646ed5762fd75dc4ddd8c0f200cb05019d67b592f6fc821c49479ab48640292eacb3b7c4be",
		hex.EncodeToString(*calc.Sum("shake256-512")))
}

func TestCalcTiger192(t *testing.T) {

	// from https://en.wikipedia.org/wiki/Tiger_(cryptography)#Examples
	calc := NewCalculator([]byte("The quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "6d12a41e72e644f017b6f0e2f7b44c6285f06dd5d2c5b075", hex.EncodeToString(*calc.Sum("tiger192")))
}

func TestCalcWhirlpool(t *testing.T) {

	// from https://en.wikipedia.org/wiki/Whirlpool_(cryptography)#Whirlpool_hashes
	calc := NewCalculator([]byte("The quick brown fox jumps over the lazy dog"))
	assert.Equal(t, "b97de512e91e3828b40d2b0fdce9ceb3c4a71f9bea8d88e75c4fa854df36725fd2b52eb6544edcacd6f8beddfea403cb55ae31f03ad62a5ef54e42ee82c3fb35", hex.EncodeToString(*calc.Sum("whirlpool")))
}
