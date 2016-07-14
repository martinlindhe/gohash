# TODO cmd/findhash

* --min-length  + --max-length   if not equal, brute force different lengths

* performance: if ran in random mode, spawn X goroutines (?) that work independently,
    to use max cpu by one app

* performance: if ran in seq mode, spawn 2 goroutines, one working from start and one from end.

* performance: if ran in seq mode, save snapshots to ~/.config/gohash.yml regularry


# TODO encodings

XXEncoded   https://en.wikipedia.org/wiki/Xxencoding
