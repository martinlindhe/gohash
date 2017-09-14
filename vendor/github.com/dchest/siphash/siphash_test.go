// Written in 2012 by Dmitry Chestnykh.
//
// To the extent possible under law, the author have dedicated all copyright
// and related and neighboring rights to this software to the public domain
// worldwide. This software is distributed without any warranty.
// http://creativecommons.org/publicdomain/zero/1.0/

package siphash

import (
	"bytes"
	"encoding/binary"
	"testing"
)

var zeroKey = make([]byte, 16)

var golden = []struct {
	k []byte
	m []byte
	r uint64
}{
	{
		[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
		[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e},
		0xa129ca6149be45e5,
	},
	{
		zeroKey,
		[]byte("Hello world"),
		0xc9e8a3021f3822d9,
	},
	{
		zeroKey,
		[]byte{}, // zero-length message
		0x1e924b9d737700d7,
	},
	{
		zeroKey,
		[]byte("12345678123"),
		0xf95d77ccdb0649f,
	},
	{
		zeroKey,
		make([]byte, 8),
		0xe849e8bb6ffe2567,
	},
	{
		zeroKey,
		make([]byte, 1535),
		0xe74d1c0ab64b2afa,
	},
}

// Test vectors from reference implementation.
//
// SipHash-2-4 output with
// k = 00 01 02 ...
// and
// in = (empty string)
// in = 00 (1 byte)
// in = 00 01 (2 bytes)
// in = 00 01 02 (3 bytes)
// ...
// in = 00 01 02 ... 3e (63 bytes)
var goldenRef = [][]byte{
	{0x31, 0x0e, 0x0e, 0xdd, 0x47, 0xdb, 0x6f, 0x72},
	{0xfd, 0x67, 0xdc, 0x93, 0xc5, 0x39, 0xf8, 0x74},
	{0x5a, 0x4f, 0xa9, 0xd9, 0x09, 0x80, 0x6c, 0x0d},
	{0x2d, 0x7e, 0xfb, 0xd7, 0x96, 0x66, 0x67, 0x85},
	{0xb7, 0x87, 0x71, 0x27, 0xe0, 0x94, 0x27, 0xcf},
	{0x8d, 0xa6, 0x99, 0xcd, 0x64, 0x55, 0x76, 0x18},
	{0xce, 0xe3, 0xfe, 0x58, 0x6e, 0x46, 0xc9, 0xcb},
	{0x37, 0xd1, 0x01, 0x8b, 0xf5, 0x00, 0x02, 0xab},
	{0x62, 0x24, 0x93, 0x9a, 0x79, 0xf5, 0xf5, 0x93},
	{0xb0, 0xe4, 0xa9, 0x0b, 0xdf, 0x82, 0x00, 0x9e},
	{0xf3, 0xb9, 0xdd, 0x94, 0xc5, 0xbb, 0x5d, 0x7a},
	{0xa7, 0xad, 0x6b, 0x22, 0x46, 0x2f, 0xb3, 0xf4},
	{0xfb, 0xe5, 0x0e, 0x86, 0xbc, 0x8f, 0x1e, 0x75},
	{0x90, 0x3d, 0x84, 0xc0, 0x27, 0x56, 0xea, 0x14},
	{0xee, 0xf2, 0x7a, 0x8e, 0x90, 0xca, 0x23, 0xf7},
	{0xe5, 0x45, 0xbe, 0x49, 0x61, 0xca, 0x29, 0xa1},
	{0xdb, 0x9b, 0xc2, 0x57, 0x7f, 0xcc, 0x2a, 0x3f},
	{0x94, 0x47, 0xbe, 0x2c, 0xf5, 0xe9, 0x9a, 0x69},
	{0x9c, 0xd3, 0x8d, 0x96, 0xf0, 0xb3, 0xc1, 0x4b},
	{0xbd, 0x61, 0x79, 0xa7, 0x1d, 0xc9, 0x6d, 0xbb},
	{0x98, 0xee, 0xa2, 0x1a, 0xf2, 0x5c, 0xd6, 0xbe},
	{0xc7, 0x67, 0x3b, 0x2e, 0xb0, 0xcb, 0xf2, 0xd0},
	{0x88, 0x3e, 0xa3, 0xe3, 0x95, 0x67, 0x53, 0x93},
	{0xc8, 0xce, 0x5c, 0xcd, 0x8c, 0x03, 0x0c, 0xa8},
	{0x94, 0xaf, 0x49, 0xf6, 0xc6, 0x50, 0xad, 0xb8},
	{0xea, 0xb8, 0x85, 0x8a, 0xde, 0x92, 0xe1, 0xbc},
	{0xf3, 0x15, 0xbb, 0x5b, 0xb8, 0x35, 0xd8, 0x17},
	{0xad, 0xcf, 0x6b, 0x07, 0x63, 0x61, 0x2e, 0x2f},
	{0xa5, 0xc9, 0x1d, 0xa7, 0xac, 0xaa, 0x4d, 0xde},
	{0x71, 0x65, 0x95, 0x87, 0x66, 0x50, 0xa2, 0xa6},
	{0x28, 0xef, 0x49, 0x5c, 0x53, 0xa3, 0x87, 0xad},
	{0x42, 0xc3, 0x41, 0xd8, 0xfa, 0x92, 0xd8, 0x32},
	{0xce, 0x7c, 0xf2, 0x72, 0x2f, 0x51, 0x27, 0x71},
	{0xe3, 0x78, 0x59, 0xf9, 0x46, 0x23, 0xf3, 0xa7},
	{0x38, 0x12, 0x05, 0xbb, 0x1a, 0xb0, 0xe0, 0x12},
	{0xae, 0x97, 0xa1, 0x0f, 0xd4, 0x34, 0xe0, 0x15},
	{0xb4, 0xa3, 0x15, 0x08, 0xbe, 0xff, 0x4d, 0x31},
	{0x81, 0x39, 0x62, 0x29, 0xf0, 0x90, 0x79, 0x02},
	{0x4d, 0x0c, 0xf4, 0x9e, 0xe5, 0xd4, 0xdc, 0xca},
	{0x5c, 0x73, 0x33, 0x6a, 0x76, 0xd8, 0xbf, 0x9a},
	{0xd0, 0xa7, 0x04, 0x53, 0x6b, 0xa9, 0x3e, 0x0e},
	{0x92, 0x59, 0x58, 0xfc, 0xd6, 0x42, 0x0c, 0xad},
	{0xa9, 0x15, 0xc2, 0x9b, 0xc8, 0x06, 0x73, 0x18},
	{0x95, 0x2b, 0x79, 0xf3, 0xbc, 0x0a, 0xa6, 0xd4},
	{0xf2, 0x1d, 0xf2, 0xe4, 0x1d, 0x45, 0x35, 0xf9},
	{0x87, 0x57, 0x75, 0x19, 0x04, 0x8f, 0x53, 0xa9},
	{0x10, 0xa5, 0x6c, 0xf5, 0xdf, 0xcd, 0x9a, 0xdb},
	{0xeb, 0x75, 0x09, 0x5c, 0xcd, 0x98, 0x6c, 0xd0},
	{0x51, 0xa9, 0xcb, 0x9e, 0xcb, 0xa3, 0x12, 0xe6},
	{0x96, 0xaf, 0xad, 0xfc, 0x2c, 0xe6, 0x66, 0xc7},
	{0x72, 0xfe, 0x52, 0x97, 0x5a, 0x43, 0x64, 0xee},
	{0x5a, 0x16, 0x45, 0xb2, 0x76, 0xd5, 0x92, 0xa1},
	{0xb2, 0x74, 0xcb, 0x8e, 0xbf, 0x87, 0x87, 0x0a},
	{0x6f, 0x9b, 0xb4, 0x20, 0x3d, 0xe7, 0xb3, 0x81},
	{0xea, 0xec, 0xb2, 0xa3, 0x0b, 0x22, 0xa8, 0x7f},
	{0x99, 0x24, 0xa4, 0x3c, 0xc1, 0x31, 0x57, 0x24},
	{0xbd, 0x83, 0x8d, 0x3a, 0xaf, 0xbf, 0x8d, 0xb7},
	{0x0b, 0x1a, 0x2a, 0x32, 0x65, 0xd5, 0x1a, 0xea},
	{0x13, 0x50, 0x79, 0xa3, 0x23, 0x1c, 0xe6, 0x60},
	{0x93, 0x2b, 0x28, 0x46, 0xe4, 0xd7, 0x06, 0x66},
	{0xe1, 0x91, 0x5f, 0x5c, 0xb1, 0xec, 0xa4, 0x6c},
	{0xf3, 0x25, 0x96, 0x5c, 0xa1, 0x6d, 0x62, 0x9f},
	{0x57, 0x5f, 0xf2, 0x8e, 0x60, 0x38, 0x1b, 0xe5},
	{0x72, 0x45, 0x06, 0xeb, 0x4c, 0x32, 0x8a, 0x95},
}

func TestSum64(t *testing.T) {
	for i, v := range golden {
		h := New(v.k)
		h.Write(v.m)
		if sum := h.Sum64(); sum != v.r {
			t.Errorf(`%d: expected "%x", got "%x"`, i, v.r, sum)
		}
	}
}

func TestSum(t *testing.T) {
	var r [8]byte
	for i, v := range golden {
		binary.LittleEndian.PutUint64(r[:], v.r)
		h := New(v.k)
		h.Write(v.m)
		if sum := h.Sum(nil); !bytes.Equal(sum, r[:]) {
			t.Errorf(`%d: expected "%x", got "%x"`, i, r, sum)
		}
	}

	var k [16]byte
	var in [64]byte
	for i := range k {
		k[i] = byte(i)
	}

	for i := 0; i < 64; i++ {
		in[i] = byte(i)
		h := New(k[:])
		h.Write(in[:i])
		if sum := h.Sum(nil); !bytes.Equal(sum, goldenRef[i]) {
			t.Errorf(`%d: expected "%x", got "%x"`, i, goldenRef[i], sum)
		}
	}
}

func TestHash(t *testing.T) {
	var k0, k1 uint64
	for i, v := range golden {
		k0 = binary.LittleEndian.Uint64(v.k[0:8])
		k1 = binary.LittleEndian.Uint64(v.k[8:16])
		if sum := Hash(k0, k1, v.m); sum != v.r {
			t.Errorf(`%d: expected "%x", got "%x"`, i, v.r, sum)
		}
	}

	var k [16]byte
	var in [64]byte
	for i := range k {
		k[i] = byte(i)
	}
	k0 = binary.LittleEndian.Uint64(k[0:8])
	k1 = binary.LittleEndian.Uint64(k[8:16])

	for i := 0; i < 64; i++ {
		in[i] = byte(i)
		ref := binary.LittleEndian.Uint64(goldenRef[i])
		if sum := Hash(k0, k1, in[:i]); sum != ref {
			t.Errorf(`%d: expected "%x", got "%x"`, i, ref, sum)
		}
	}
}

var key = zeroKey
var key0, key1 uint64
var bench = New(key)
var buf = make([]byte, 8<<10)

func BenchmarkHash8(b *testing.B) {
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		Hash(key0, key1, buf[:8])
	}
}

