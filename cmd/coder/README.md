# About

coder is a command line tool to encode / decode pipes or files
between various binary-to-text encodings.


## Installation

    go get -u github.com/martinlindhe/gohash/cmd/coder


### Encode

    printf "hello" | coder base64


### Decode

    cat file.base64 | coder base64 -d


## Available encodings

```
$ coder --list-encodings
[ascii85 base32 base36 base58 base64 base91 binary
 bubblebabble decimal hex hexup octal z85]
```
