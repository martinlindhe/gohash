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

md5     128 bit     16 byte
sha1    160 bit     20 byte
sha512  512 bit     64 byte
