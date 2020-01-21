# ISSUES - 21 JAN 2020

1. go-isatty fails to compile to wasm with tinygo,
    pr sent https://github.com/mattn/go-isatty/pull/51

2. tinygo 0.11.0 fails to compile math/big:

    ../../../../../../../../usr/local/Cellar/go/1.13.6/libexec/src/math/big/float.go:559:4: interp: branch on a non-constant

    seems to be issue https://github.com/tinygo-org/tinygo/issues/437

3. didnt get console.log() from tinygo with log.Println() ?


4. with go 1.13, wasm.exports is not supported

    https://github.com/golang/go/issues/25612
