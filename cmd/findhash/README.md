# About

`findhash` is a command line tool that searches for plaintext matching known hashes


### Installation

    go get -u github.com/martinlindhe/gohash/cmd/findhash


### Sequential brute force

For example, try to find a hidden .onion url from cicada 3301

    findhash 36367763ab73783c7af284446c59466b4cd653239a311cb7116d4618dee09a8425893dc7500b464fdaf1672d7bef5e891c6e2274568926a49fb4f45132c2a8b4 \
    --algo=sha512 --suffix=.onion --allowed=abcdefghijklmnopqrstuvwxyz234567 --min-length=16

Use `--reverse` flag to start from the end


### Random brute force

    findhash 36367763ab73783c7af284446c59466b4cd653239a311cb7116d4618dee09a8425893dc7500b464fdaf1672d7bef5e891c6e2274568926a49fb4f45132c2a8b4 \
    --algo=sha512 --suffix=.onion --allowed=abcdefghijklmnopqrstuvwxyz234567 --min-length=16 --random


### Dictionary

    findhash 36367763ab73783c7af284446c59466b4cd653239a311cb7116d4618dee09a8425893dc7500b464fdaf1672d7bef5e891c6e2274568926a49fb4f45132c2a8b4 \
    --dictionary=dictionary.txt

Tries possible hashes based on `--hash` length
