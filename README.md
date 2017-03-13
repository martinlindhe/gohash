# About

[![Travis-CI](https://api.travis-ci.org/martinlindhe/gohash.svg)](https://travis-ci.org/martinlindhe/gohash)
[![codecov.io](https://codecov.io/github/martinlindhe/gohash/coverage.svg?branch=master)](https://codecov.io/github/martinlindhe/gohash?branch=master)
[![GoDoc](https://godoc.org/github.com/martinlindhe/gohash?status.svg)](https://godoc.org/github.com/martinlindhe/gohash)
[![Go Report Card](https://goreportcard.com/badge/github.com/martinlindhe/gohash)](https://goreportcard.com/report/github.com/martinlindhe/gohash)

Command line tools and library to work with hashes and various encodings.


### Commands

[coder](cmd/coder)        encode / decode pipes or files between various encodings

[hasher](cmd/hasher)      calculate hashes from stdin or files

[findhash](cmd/findhash)  search for plaintext matching known hashes


### Install everything

	go get -u github.com/martinlindhe/gohash/...


### Library example

```go
import "github.com/martinlindhe/coder"

// pretend to be hash() from PHP
func hash(algo string, b *[]byte) string {

	calc := gohash.NewCalculator(*b)
	return hex.EncodeToString(*calc.Sum(algo))
}
```


### Hash algorithms

Set algo with `hasher <id>`, list all supported hashes
with `hasher --list-hashes`

| id                | Algorithm            | key size | key size | year |
| ----------------- | -------------------- | --------:| --------:| ---- |
| adler32           | Adler-32             | 32 bit   | 4 byte   | 1995 |
| blake224          | BLAKE-224            | 224 bit  | 28 byte  | 2008 |
| blake256          |  BLAKE-256           | 256 bit  | 32 byte  | 2008 |
| blake384          | BLAKE-384            | 384 bit  | 48 byte  | 2008 |
| blake512          | BLAKE-512            | 512 bit  | 64 byte  | 2008 |
| blake2b-512       | BLAKE2b-512          | 512 bit  | 64 byte  | 2012 |
| blake2s-256       | BLAKE2s-256          | 256 bit  | 32 byte  | 2012 |
| crc8-atm          | Crc-8 (ATM)          | 8 bit    | 1 byte   | ?    |
| crc16-ccitt       | Crc-16 (CCITT)       | 16 bit   | 2 byte   | ?    |
| crc16-ccitt-false | Crc-16 (CCITT-False) | 16 bit   | 2 byte   | ?    |
| crc16-ibm         | Crc-16 (IBM)         | 16 bit   | 2 byte   | ?    |
| crc16-scsi        | Crc-16 (SCSI)        | 16 bit   | 2 byte   | ?    |
| crc24-openpgp     | Crc-24 (OpenPGP)     | 24 bit   | 3 byte   | ?    |
| crc32-ieee        | Crc-32 (IEEE)        | 32 bit   | 4 byte   | ?    |
| crc32-castagnoli  | Crc-32 (Castagnoli)  | 32 bit   | 4 byte   | ?    |
| crc32-koopman     | Crc-32 (Koopman)     | 32 bit   | 4 byte   | ?    |
| crc64-iso         | Crc-64 (ISO)         | 64 bit   | 8 byte   | ?    |
| crc64-ecma        | Crc-64 (ECMA)        | 64 bit   | 8 byte   | ?    |
| fnv1-32           | FNV-1 32             | 32 bit   | 4 byte   | 1991 |
| fnv1a-32          | FNV-1a 32            | 32 bit   | 4 byte   | 1991 |
| fnv1-64           | FNV-1 64             | 64 bit   | 8 byte   | 1991 |
| fnv1a-64          | FNV-1a 64            | 64 bit   | 8 byte   | 1991 |
| gost              | GOST                 | 256 bit  | 32 byte  | 1994 |
| md2               | MD2                  | 128 bit  | 16 byte  | 1989 |
| md4               | MD4                  | 128 bit  | 16 byte  | 1990 |
| md5               | MD5                  | 128 bit  | 16 byte  | 1992 |
| ripemd160         | RIPEMD-160           | 160 bit  | 20 byte  | 1996 |
| sha1              | SHA1                 | 160 bit  | 20 byte  | 1995 |
| sha224            | SHA2-224             | 224 bit  | 28 byte  | 2001 |
| sha256            | SHA2-256             | 256 bit  | 32 byte  | 2001 |
| sha384            | SHA2-384             | 384 bit  | 48 byte  | 2001 |
| sha512            | SHA2-512             | 512 bit  | 64 byte  | 2001 |
| sha512-224        | SHA2-512/224         | 224 bit  | 28 byte  | 2001 |
| sha512-256        | SHA2-512/256         | 256 bit  | 32 byte  | 2001 |
| sha3-224          | SHA3-224             | 224 bit  | 28 byte  | 2015 |
| sha3-256          | SHA3-256             | 256 bit  | 32 byte  | 2015 |
| sha3-384          | SHA3-384             | 384 bit  | 48 byte  | 2015 |
| sha3-512          | SHA3-512             | 512 bit  | 64 byte  | 2015 |
| shake128-256      | SHA3-SHAKE128        | 256 bit  | 32 byte  | 2015 |
| shake256-512      | SHA3-SHAKE256        | 512 bit  | 64 byte  | 2015 |
| siphash-2-4       | SipHash-2-4          | 64 bit   | 8 byte   | 2012 |
| skein512-256      | Skein-512-256        | 256 bit  | 32 byte  | 2008? |
| skein512-512      | Skein-512-512        | 512 bit  | 64 byte  | 2008? |
| tiger192          | Tiger                | 192 bit  | 24 byte  | 1996 |
| whirlpool         | Whirlpool            | 512 bit  | 64 byte  | 2000 |


### Binary-to-text encodings

Set algo with `hasher --encoding=<id>`, list all supported encodings
with `hasher --list-encodings`

| id                | Algorithm              |
| ----------------- | ---------------------- |
| ascii85           | Ascii-85               |
| base32            | Base-32                |
| base36            | Base-36                |
| base58            | Base-58                |
| base64            | Base-64                |
| base91            | Base-91                |
| bubblebabble      | Bubble Babble          |
| binary            | Binary "1010"          |
| decimal           | Decimal "13 0 99"      |
| hex               | Hex "3f997a"           |
| hexup             | Hex "3F997A"           |
| octal             | Octal "0129 0226 0120" |
| z85               | Z85                    |


### License

Under [MIT](LICENSE)
