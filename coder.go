package gohash

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	b58 "github.com/jbenet/go-base58"
	"github.com/martinlindhe/bubblebabble"
	"github.com/tilinna/z85"
)

// Coder is used to encode and decode various binary-to-text encodings
type Coder struct {
	encoding string
}

// NewCoder creates a new Coder
func NewCoder(encoding string) *Coder {

	return &Coder{
		encoding: resolveEncodingAliases(encoding),
	}
}

var (
	separator = " "

	coders = map[string]func([]byte) (string, error){
		"ascii85":      encodeASCII85,
		"base32":       encodeBase32,
		"base36":       encodeBase36,
		"base58":       encodeBase58,
		"base64":       encodeBase64,
		"bubblebabble": encodeBubbleBabble,
		"binary":       encodeBinary,
		"decimal":      encodeDecimal,
		"hex":          encodeHex,
		"hexup":        encodeHexUpper,
		"octal":        encodeOctal,
		"z85":          encodeZ85,
	}
)

// Encode encodes src into some encoding
func (c *Coder) Encode(src []byte) (string, error) {

	if coder, ok := coders[c.encoding]; ok {
		return coder(src)
	}
	return "", fmt.Errorf("unknown encoding: %s", c.encoding)
}

func encodeASCII85(src []byte) (string, error) {
	buf := make([]byte, ascii85.MaxEncodedLen(len(src)))
	n := ascii85.Encode(buf, src)
	buf = buf[0:n]
	return string(buf), nil
}

func encodeBase32(src []byte) (string, error) {
	return base32.StdEncoding.EncodeToString(src), nil
}

func encodeBase36(src []byte) (string, error) {
	return "", fmt.Errorf("FIXME finish base36 lib")
}

func encodeBase58(src []byte) (string, error) {
	return b58.Encode(src), nil
}

func encodeBase64(src []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(src), nil
}

func encodeBinary(src []byte) (string, error) {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%08b", b) + separator
	}

	return strings.TrimRight(res, separator), nil
}

func encodeBubbleBabble(src []byte) (string, error) {
	return bubblebabble.EncodeToString(src), nil
}

func encodeDecimal(src []byte) (string, error) {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%d", b) + separator
	}

	return strings.TrimRight(res, separator), nil
}

func encodeHex(src []byte) (string, error) {
	return hex.EncodeToString(src), nil
}

func encodeHexUpper(src []byte) (string, error) {
	return strings.ToUpper(hex.EncodeToString(src)), nil
}

func encodeOctal(src []byte) (string, error) {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%#o", b) + separator
	}

	return strings.TrimRight(res, separator), nil
}

func encodeZ85(src []byte) (string, error) {
	src4pad := src

	// pad size, input must be divisible by 4
	if len(src4pad)%4 != 0 {
		l := len(src4pad) + 4 - (len(src4pad) % 4)
		src4pad = make([]byte, l)
		for i, b := range src {
			src4pad[i] = b
		}
	}

	b85 := make([]byte, z85.EncodedLen(len(src4pad)))
	_, err := z85.Encode(b85, src4pad)
	if err != nil {
		return "", err
	}
	return string(b85), nil
}

func resolveEncodingAliases(s string) string {

	s = strings.ToLower(s)
	if s == "base85" {
		return "ascii85"
	}
	if s == "bb" {
		return "bubblebabble"
	}
	if s == "bin" {
		return "binary"
	}
	if s == "dec" {
		return "decimal"
	}
	if s == "base16" {
		return "hex"
	}
	if s == "oct" {
		return "octal"
	}
	return s
}
