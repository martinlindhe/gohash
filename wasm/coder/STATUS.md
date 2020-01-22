# ISSUES - 21 JAN 2020

1. tinygo 0.11.0 fails to compile math/big:

    ../../../../../../../../usr/local/Cellar/go/1.13.6/libexec/src/math/big/float.go:559:4: interp: branch on a non-constant

    https://tinygo.org/lang-support/stdlib/#math-big
    https://github.com/tinygo-org/tinygo/issues/437


2. with go 1.13, wasm.exports is not supported

    https://github.com/golang/go/issues/25612
