== README

Hasher in golang


== TODO

* if ran in random mode, spawn X goroutines (?) that work independently,
    to use max cpu by one app

* if ran in seq mode, spawn 2 goroutines, one working from start and one from end.

* if ran in seq mode, save snapshots to ~/.config/gohash.yml regularry


* support Apache MD5 (htpasswd)


more benchmarks







# Hash sizes


id      Algorithm   keysize  keysize
md5     MD5         128 bit  16 byte
sha1    SHA1        160 bit  20 byte
sha224  SHA2-224    224 bit  28 byte
sha256  SHA2-256    256 bit  32 byte
sha384  SHA2-384    384 bit  48 byte
sha512  SHA2-512    512 bit  64 byte
