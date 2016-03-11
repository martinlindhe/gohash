# About

Hasher in golang

STATUS: priv


# TODO

* XXX 1: in dict mode it makes sense to try reverse of each input line

* XXX 2: make a hash cli tool using all this: work on stdin or files


# TODO findhash:
* performance: if ran in random mode, spawn X goroutines (?) that work independently,
    to use max cpu by one app

* performance: if ran in seq mode, spawn 2 goroutines, one working from start and one from end.

* performance: if ran in seq mode, save snapshots to ~/.config/gohash.yml regularry


* support Apache MD5 (htpasswd)


more benchmarks




# Hash sizes

Set algo with --algo=<id>

id            Algorithm           key size key size     Notes
adler32       Adler-32            32 bit   4 byte
blake224      BLAKE-224           224 bit  28 byte
blake256      BLAKE-256           256 bit  32 byte
blake384      BLAKE-384           384 bit  48 byte
blake512      BLAKE-512           512 bit  64 byte
crc32         Crc-32 (IEEE)       32 bit   4 byte       php's hash() calls this "crc32b"
crc32c        Crc-32 (Castagnoli) 32 bit   4 byte       XXX haven't verified calculations
crc32k        Crc-32 (Koopman)    32 bit   4 byte       XXX haven't verified calculations
gost          GOST                256 bit  32 byte      XXX cant use, see https://github.com/stargrave/gogost/issues/1
md2           MD2                 128 bit  16 byte
md4           MD4                 128 bit  16 byte
md5           MD5                 128 bit  16 byte
md6           MD6                   --variable--        XXX no golang impl found
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
tiger192      Tiger               192 bit  24 byte      php's hash() calls this "tiger192,3"
whirlpool     Whirlpool           512 bit  64 byte

TODO skein:
skein256-256  Skein-256-256       256 bit  32 byte      XXX
skein512-256  Skein-512-256       256 bit  32 byte      XXX
skein512-512  Skein-512-512       512 bit  64 byte      XXX




NOTE no golang impl for these ripemd forms:
ripemd128     32 789d569f08ed7055e94b4289a4195012
ripemd256     64 cc1d2594aece0a064b7aed75a57283d9490fd5705ed3d66bf9a
ripemd320     80 eb0cf45114c56a8421fbcb33430fa22e0cd607560a88bbe14ce


TODO later, port a sha0 lib to golang and add sha0 support
    http://home.utah.edu/~nahaj/ada/sha/sha-0/sha-0.tar


Algorithms: https://en.wikipedia.org/wiki/Comparison_of_cryptographic_hash_functions




  string(9) "ripemd128"
  string(9) "ripemd256"
  string(9) "ripemd320"

  string(10) "tiger128,3"
  [14]=>
  string(10) "tiger160,3"
  [15]=>
  string(10) "tiger192,3"
  [16]=>
  string(10) "tiger128,4"
  [17]=>
  string(10) "tiger160,4"
  [18]=>
  string(10) "tiger192,4"
  [19]=>
  string(6) "snefru"
  [20]=>
  string(9) "snefru256"
  [21]=>


  string(6) "fnv132"
  [27]=>
  string(7) "fnv1a32"
  [28]=>
  string(6) "fnv164"
  [29]=>
  string(7) "fnv1a64"
  [30]=>
  string(5) "joaat"
  [31]=>

  string(10) "haval128,3"
  [32]=>
  string(10) "haval160,3"
  [33]=>
  string(10) "haval192,3"
  [34]=>
  string(10) "haval224,3"
  [35]=>
  string(10) "haval256,3"
  [36]=>
  string(10) "haval128,4"
  [37]=>
  string(10) "haval160,4"
  [38]=>
  string(10) "haval192,4"
  [39]=>
  string(10) "haval224,4"
  [40]=>
  string(10) "haval256,4"
  [41]=>
  string(10) "haval128,5"
  [42]=>
  string(10) "haval160,5"
  [43]=>
  string(10) "haval192,5"
  [44]=>
  string(10) "haval224,5"
  [45]=>
  string(10) "haval256,5"
}
