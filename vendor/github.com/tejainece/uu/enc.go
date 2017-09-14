package uu

const (
	//MaxBytesPerLine is maximum bytes allowed in per line
	MaxBytesPerLine = 45

	//UUCharStart is integer representation of first allowed Uuencode char
	//Equvalent to 'space' in ASCII table
	UUCharStart = 0x20

	//UUCharEnd is the integer representation of last allowed Uuencode char
	//Equivalent to '`' in ASCII table
	UUCharEnd = 0x60

	//UUCharPseudoZero represents pseudo zero
	UUCharPseudoZero = 0x60

	//PCharNewline is integer representation of newline character
	PCharNewline = 0x0A

	//PCharCR is integer representation of carriage return character
	PCharCR = 0x0D
)

//ToUUPack converts 3 bytes into UUPack
func ToUUPack(aBytes [3]byte) [4]byte {
	lRet := [4]byte{}

	lInt32 := uint32(0)
	for cIdx := 0; cIdx < len(aBytes); cIdx++ {
		lInt32 <<= 8
		lInt32 |= uint32(aBytes[cIdx])
	}

	for cIdx := 0; cIdx < len(lRet); cIdx++ {
		lRet[len(lRet)-1-cIdx] = UUCharStart + byte(lInt32&0x3F)
		lInt32 >>= 6
	}

	return lRet
}

//EncodeLine encodes the given bytes into Uuencode line
func EncodeLine(aBytes []byte) []byte {
	if len(aBytes) > MaxBytesPerLine {
		return nil
	}

	if len(aBytes) == 0 {
		return []byte{UUCharPseudoZero, PCharCR, PCharNewline}
	}

	var lRet []byte

	lRet = append(lRet, byte(UUCharStart+len(aBytes)))

	for cIdx := 0; cIdx < (len(aBytes) / 3); cIdx++ {
		bIn := [3]byte{}
		copy(bIn[:], aBytes[cIdx*3:((cIdx+1)*3)])
		bOut := ToUUPack(bIn)
		lRet = append(lRet, bOut[0], bOut[1], bOut[2], bOut[3])
	}

	{
		bLeft := len(aBytes) % 3
		if bLeft > 0 {
			bIn := [3]byte{}
			copy(bIn[0:bLeft], aBytes[len(aBytes)-bLeft:])
			bOut := ToUUPack(bIn)
			lRet = append(lRet, bOut[0], bOut[1], bOut[2], bOut[3])
		}
	}

	lRet = append(lRet, PCharCR, PCharNewline)

	return lRet
}
