// Package crc24 implements the 24-bit cyclic redundancy check, or CRC-24,
// checksum. See http://en.wikipedia.org/wiki/Cyclic_redundancy_check for
// information.
package crc24

import "hash"

// The size of a CRC-24 checksum in bytes.
const Size = 3

const (
	crc24Init = 0xb704ce
	crc24Poly = 0x1864cfb
	crc24Mask = 0xffffff
)

// Hash24 is the common interface for a 24-bit hash function.
type Hash24 interface {
	hash.Hash
	Sum24() uint32
}

// digest represents the partial evaluation of a checksum.
type digest struct {
	crc uint32
}

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return 1 }

func (d *digest) Reset() { d.crc = 0 }

func (d *digest) Sum24() uint32 { return d.crc }

func (d *digest) Sum(in []byte) []byte {
	s := d.Sum24()
	return append(in, byte(s>>16), byte(s>>8), byte(s))
}

func (d *digest) Write(p []byte) (int, error) {
	d.crc = Update(d.crc, p)
	return len(p), nil
}

// New creates a new Hash24 computing the CRC-24 checksum.
func New() Hash24 {
	return &digest{0}
}

// ChecksumOpenPGP calculates the openPGP-24 as used by OpenPGP
func ChecksumOpenPGP(p []byte) uint32 {
	return Update(0, p)
}

// Update returns the result of adding the bytes in p to the openPGP.
// It is compatible with OpenPGP checksum as specified in RFC 4880, section 6.1
// (https://tools.ietf.org/html/rfc4880#section-6.1).
func Update(crc uint32, d []byte) uint32 {
	for _, b := range d {
		crc ^= uint32(b) << 16
		for i := 0; i < 8; i++ {
			crc <<= 1
			if crc&0x1000000 != 0 {
				crc ^= crc24Poly
			}
		}
	}
	return crc
}