# About

Hasher in golang

If you are familiar with php's [hash()](php.net/manual/en/function.hash.php):
```
func hash(algo string, b *[]byte) string {

	calc := gohash.NewCalculator(*b)
	return hex.EncodeToString(*calc.Sum(algo))
}
```



# Hash algorithms

Set algo with `hasher --algo=<id>`, list all supported hashes
with `hasher --list-hashes`

```
id                Algorithm            key size key size
adler32           Adler-32             32 bit   4 byte
blake224          BLAKE-224            224 bit  28 byte
blake256          BLAKE-256            256 bit  32 byte
blake384          BLAKE-384            384 bit  48 byte
blake512          BLAKE-512            512 bit  64 byte
blake2b-512       BLAKE2b-512          512 bit  64 byte
blake2s-256       BLAKE2s-256          256 bit  32 byte
crc8-atm          Crc-8 (ATM)          8 bit    1 byte
crc16-ccitt       Crc-16 (CCITT)       16 bit   2 byte
crc16-ccitt-false Crc-16 (CCITT-False) 16 bit   2 byte
crc16-ibm         Crc-16 (IBM)         16 bit   2 byte
crc16-scsi        Crc-16 (SCSI)        16 bit   2 byte
crc32-ieee        Crc-32 (IEEE)        32 bit   4 byte
crc32-castagnoli  Crc-32 (Castagnoli)  32 bit   4 byte
crc32-koopman     Crc-32 (Koopman)     32 bit   4 byte
fnv1-32           FNV-1 32             32 bit   4 byte
fnv1a-32          FNV-1a 32            32 bit   4 byte
fnv1-64           FNV-1 64             64 bit   8 byte
fnv1a-64          FNV-1a 64            64 bit   8 byte
gost              GOST                 256 bit  32 byte
md2               MD2                  128 bit  16 byte
md4               MD4                  128 bit  16 byte
md5               MD5                  128 bit  16 byte
ripemd160         RIPEMD-160           160 bit  20 byte
sha1              SHA1                 160 bit  20 byte
sha224            SHA2-224             224 bit  28 byte
sha256            SHA2-256             256 bit  32 byte
sha384            SHA2-384             384 bit  48 byte
sha512            SHA2-512             512 bit  64 byte
sha512-224        SHA2-512/224         224 bit  28 byte
sha512-256        SHA2-512/256         256 bit  32 byte
sha3-224          SHA3-224             224 bit  28 byte
sha3-256          SHA3-256             256 bit  32 byte
sha3-384          SHA3-384             384 bit  48 byte
sha3-512          SHA3-512             512 bit  64 byte
shake128-256      SHA3-SHAKE128        256 bit  32 byte
shake256-512      SHA3-SHAKE256        512 bit  64 byte
siphash-2-4       SipHash-2-4          64 bit   8 byte
skein512-256      Skein-512-256        256 bit  32 byte
skein512-512      Skein-512-512        512 bit  64 byte
tiger192          Tiger                192 bit  24 byte
whirlpool         Whirlpool            512 bit  64 byte
```
