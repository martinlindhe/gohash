// Package bubblebabble implements Bubble Babble encoding and decoding, as specified
// by http://wiki.yak.net/589.
package bubblebabble

import "strconv"

// The table of Babble vowels.
var vow = []byte("aeiouy")

// The table of Babble consonants.
var con = []byte("bcdfghklmnprstvzx")

// updateChecksum calculates a new Babble checksum value based on the next two
// bytes of input data.
func updateChecksum(c, data1, data2 byte) byte {

	return byte((int(c)*5 + (int(data1)*7 + int(data2))) % 36)
}

// EncodedLen returns the number of bytes an encoded n bytes will take.
func EncodedLen(n int) int {

	nTuples := n / 2
	partialTuple := 3
	terminators := 2
	hyphens := nTuples
	return 5*nTuples + hyphens + partialTuple + terminators
}

// MaxDecodedLen returns the maximum number of bytes a decoding of a Babble
// string of length n will take. There may be a difference of one byte in the
// result length for the same input length depending on the content.
func MaxDecodedLen(n int) int {

	if n == 5 {
		// Only the partial tuple present.
		return 1
	}
	nTuples := (n + 1) / 6
	return nTuples * 2
}

// Encode encodes src into EncodedLen(len(src)) bytes of dst as Bubble Babble
// code.
func Encode(dst, src []byte) int {

	dst[0] = 'x'
	c := byte(1)
	numIter := len(src)/2 + 1

	for i := 0; i < numIter; i++ {
		if i+1 < numIter || len(src)%2 != 0 {
			d1 := src[i*2]

			dst[i*6+1] = vow[(((d1>>6)&3)+c)%6]
			dst[i*6+2] = con[(d1>>2)&15]
			dst[i*6+3] = vow[((d1&3)+c/6)%6]

			if i+1 < numIter {
				d2 := src[i*2+1]
				// Haven't written the last part yet.
				dst[i*6+4] = con[(d2>>4)&15]
				dst[i*6+5] = '-'
				dst[i*6+6] = con[(d2&15)%36]
				c = updateChecksum(c, d1, d2)
			}
		} else {
			// Last part for even data length.
			dst[i*6+1] = vow[c%6]
			dst[i*6+2] = con[16]
			dst[i*6+3] = vow[c/6]
		}
	}
	dst[(len(src)/2)*6+4] = 'x'
	return EncodedLen(len(src))
}

// EncodeToString returns the Bubble Babble encoding of src.
func EncodeToString(src []byte) string {

	dst := make([]byte, EncodedLen(len(src)))
	Encode(dst, src)
	return string(dst)
}

type CorruptInputError int64

func (e CorruptInputError) Error() string {
	return "illegal bubblebabble data at input byte " + strconv.FormatInt(int64(e), 10)
}

// Decode decodes a Babble string into the corresponding byte array. Returns
// the number of bytes decoded, and an error if the string isn't a Babble string.
func Decode(dst, src []byte) (int, error) {

	n, err := decode(dst, src)

	dst = dst[0:n] // XXX shrink dont seem to affect the resulting dst...

	return n, err
}

