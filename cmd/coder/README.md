# About

coder is a command line tool to encode / decode pipes or files
between various binary-to-text encodings.


## Installation

    go get -u github.com/martinlindhe/gohash/cmd/coder


### Encode

    printf "hello" | coder base64


### Decode

    cat file.base64 | coder base64 -d


### Chained encodings

This should work on Linux
    $ echo "hello" | base64 | xxd -p
    614756736247384b0a

Decode
    $ echo "614756736247384b0a" | xxd -r -p | base64 -d
    hello

## Available encodings

```
$ coder --list-encodings
[ascii85 base32 base36 base58 base64 base91 binary
 bubblebabble decimal hex hexup octal z85]
```