func BenchmarkHash16(b *testing.B) {
	b.SetBytes(16)
	for i := 0; i < b.N; i++ {
		Hash(key0, key1, buf[:16])
	}
}

func BenchmarkHash40(b *testing.B) {
	b.SetBytes(40)
	for i := 0; i < b.N; i++ {
		Hash(key0, key1, buf[:40])
	}
}

func BenchmarkHash64(b *testing.B) {
	b.SetBytes(64)
	for i := 0; i < b.N; i++ {
		Hash(key0, key1, buf[:64])
	}
}

func BenchmarkHash128(b *testing.B) {
	b.SetBytes(128)
	for i := 0; i < b.N; i++ {
		Hash(key0, key1, buf[:128])
	}
}

func BenchmarkHash1K(b *testing.B) {
	b.SetBytes(1024)
	for i := 0; i < b.N; i++ {
		Hash(key0, key1, buf[:1024])
	}
}

func BenchmarkHash8K(b *testing.B) {
	b.SetBytes(int64(len(buf)))
	for i := 0; i < b.N; i++ {
		Hash(key0, key1, buf)
	}
}

func BenchmarkFull8(b *testing.B) {
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:8])
		bench.Sum64()
	}
}

func BenchmarkFull16(b *testing.B) {
	b.SetBytes(16)
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:16])
		bench.Sum64()
	}
}

func BenchmarkFull40(b *testing.B) {
	b.SetBytes(24)
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:16])
		bench.Sum64()
	}
}

func BenchmarkFull64(b *testing.B) {
	b.SetBytes(64)
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:64])
		bench.Sum64()
	}
}

func BenchmarkFull128(b *testing.B) {
	b.SetBytes(128)
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:64])
		bench.Sum64()
	}
}

func BenchmarkFull1K(b *testing.B) {
	b.SetBytes(1024)
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:1024])
		bench.Sum64()
	}
}

func BenchmarkFull8K(b *testing.B) {
	b.SetBytes(int64(len(buf)))
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf)
		bench.Sum64()
	}
}
