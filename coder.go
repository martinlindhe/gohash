package gohash

import (
	"bytes"
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/bproctor/base91"
	b58 "github.com/jbenet/go-base58"
	"github.com/martinlindhe/base36"
	"github.com/martinlindhe/bubblebabble"
	"github.com/martinlindhe/uu"
	"github.com/tilinna/z85"
)

// Coder is used to encode and decode various binary-to-text encodings
type Coder struct {
	encoding string
}

var (
	separator = " "
	encoders  = map[string]func(r io.Reader) ([]byte, error){
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
		"reverse":      encodeReverse,
		"rot13":        encodeROT13,
		"rot47":        encodeROT47,
		"uu":           encodeUU,
		"z85":          encodeZ85,
	}

	decoders = map[string]func(r io.Reader) ([]byte, error){
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
		"reverse":      encodeReverse,
		"rot13":        encodeROT13,
		"rot47":        encodeROT47,
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

// Encode encodes `r` into some encoding
func (c *Coder) Encode(r io.Reader) ([]byte, error) {
	if coder, ok := encoders[c.encoding]; ok {
		return coder(r)
	}
	return nil, fmt.Errorf("unknown encoding: %s", c.encoding)
}

// Decode decodes `r` from some encoding
func (c *Coder) Decode(r io.Reader) ([]byte, error) {
	if coder, ok := decoders[c.encoding]; ok {
		return coder(r)
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
func RecodeInput(encodings []string, r io.Reader, decode bool) ([]byte, error) {
	var err error
	data, _ := ioutil.ReadAll(r)

	for _, enc := range encodings {
		coder := NewCoder(enc)
		datan := bytes.NewReader(data)
		if decode {
			data, err = coder.Decode(datan)
		} else {
			data, err = coder.Encode(datan)
		}
		if err != nil {
			break
		}
	}
	return data, err
}

func encodeASCII85(r io.Reader) ([]byte, error) {
	var b bytes.Buffer
	h := ascii85.NewEncoder(&b)
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	err := h.Close()
	return b.Bytes(), err
}

func decodeASCII85(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	// Often, ascii85-encoded data is wrapped in <~ and ~> symbols. Decode() expects these to have been stripped by the caller.
	if len(src) >= 4 && src[0] == '<' && src[1] == '~' && src[len(src)-2] == '~' && src[len(src)-1] == '>' {
		src = src[2 : len(src)-2]
	}
	r = bytes.NewBuffer(src)
	dec := ascii85.NewDecoder(r)
	return ioutil.ReadAll(dec)
}

func encodeBase32(r io.Reader) ([]byte, error) {
	var b bytes.Buffer
	h := base32.NewEncoder(base32.StdEncoding, &b)
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	err := h.Close()
	return b.Bytes(), err
}

func decodeBase32(r io.Reader) ([]byte, error) {
	h := base32.NewDecoder(base32.StdEncoding, r)
	return ioutil.ReadAll(h)
}

func encodeBase36(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	return base36.EncodeBytesAsBytes(src), nil
}

func decodeBase36(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	return base36.DecodeToBytes(string(src)), nil
}

func encodeBase58(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	return []byte(b58.Encode(src)), nil
}

func decodeBase58(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	return b58.Decode(string(src)), nil
}

func encodeBase64(r io.Reader) ([]byte, error) {
	var b bytes.Buffer
	h := base64.NewEncoder(base64.StdEncoding, &b)
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	err := h.Close()
	return b.Bytes(), err
}

func decodeBase64(r io.Reader) ([]byte, error) {
	h := base64.NewDecoder(base64.StdEncoding, r)
	return ioutil.ReadAll(h)
}

func encodeBase91(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	return base91.Encode(src), nil
}

func decodeBase91(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	if len(src) == 0 {
		return []byte{}, nil
	}
	return base91.Decode(src), nil
}

func encodeBinary(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	var res bytes.Buffer
	for n, b := range src {
		fmt.Fprintf(&res, "%08b", b)
		if n < len(src)-1 {
			res.WriteString(separator)
		}
	}
	return res.Bytes(), nil
}

func decodeBinary(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	if len(src) == 0 {
		return []byte{}, nil
	}
	parts := strings.Split(string(src), separator)
	res := make([]byte, len(parts))

	for i, part := range parts {
		b, _ := strconv.ParseUint(part, 2, 8)
		res[i] = byte(b)
	}
	return res, nil
}

func encodeBubbleBabble(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	return []byte(bubblebabble.EncodeToString(src)), nil
}

func decodeBubbleBabble(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	return bubblebabble.DecodeString(string(src))
}

func encodeDecimal(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	var res bytes.Buffer
	for n, b := range src {
		fmt.Fprintf(&res, "%d", b)
		if n < len(src)-1 {
			res.WriteString(separator)
		}
	}
	return res.Bytes(), nil
}

func decodeDecimal(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	if len(src) == 0 {
		return []byte{}, nil
	}
	parts := strings.Split(string(src), separator)
	res := make([]byte, len(parts))

	for i, part := range parts {
		b, _ := strconv.ParseUint(part, 10, 8)
		res[i] = byte(b)
	}
	return res, nil
}

func encodeHex(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst, nil
}

func encodeHexUpper(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	return []byte(strings.ToUpper(hex.EncodeToString(src))), nil
}

func decodeHex(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	s := stripSeparators(string(src))
	res, err := hex.DecodeString(s)
	return res, err
}

func encodeOctal(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	var res bytes.Buffer
	for n, b := range src {
		fmt.Fprintf(&res, "%#o", b)
		if n < len(src)-1 {
			res.WriteString(separator)
		}
	}
	return res.Bytes(), nil
}

func decodeOctal(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

	if len(src) == 0 {
		return []byte{}, nil
	}

	parts := strings.Split(string(src), separator)
	res := make([]byte, len(parts))

	for i, part := range parts {
		b, _ := strconv.ParseUint(part, 8, 8)
		res[i] = byte(b)
	}
	return res, nil
}

// encode/decode bytes in reverse order
func encodeReverse(r io.Reader) ([]byte, error) {
	a, _ := ioutil.ReadAll(r)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a, nil
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

// encode/decode bytes in ROT13
func encodeROT13(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	for i, b := range src {
		src[i] = rot13(b)
	}
	return src, nil
}

// encode/decode bytes in ROT47
func encodeROT47(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	for i, b := range src {
		if b > 32 && b < 80 {
			src[i] += 47
		} else if b > 79 && b < 127 {
			src[i] -= 47
		}
	}
	return src, nil
}

func encodeUU(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	return uu.Encode(src, "file.txt", "644")
}

func decodeUU(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	dec, err := uu.Decode(src)
	return dec.Data, err
}

func encodeZ85(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)

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

func decodeZ85(r io.Reader) ([]byte, error) {
	src, _ := ioutil.ReadAll(r)
	src = []byte(strings.TrimRight(string(src), "\r\n"))
	dst := make([]byte, z85.DecodedLen(len(src)))
	_, err := z85.Decode(dst, src)
	dst = stripZ85Padding(dst)
	return dst, err
}

func stripZ85Padding(b []byte) []byte {
	// trim trailing 0:s for z85, because "The binary frame SHALL have a length that is
	// divisible by 4 with no remainder.", and trailing 0:s will not remain in the decode of
	// such sequence. See https://rfc.zeromq.org/spec:32/Z85/

	n := len(b)
	for ; n > 0; n-- {
		if b[n-1] != 0 {
			break
		}
	}
	return b[0:n]
}

// defaults to "hex" if encoding is unspecified
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
	if s == "base16" || s == "hexadecimal" {
		return "hex"
	}
	if s == "oct" {
		return "octal"
	}
	return s
}

// removes spaces and ; from string
func stripSeparators(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || r == ';' {
			return -1
		}
		return r
	}, str)
}
