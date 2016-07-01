# TODO cmd/coder

* ability to combine encodings:
    base64+ascii85


# TODO cmd/hasher

* --algo.  Ability to combine algorithms can be combinded by the plus character, e.g:
      "sha1+crc32",  "bsd+crc24+xor8".


# TODO cmd/findhash

* --min-length  + --max-length   if not equal, brute force different lengths

* performance: if ran in random mode, spawn X goroutines (?) that work independently,
    to use max cpu by one app

* performance: if ran in seq mode, spawn 2 goroutines, one working from start and one from end.

* performance: if ran in seq mode, save snapshots to ~/.config/gohash.yml regularry


# TODO encodings

UUEncoded   https://en.wikipedia.org/wiki/Uuencoding
    The program uudecode reverses the effect of uuencode, recreating the original binary file exactly

XXEncoded   https://en.wikipedia.org/wiki/Xxencoding

