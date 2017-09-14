package uu

import (
	"github.com/tejainece/hexutils"
	"math"
)

const (
	//MaxBytesPerEncLine is maximum number of bytes allowed in an UULine
	MaxBytesPerEncLine = 61
)

//DecodeUUByte decodes given UUByte into normal byte
func DecodeUUByte(aData byte) (byte, bool) {
	if aData < UUCharStart || aData > UUCharEnd {
		return 0, false
	}

	lRet := aData - UUCharStart

	if lRet == 0x40 {
		lRet = 0
	}

	return lRet, true
}

//DecodePack decodes given UUPack into byte array
func DecodePack(aBytes [4]byte) ([3]byte, bool) {
	lRet := [3]byte{}

	lInt32 := uint32(0)
	for cIdx := 0; cIdx < len(aBytes); cIdx++ {
		bByte, bErr := DecodeUUByte(aBytes[cIdx])

		if !bErr {
			return lRet, false
		}

		lInt32 <<= 6
		lInt32 |= uint32(bByte)
	}

	lRet[2] = hexutils.GetByte0FromUInt32(lInt32)
	lRet[1] = hexutils.GetByte1FromUInt32(lInt32)
	lRet[0] = hexutils.GetByte2FromUInt32(lInt32)

	return lRet, true
}

//DecodeLine decodes given UULine and returns data
func DecodeLine(aBytes []byte) ([]byte, error) {
	if len(aBytes) > MaxBytesPerEncLine {
		return nil, ErrInvalidLineLen
	}

	lLen, lSts := DecodeUUByte(aBytes[0])
	if !lSts {
		return nil, ErrInvalidData
	}

	lEncLen := int(math.Ceil(float64(lLen)/3)) * 4

	//+3 because of length, termination CR, termination LF
	lExpLineLen := lEncLen + 3

	if len(aBytes) != lExpLineLen {
		return nil, ErrInvalidDataLen
	}

	lRet := make([]byte, 0, lEncLen)

	lBytes := aBytes[1 : len(aBytes)-2]
	for cIdx := 0; cIdx < lEncLen; cIdx += 4 {
		bTempBuf := [4]byte{}
		copy(bTempBuf[:], lBytes[cIdx:cIdx+4])
		bBytes, bErr := DecodePack(bTempBuf)

		if !bErr {
			return nil, ErrInvalidData
		}

		lRet = append(lRet, bBytes[0], bBytes[1], bBytes[2])
	}

	return lRet[:lLen], nil
}
