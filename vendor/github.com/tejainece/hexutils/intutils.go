package hexutils

import (
//"fmt"
)

//MakeUInt16BE makes uint16 from byte array
func MakeUInt16BE(aBytes [2]byte) uint16 {
	lRet := uint16(0)

	lRet = uint16(aBytes[0])
	lRet = lRet << 8
	lRet |= uint16(aBytes[1])

	return lRet
}

//MakeUInt16LE makes uint16 from byte array
func MakeUInt16LE(aBytes [2]byte) uint16 {
	lRet := uint16(0)

	lRet = uint16(aBytes[1])
	lRet = lRet << 8
	lRet |= uint16(aBytes[0])

	return lRet
}

//MakeUInt32FromUInt16 makes uint32 from two uint16
func MakeUInt32FromUInt16(aHi, aLo uint16) uint32 {
	lRet := uint32(0)

	lRet = uint32(aHi)
	lRet = lRet << 16
	lRet |= uint32(aLo)

	return lRet
}

func MakeUInt32FromUInt8(aB3, aB2, aB1, aB0 uint8) uint32 {
	lRet := uint32(0)

	lRet |= uint32(aB3)
	lRet <<= 8

	lRet |= uint32(aB2)
	lRet <<= 8

	lRet |= uint32(aB1)
	lRet <<= 8

	lRet |= uint32(aB0)

	return lRet
}

func MakeUInt32FromByte(aB3, aB2, aB1, aB0 byte) uint32 {
	lRet := uint32(0)

	lRet |= uint32(aB3)
	lRet <<= 8

	lRet |= uint32(aB2)
	lRet <<= 8

	lRet |= uint32(aB1)
	lRet <<= 8

	lRet |= uint32(aB0)

	return lRet
}

func GetByte0FromUInt32(aData uint32) byte {
	return byte(aData & 0xFF)
}

func GetByte1FromUInt32(aData uint32) byte {
	return byte((aData >> 8) & 0xFF)
}

func GetByte2FromUInt32(aData uint32) byte {
	return byte((aData >> 16) & 0xFF)
}

func GetByte3FromUInt32(aData uint32) byte {
	return byte((aData >> 24) & 0xFF)
}

func GetUInt80FromUInt32(aData uint32) uint8 {
	return uint8(aData & 0xFF)
}

func GetUInt81FromUInt32(aData uint32) uint8 {
	return uint8((aData >> 8) & 0xFF)
}

func GetUInt82FromUInt32(aData uint32) uint8 {
	return uint8((aData >> 16) & 0xFF)
}

func GetUInt83FromUInt32(aData uint32) uint8 {
	return uint8((aData >> 24) & 0xFF)
}

func GetUInt160FromUInt32(aData uint32) uint16 {
	return uint16(aData & 0xFFFF)
}

func GetUInt161FromUInt32(aData uint32) uint16 {
	return uint16((aData >> 16) & 0xFFFF)
}
