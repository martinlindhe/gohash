// Written in 2011-2012 by Dmitry Chestnykh.
//
// To the extent possible under law, the author have dedicated all copyright
// and related and neighboring rights to this software to the public domain
// worldwide. This software is distributed without any warranty.
// http://creativecommons.org/publicdomain/zero/1.0/

// Package blake512 implements BLAKE-512 and BLAKE-384 hash functions (SHA-3
// candidate).
package blake512

import "hash"

// The block size of the hash algorithm in bytes.
const BlockSize = 128

// The size of BLAKE-512 hash in bytes.
const Size = 64

// The size of BLAKE-384 hash in bytes.
const Size384 = 48

type digest struct {
	hashSize int             // hash output size in bits (384 or 512)
	h        [8]uint64       // current chain value
	s        [4]uint64       // salt (zero by default)
	t        uint64          // message bits counter
	nullt    bool            // special case for finalization: skip counter
	x        [BlockSize]byte // buffer for data not yet compressed
	nx       int             // number of bytes in buffer
}

var (
	// Initialization values.
	iv512 = [8]uint64{
		0x6A09E667F3BCC908, 0xBB67AE8584CAA73B,
		0x3C6EF372FE94F82B, 0xA54FF53A5F1D36F1,
		0x510E527FADE682D1, 0x9B05688C2B3E6C1F,
		0x1F83D9ABFB41BD6B, 0x5BE0CD19137E2179}

	iv384 = [8]uint64{
		0xCBBB9D5DC1059ED8, 0x629A292A367CD507,
		0x9159015A3070DD17, 0x152FECD8F70E5939,
		0x67332667FFC00B31, 0x8EB44A8768581511,
		0xDB0C2E0D64F98FA7, 0x47B5481DBEFA4FA4}
)

// Reset resets the state of digest. It leaves salt intact.
func (d *digest) Reset() {
	if d.hashSize == 384 {
		d.h = iv384
	} else {
		d.h = iv512
	}
	d.t = 0
	d.nx = 0
	d.nullt = false
}

func (d *digest) Size() int { return d.hashSize >> 3 }

func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(p []byte) (nn int, err error) {
	nn = len(p)
	if d.nx > 0 {
		n := len(p)
		if n > BlockSize-d.nx {
			n = BlockSize - d.nx
		}
		d.nx += copy(d.x[d.nx:], p)
		if d.nx == BlockSize {
			block(d, d.x[:])
			d.nx = 0
		}
		p = p[n:]
	}
	if len(p) >= BlockSize {
		n := len(p) &^ (BlockSize - 1)
		block(d, p[:n])
		p = p[n:]
	}
	if len(p) > 0 {
		d.nx = copy(d.x[:], p)
	}
	return
}

// Sum returns the calculated checksum.
func (d0 *digest) Sum(in []byte) []byte {
	// Make a copy of d0 so that caller can keep writing and summing.
	d := *d0

	nx := uint64(d.nx)
	l := d.t + nx<<3
	len := make([]byte, 16)
	// len[0 .. 7] = 0, because our counter has only 64 bits.
	len[8] = byte(l >> 56)
	len[9] = byte(l >> 48)
	len[10] = byte(l >> 40)
	len[11] = byte(l >> 32)
	len[12] = byte(l >> 24)
	len[13] = byte(l >> 16)
	len[14] = byte(l >> 8)
	len[15] = byte(l)

	if nx == 111 {
		// One padding byte.
		d.t -= 8
		if d.hashSize == 384 {
			d.Write([]byte{0x80})
		} else {
			d.Write([]byte{0x81})
		}
	} else {
		pad := [129]byte{0x80}
		if nx < 111 {
			// Enough space to fill the block.
			if nx == 0 {
				d.nullt = true
			}
			d.t -= 888 - nx<<3
			d.Write(pad[0 : 111-nx])
		} else {
			// Need 2 compressions.
			d.t -= 1024 - nx<<3
			d.Write(pad[0 : 128-nx])
			d.t -= 888
			d.Write(pad[1:112])
			d.nullt = true
		}
		if d.hashSize == 384 {
			d.Write([]byte{0x00})
		} else {
			d.Write([]byte{0x01})
		}
		d.t -= 8
	}
	d.t -= 128
	d.Write(len)

	out := make([]byte, d.Size())
	j := 0
	for _, s := range d.h[:d.hashSize>>6] {
		out[j+0] = byte(s >> 56)
		out[j+1] = byte(s >> 48)
		out[j+2] = byte(s >> 40)
		out[j+3] = byte(s >> 32)
		out[j+4] = byte(s >> 24)
		out[j+5] = byte(s >> 16)
		out[j+6] = byte(s >> 8)
		out[j+7] = byte(s >> 0)
		j += 8
	}
	return append(in, out...)
}

func (d *digest) setSalt(s []byte) {
	if len(s) != 32 {
		panic("salt length must be 32 bytes")
	}
	j := 0
	for i := 0; i < 4; i++ {
		d.s[i] = uint64(s[j])<<56 | uint64(s[j+1])<<48 | uint64(s[j+2])<<40 |
			uint64(s[j+3])<<32 | uint64(s[j+4])<<24 | uint64(s[j+5])<<16 |
			uint64(s[j+6])<<8 | uint64(s[j+7])
		j += 8
	}
}

// New returns a new hash.Hash computing the BLAKE-512 checksum.
func New() hash.Hash {
	return &digest{
		hashSize: 512,
		h:        iv512,
	}
}

// NewSalt is like New but initializes salt with the given 32-byte slice.
func NewSalt(salt []byte) hash.Hash {
	d := &digest{
		hashSize: 512,
		h:        iv512,
	}
	d.setSalt(salt)
	return d
}

// New384 returns a new hash.Hash computing the BLAKE-384 checksum.
func New384() hash.Hash {
	return &digest{
		hashSize: 384,
		h:        iv384,
	}
}

// New384Salt is like New384 but initializes salt with the given 32-byte slice.
func New384Salt(salt []byte) hash.Hash {
	d := &digest{
		hashSize: 384,
		h:        iv384,
	}
	d.setSalt(salt)
	return d
}
