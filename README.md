# About

Hasher in golang


If you are familiar with php's [hash()](php.net/manual/en/function.hash.php):
```
func hash(algo string, b *[]byte) string {

	calc := gohash.NewCalculator(*b)
	return hex.EncodeToString(*calc.Sum(algo))
}
```

STATUS: priv


# TODO

* XXX 1: in dict mode it makes sense to try reverse of each input line


# TODO findhash:
* performance: if ran in random mode, spawn X goroutines (?) that work independently,
    to use max cpu by one app

* performance: if ran in seq mode, spawn 2 goroutines, one working from start and one from end.

* performance: if ran in seq mode, save snapshots to ~/.config/gohash.yml regularry


more benchmarks




# Hash sizes

Set algo with --algo=<id>

id            Algorithm           key size key size     Notes
adler32       Adler-32            32 bit   4 byte
blake224      BLAKE-224           224 bit  28 byte
blake256      BLAKE-256           256 bit  32 byte
blake384      BLAKE-384           384 bit  48 byte
blake512      BLAKE-512           512 bit  64 byte
crc32         Crc-32 (IEEE)       32 bit   4 byte
crc32c        Crc-32 (Castagnoli) 32 bit   4 byte
crc32k        Crc-32 (Koopman)    32 bit   4 byte
fnv1-32       FNV-1 32            32 bit   4 byte
fnv1a-32      FNV-1a 32           32 bit   4 byte
fnv1-64       FNV-1 64            64 bit   8 byte
fnv1a-64      FNV-1a 64           64 bit   8 byte
gost          GOST                256 bit  32 byte      FIXME works with patch https://github.com/stargrave/gogost/pull/2
md2           MD2                 128 bit  16 byte
md4           MD4                 128 bit  16 byte
md5           MD5                 128 bit  16 byte
ripemd160     RIPEMD-160          160 bit  20 byte
sha1          SHA1                160 bit  20 byte
sha224        SHA2-224            224 bit  28 byte
sha256        SHA2-256            256 bit  32 byte
sha384        SHA2-384            384 bit  48 byte
sha512        SHA2-512            512 bit  64 byte
sha512-224    SHA2-512/224        224 bit  28 byte
sha512-256    SHA2-512/256        256 bit  32 byte
sha3-224      SHA3-224            224 bit  28 byte
sha3-256      SHA3-256            256 bit  32 byte
sha3-384      SHA3-384            384 bit  48 byte
sha3-512      SHA3-512            512 bit  64 byte
shake128-256  SHA3-SHAKE128       256 bit  32 byte
shake256-512  SHA3-SHAKE256       512 bit  64 byte
siphash-2-4   SipHash-2-4         64 bit   8 byte
tiger192      Tiger               192 bit  24 byte
whirlpool     Whirlpool           512 bit  64 byte

TODO skein:
skein256-256  Skein-256-256       256 bit  32 byte      XXX
skein512-256  Skein-512-256       256 bit  32 byte      XXX
skein512-512  Skein-512-512       512 bit  64 byte      XXX



NOTE no golang impl for these ripemd forms:
ripemd128     RIPEMD-128          128 bit  16 byte
ripemd256     RIPEMD-256          256 bit  32 byte
ripemd320     RIPEMD-320          320 bit  40 byte

TODO later, sha0:
sha0          SHA0                160 bit  20 byte      XXX no golang impl found

TODO later, md6:
md6           MD6                   --variable--        XXX no golang impl found


Algorithms: https://en.wikipedia.org/wiki/Comparison_of_cryptographic_hash_functions



TODO JH, sha3-finalist, https://en.wikipedia.org/wiki/JH_(hash_function)

TODO Grøstl, sha3-finalist, https://en.wikipedia.org/wiki/Gr%C3%B8stl
    https://github.com/ctz/groestl/blob/master/groestl.py

TODO never(?), panama, https://en.wikipedia.org/wiki/Panama_(cryptography)

TODO never(?), ECOH, sha3-disqualified, https://en.wikipedia.org/wiki/Elliptic_curve_only_hash


TODO later, tiger variants (from php):
tiger128,3
tiger160,3
tiger192,3
tiger128,4
tiger160,4
tiger192,4


TODO never(?), snefru (from php), https://en.wikipedia.org/wiki/Snefru
snefru
snefru256

TODO never(?), haval (from php), https://en.wikipedia.org/wiki/HAVAL
haval128,3
haval160,3
haval192,3
haval224,3
haval256,3
haval128,4
haval160,4
haval192,4
haval224,4
haval256,4
haval128,5
haval160,5
haval192,5
haval224,5
haval256,5
