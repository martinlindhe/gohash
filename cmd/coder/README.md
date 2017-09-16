# About

coder is a command line tool to encode / decode pipes or files
between various binary-to-text encodings.


## Installation

    go get -u github.com/martinlindhe/gohash/cmd/coder


### Encode

    echo "hello" | coder base64


### Decode

    cat file.base64 | coder -d base64


### Chained encode

Combine multiple encodings in one step:

    echo -n "hello" | coder -e base64+hex

This is equivalent to the following:

    echo -n "hello" | base64 -w 0 | xxd -p


### Chained decode

    echo -n "614756736247383d" | coder -d -n hex+base64

This is equivalent to the following:

    echo -n "614756736247383d" | xxd -r -p | base64 -d


## Available encodings

```
$ coder --list-encodings
[ascii85 base32 base36 base58 base64 base91 binary
 bubblebabble decimal hex hexup octal uu z85]
```