func decode(dst, src []byte) (int, error) {

	nTuples := len(src) / 6
	c := byte(1)

	n := 0
	var err error

	// Babble strings must be made of one or more hyphen-separated groups of five characters.
	switch {
	case len(src) == 5:
		// One group, ok
	case len(src) > 5 && len(src)%6 == 5:
		// More than one groups, ok.
	default:
		// Bad string length
		return n, CorruptInputError(0)
	}

	// Babble strings must start and end with 'x'.
	if src[0] != 'x' {
		return n, CorruptInputError(0)
	}
	if src[len(src)-1] != 'x' {
		return n, CorruptInputError(len(src) - 1)
	}

	src = src[1:]
	offset := int64(1)

	// Decode the full tuples.
	for i := 0; i < nTuples; i++ {
		t, err := decodeTuple(offset, src, true)
		if err != nil {
			return n, err
		}

		d1, err := decode3WayByte(offset, t[0], t[1], t[2], c)
		if err != nil {
			return n, err
		}

		d2, err := decode2WayByte(offset+int64(4), t[3], t[4])
		if err != nil {
			return n, err
		}
		c = updateChecksum(c, d1, d2)
		dst[i*2] = d1
		dst[i*2+1] = d2

		src = src[6:]
		offset += 6
	}

	// Decode the final partial tuple.
	t, err := decodeTuple(offset, src, false)
	if err != nil {
		return n, err
	}

	if t[1] == 16 {
		// No last byte, final tuple is just checksum data.
		n = nTuples * 2
		if t[0] != c%6 {
			return n, CorruptInputError(offset)
		}
		if t[2] != c/6 {
			return n, CorruptInputError(offset + 2)
		}
	} else {
		// Partial tuple contains one last byte of data, decode it.
		n = nTuples*2 + 1
		d, err := decode3WayByte(offset, t[0], t[1], t[2], c)
		if err != nil {
			return n, err
		}
		dst[nTuples*2] = d
	}

	return n, nil
}

// DecodeString decodes a babble string, returning the resulting byte array.
func DecodeString(src string) (result []byte, err error) {

	result = make([]byte, MaxDecodedLen(len(src)))
	n, err := Decode(result, []byte(src))
	if err != nil {
		return
	}
	result = result[0:n]
	return
}

// devowel converts Babble vowels into the corresponding data values.
func devowel(char byte) (idx byte, ok bool) {

	for i, c := range vow {
		if char == c {
			return byte(i), true
		}
	}
	return 0, false
}

// deconsonant converts Babble consonants into the corresponding data values.
func deconsonant(char byte) (idx byte, ok bool) {

	for i, c := range con {
		if char == c {
			return byte(i), true
		}
	}
	return 0, false
}

// hyphen returns an error if the parameter character is not '-'. It has the
// same function signature as devowel and deconsonant so that it's func value
// can be used in the same type context as theirs.
func hyphen(char byte) (dummy byte, ok bool) { return 0, char == '-' }

// decodeTuple converts a full Bubble Babble string tuple or a data-carrying
// partial tuple into the corresponding byte tuple.
func decodeTuple(offset int64, src []byte, decodeFullTuple bool) (result [5]byte, err error) {

	lut := [](func(byte) (byte, bool)){devowel, deconsonant, devowel, deconsonant, hyphen, deconsonant}
	idx := []int{0, 1, 2, 3, -1, 4}
	for i := 0; i < 6; i++ {
		val, ok := lut[i](src[i])
		if !ok {
			err = CorruptInputError(offset + int64(i))
		}
		if idx[i] >= 0 {
			result[idx[i]] = val
		}
		if i == 2 && !decodeFullTuple {
			return
		}
	}
	return
}

// decode3WayByte decodes a byte that has been encoded into three Babble
// characters. Returns an error if the data is invalid or if it fails a
// checksum check.
func decode3WayByte(offset int64, a1, a2, a3 byte, c byte) (result byte, err error) {

	high2 := (int(a1) - int(c%6) + 6) % 6
	if high2 >= 4 {
		err = CorruptInputError(offset)
		return
	}
	if a2 > 16 {
		err = CorruptInputError(offset + 1)
		return
	}
	mid4 := int(a2)
	low2 := (int(a3) - int(c/6%6) + 6) % 6
	if low2 >= 4 {
		err = CorruptInputError(offset + 2)
		return
	}
	result = byte(high2<<6) | byte(mid4<<2) | byte(low2)
	return
}

// decode2WayByte decodes a byte that has been encoded into two Babble
// characters. This type of encoding uses all the available bits to represent
// data, so a checksum value is not used.
func decode2WayByte(offset int64, a1, a2 byte) (result byte, err error) {

	if a1 > 16 {
		err = CorruptInputError(offset)
		return
	}
	if a2 > 16 {
		err = CorruptInputError(offset + 1)
		return
	}

	result = (a1 << 4) | a2
	return
}
