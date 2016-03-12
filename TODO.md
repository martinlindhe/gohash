# TODO XXX

hasher_test.go: rewrite to use a table like calc_test.go


# TODO hasher:



    --algo.  Algorithms can be combinded by the plus  character,  e.  g.
              "sha1+crc32",  "bsd+crc24+xor8".


              -e --encoding, output encoding:

                         bin          Binary
                         dec          Decimal
                         oct          Octal
                         hex          Hexadecimal in lowercase (same as -x)
                         hexup        Hexadecimal in uppercase (same as -X)
                         base16       Base 16 (as defined by RFC 3548)
                         base32       Base 32 (as defined by RFC 3548)
                         base64       Base 64 (as defined by RFC 3548)
                         bb           BubbleBabble (used by OpenSSH and SSH2)
                         z85          ZeroMQ Base-85



# TODO findhash:


--min-length  + --max-length   if not equal, brute force different lengths



* performance: if ran in random mode, spawn X goroutines (?) that work independently,
    to use max cpu by one app

* performance: if ran in seq mode, spawn 2 goroutines, one working from start and one from end.

* performance: if ran in seq mode, save snapshots to ~/.config/gohash.yml regularry


more benchmarks





# TODO - unsupported hashes

NOTE no golang impl for these ripemd forms:
ripemd128     RIPEMD-128          128 bit  16 byte
ripemd256     RIPEMD-256          256 bit  32 byte
ripemd320     RIPEMD-320          320 bit  40 byte

TODO later, sha0:
sha0          SHA0                160 bit  20 byte      XXX no golang impl found

TODO later, md6:
md6           MD6                   --variable--        XXX no golang impl found




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

TODO later, skein:
skein256-256  Skein-256-256       256 bit  32 byte
    not available in github.com/dchest/skein
    also Skein-1024-384, and more forms



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
