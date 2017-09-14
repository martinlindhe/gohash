package z85

import (
	"bytes"
	"testing"
)

func TestEncodingAndDecoding(t *testing.T) {
	tests := []struct {
		src []byte
		z85 string
		err error
	}{
		{nil, "", nil},
		{[]byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B}, "HelloWorld", nil},
		{[]byte{0x8E, 0x0B, 0xDD, 0x69, 0x76, 0x28, 0xB9, 0x1D,
			0x8F, 0x24, 0x55, 0x87, 0xEE, 0x95, 0xC5, 0xB0,
			0x4D, 0x48, 0x96, 0x3F, 0x79, 0x25, 0x98, 0x77,
			0xB4, 0x9C, 0xD9, 0x06, 0x3A, 0xEA, 0xD3, 0xB7,
		}, "JTKVSB%%)wK0E.X)V>+}o?pNmC{O&4W4b!Ni{Lh6", nil},

		// Alignment error
		{[]byte{1, 2, 3}, "abcdefghi", ErrLength},

		// Invalid data
		{nil, `     `, InvalidByteError(' ')}, // ' ' < minDigit
		{nil, `~~~~~`, InvalidByteError('~')}, // '~' > maxDigit
		{nil, `"""""`, InvalidByteError('"')}, // '"' is in range, but not part of the z85 alphabet
	}

	for _, test := range tests {

		// Encode
		if _, ok := test.err.(InvalidByteError); !ok {
			dst := make([]byte, EncodedLen(len(test.src)))
			n, err := Encode(dst, test.src)
			if err != test.err {
				t.Fatalf("got %q, want %q", err, test.err)
			}
			if err == nil {
				if !bytes.Equal([]byte(test.z85), dst) {
					t.Fatalf("got %q, want %q", dst, test.z85)
				}
				if n != len(test.z85) {
					t.Fatalf("got %d, want %d", n, len(test.z85))
				}
			}
		}

		// Decode
		dst := make([]byte, DecodedLen(len(test.z85)))
		n, err := Decode(dst, []byte(test.z85))
		if err != test.err {
			t.Fatalf("got %q, want %q", err, test.err)
		}
		if err == nil {
			if !bytes.Equal(test.src, dst) {
				t.Fatalf("got %d, want %d", dst, test.src)
			}
			if n != len(dst) {
				t.Fatalf("got %d, want %d", n, len(dst))
			}
		}
	}
}
