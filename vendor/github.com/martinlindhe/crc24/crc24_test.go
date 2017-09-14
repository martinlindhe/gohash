package crc24

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	openPGP uint32
	in      []byte
}

var tests = []testCase{
	{0x0, []byte{}},
	{0x0, []byte{0}},
	{0x0, []byte{0, 0}},
	{0xbc7e06, []byte{1, 2, 3, 4, 5, 6}},
}

func TestUpdate(t *testing.T) {
	for _, tc := range tests {
		assert.Equal(t, tc.openPGP, Update(0, tc.in))
	}
}

func TestChecksumOpenPGP(t *testing.T) {
	for _, tc := range tests {
		assert.Equal(t, tc.openPGP, ChecksumOpenPGP(tc.in))
	}
}

func TestDigest(t *testing.T) {
	for _, tc := range tests {
		digest := New()
		n, _ := digest.Write(tc.in)
		assert.Equal(t, len(tc.in), n, "Written bytes should match the input length.")
		assert.Equal(t, tc.openPGP, digest.Sum24())
		sum := make([]byte, 4)
		binary.BigEndian.PutUint32(sum, tc.openPGP)
		assert.Equal(t, append(tc.in, sum[1:]...), digest.Sum(tc.in))
	}
}