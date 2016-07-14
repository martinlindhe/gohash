package gohash

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/bproctor/base91"
	b58 "github.com/jbenet/go-base58"
	"github.com/martinlindhe/base36"
	"github.com/martinlindhe/bubblebabble"
	"github.com/tejainece/uu"
	"github.com/tilinna/z85"
)

// Coder is used to encode and decode various binary-to-text encodings
type Coder struct {
	encoding string
}

var (
	separator = " "
	encoders  = map[string]func([]byte) ([]byte, error){
		"ascii85":      encodeASCII85,
		"base32":       encodeBase32,
		"base36":       encodeBase36,
		"base58":       encodeBase58,
		"base64":       encodeBase64,
		"base91":       encodeBase91,
		"bubblebabble": encodeBubbleBabble,
		"binary":       encodeBinary,
		"decimal":      encodeDecimal,
		"hex":          encodeHex,
		"hexup":        encodeHexUpper,
		"octal":        encodeOctal,
		"uu":           encodeUU,
		"z85":          encodeZ85,
	}

	decoders = map[string]func(string) ([]byte, error){
		"ascii85":      decodeASCII85,
		"base32":       decodeBase32,
		"base36":       decodeBase36,
		"base58":       decodeBase58,
		"base64":       decodeBase64,
		"base91":       decodeBase91,
		"binary":       decodeBinary,
		"bubblebabble": decodeBubbleBabble,
		"decimal":      decodeDecimal,
		"hex":          decodeHex,
		"hexup":        decodeHex,
		"octal":        decodeOctal,
		"uu":           decodeUU,
		"z85":          decodeZ85,
	}
)

// NewCoder creates a new Coder
func NewCoder(encoding string) *Coder {

	return &Coder{
		encoding: resolveEncodingAliases(encoding),
	}
}

// Encode encodes src into some encoding
func (c *Coder) Encode(src []byte) ([]byte, error) {

	if coder, ok := encoders[c.encoding]; ok {
		return coder(src)
	}
	return nil, fmt.Errorf("unknown encoding: %s", c.encoding)
}

// Decode decodes src from some encoding
func (c *Coder) Decode(src string) ([]byte, error) {

	if coder, ok := decoders[c.encoding]; ok {
		return coder(src)
	}
	return nil, fmt.Errorf("unknown encoding: %s", c.encoding)
}

// AvailableEncodings returns the available encoding id's
func AvailableEncodings() []string {

	res := []string{}

	for key := range encoders {
		res = append(res, key)
	}

	sort.Strings(res)
	return res
}

// RecodeInput processes input `data` according to encodings, used by cmd/coder
func RecodeInput(encodings []string, data []byte, decode bool) ([]byte, error) {

	var err error

	for _, enc := range encodings {

		coder := NewCoder(enc)

		if decode {
			data, err = coder.Decode(string(data))
		} else {
			data, err = coder.Encode(data)
		}
		if err != nil {
			break
		}
	}
	return data, err
}

func encodeASCII85(src []byte) ([]byte, error) {
	buf := make([]byte, ascii85.MaxEncodedLen(len(src)))
	n := ascii85.Encode(buf, src)
	buf = buf[0:n]
	return buf, nil
}

func decodeASCII85(s string) ([]byte, error) {
	dst := make([]byte, 4*len(s))
	ndst, _, err := ascii85.Decode(dst, []byte(s), true)
	return dst[0:ndst], err
}

func encodeBase32(src []byte) ([]byte, error) {
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(src)))
	base32.StdEncoding.Encode(dst, src)
	return dst, nil
}

func decodeBase32(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func encodeBase36(src []byte) ([]byte, error) {
	return base36.EncodeBytesAsBytes(src), nil
}

func decodeBase36(s string) ([]byte, error) {
	return base36.DecodeToBytes(s), nil
}

func encodeBase58(src []byte) ([]byte, error) {
	return []byte(b58.Encode(src)), nil
}

func decodeBase58(s string) ([]byte, error) {
	return b58.Decode(s), nil
}

func encodeBase64(src []byte) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst, nil
}

func decodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func encodeBase91(src []byte) ([]byte, error) {
	return []byte(base91.Encode(string(src))), nil
}

func decodeBase91(s string) ([]byte, error) {
	return []byte(base91.Decode(s)), nil
}

func encodeBinary(src []byte) ([]byte, error) {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%08b", b) + separator
	}

	return []byte(strings.TrimRight(res, separator)), nil
}

func decodeBinary(s string) ([]byte, error) {

	if len(s) == 0 {
		return []byte{}, nil
	}

	parts := strings.Split(s, separator)
	res := make([]byte, len(parts))

	for i, part := range parts {
		b, _ := strconv.ParseInt(part, 2, 8)
		res[i] = byte(b)
	}
	return res, nil
}

func encodeBubbleBabble(src []byte) ([]byte, error) {
	return []byte(bubblebabble.EncodeToString(src)), nil
}

func decodeBubbleBabble(s string) ([]byte, error) {
	return bubblebabble.DecodeString(s)
}

func encodeDecimal(src []byte) ([]byte, error) {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%d", b) + separator
	}

	return []byte(strings.TrimRight(res, separator)), nil
}

func decodeDecimal(s string) ([]byte, error) {

	if len(s) == 0 {
		return []byte{}, nil
	}

	parts := strings.Split(s, separator)
	res := make([]byte, len(parts))

	for i, part := range parts {
		b, _ := strconv.ParseInt(part, 10, 8)
		res[i] = byte(b)
	}
	return res, nil
}

func encodeHex(src []byte) ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst, nil
}

func encodeHexUpper(src []byte) ([]byte, error) {
	return []byte(strings.ToUpper(hex.EncodeToString(src))), nil
}

func decodeHex(s string) ([]byte, error) {

	s = stripSpaces(s)
	res, err := hex.DecodeString(s)
	return res, err
}

func encodeOctal(src []byte) ([]byte, error) {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%#o", b) + separator
	}

	return []byte(strings.TrimRight(res, separator)), nil
}

func decodeOctal(s string) ([]byte, error) {

	if len(s) == 0 {
		return []byte{}, nil
	}

	parts := strings.Split(s, separator)
	res := make([]byte, len(parts))

	for i, part := range parts {
		b, _ := strconv.ParseInt(part, 8, 8)
		res[i] = byte(b)
	}
	return res, nil
}

func encodeUU(src []byte) ([]byte, error) {
	res := uu.EncodeLine(src)
	return res, nil
}

func decodeUU(s string) ([]byte, error) {
	return uu.DecodeLine([]byte(s))
}

func encodeZ85(src []byte) ([]byte, error) {
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
	return b85, err
}

func decodeZ85(s string) ([]byte, error) {

	dst := make([]byte, z85.DecodedLen(len(s)))
	n, err := z85.Decode(dst, []byte(s))

	// strip padding
	for ; n > 0; n-- {
		if dst[n-1] != 0 {
			break
		}
	}
	return dst[0:n], err
}

// defaults to "hex" if encoding is unspecified
func resolveEncodingAliases(s string) string {

	s = strings.ToLower(s)
	if s == "" {
		return "hex"
	}
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
	if s == "base16" || s == "hexadecimal" {
		return "hex"
	}
	if s == "oct" {
		return "octal"
	}
	return s
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
