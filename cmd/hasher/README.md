## About

hasher is a command line tool to calculate hashes from stdin or files


## Usage


### Hash from stdin

```
$ printf "hello" | hasher sha1
aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d  -
```


### Hash a file

```
$ hasher -i file.dat sha1
57864a11ea26b249cd63e48117852366db0737da  file.dat
```


### Available hash algorithms
```
$ hasher --list-algos
[adler32 blake224 blake256 blake2b-512 blake2s-256
 blake384 blake512 crc16-ccitt crc16-ccitt-false
 crc16-ibm crc16-scsi crc24-openpgp crc32-castagnoli
 crc32-ieee crc32-koopman crc64-ecma crc64-iso crc8-atm
 fnv1-32 fnv1-64 fnv1a-32 fnv1a-64 gost md2 md4 md5
 ripemd160 sha1 sha224 sha256 sha3-224 sha3-256
 sha3-384 sha3-512 sha384 sha512 sha512-224 sha512-256
 shake128-256 shake256-512 siphash-2-4 skein512-256
 skein512-512 tiger192 whirlpool]
```

### Available encodings

```
$ hasher --list-encodings
[ascii85 base32 base36 base58 base64 base91 binary
 bubblebabble decimal hex hexup octal z85]
```
