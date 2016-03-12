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
		encoding: encoding,
	}
}

// Encode encodes src into some encoding
func (e *Coder) Encode(src []byte) (string, error) {

	switch e.encoding {
	case "ascii85", "base85":
		buf := make([]byte, ascii85.MaxEncodedLen(len(src)))
		n := ascii85.Encode(buf, src)
		buf = buf[0:n]
		return string(buf), nil

	case "base32":
		return base32.StdEncoding.EncodeToString(src), nil

	case "base36":
		return "XXX", nil // FIXME finish base36 lib

	case "base58":
		return b58.Encode(src), nil

	case "base64":
		return base64.StdEncoding.EncodeToString(src), nil

	case "bubblebabble", "bb":
		return bubblebabble.EncodeToString(src), nil

	case "binary", "bin":
		return toBinaryString(src, " "), nil

	case "decimal", "dec":
		return toDecimalString(src, " "), nil

	case "hex", "base16":
		return hex.EncodeToString(src), nil

	case "hexup":
		return strings.ToUpper(hex.EncodeToString(src)), nil

	case "octal", "oct":
		return toOctalString(src, " "), nil

	case "z85":
		src4pad := src
		if len(src4pad)%4 != 0 {
			// The len(src) must be divisible by 4
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

	return "", fmt.Errorf("error: unknown --encoding %s", e.encoding)
}

func toBinaryString(src []byte, separator string) string {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%08b", b) + separator
	}

	return strings.TrimRight(res, separator)
}

func toOctalString(src []byte, separator string) string {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%#o", b) + separator
	}

	return strings.TrimRight(res, separator)
}

func toDecimalString(src []byte, separator string) string {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%d", b) + separator
	}

	return strings.TrimRight(res, separator)
}
