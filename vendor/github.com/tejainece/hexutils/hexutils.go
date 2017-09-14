package hexutils

import (
	"encoding/hex"
)

//Err contains information about the failure
type Err string

//Error returns printable string of the failure
func (meErr Err) Error() string {
	return string(meErr)
}

const (
	//ErrNotHex is error when the hex digit array contains non-hex digit
	ErrNotHex = Err("Encountered non-hex digit!")
)

//ToUInt8 converts hex digits into uint8. Digits are assumed to be
//in big endian format.
func ToUInt8(aBytes [2]byte) (uint8, error) {
	lDec := [1]byte{}

	lNum, lErr := hex.Decode(lDec[:], aBytes[:])

	if (lNum != len(lDec)) || (lErr != nil) {
		return 0, ErrNotHex
	}

	return uint8(lDec[0]), nil
}

//ToUInt8 converts hex bytes into uint8. Digits are assumed to be
//in big endian format within bytes.
func ToUInt16(aBytes [4]byte) (uint16, error) {
	lDec := [2]byte{}

	lNum, lErr := hex.Decode(lDec[:], aBytes[:])

	if (lNum != len(lDec)) || (lErr != nil) {
		return 0, ErrNotHex
	}

	return MakeUInt16BE(lDec), nil
}
