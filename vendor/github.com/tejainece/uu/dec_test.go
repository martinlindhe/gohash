package uu

import (
	"gopkg.in/check.v1"
)

//TstDec is the test suit
type TstDec struct{}

var _ = check.Suite(&TstDec{})

//TestDecodePack1 tests DecodePack
func (s *TstDec) TestDecodePack1(c *check.C) {
	lRes, lSts := DecodePack([4]byte{0x25, 0x20, 0x5E, 0x48})
	c.Check(lSts, check.Equals, true)
	c.Check(lRes, check.Equals, [3]byte{0x14, 0x0F, 0xA8})
}

//TestFindLevel tests hello world
func (s *TstDec) TestDecodePack2(c *check.C) {
	lRes, lSts := DecodePack([4]byte{0x25, 0x50, 0x20, 0x20})
	c.Check(lSts, check.Equals, true)
	c.Check(lRes, check.Equals, [3]byte{0x17, 0x00, 0x00})
}

//TestFindLevel tests hello world
func (s *TstDec) TestDecodePack3(c *check.C) {
	lRes, lSts := DecodePack([4]byte{0x25, 0x50, 0x60, 0x60})
	c.Check(lSts, check.Equals, true)
	c.Check(lRes, check.Equals, [3]byte{0x17, 0x00, 0x00})
}

//TestFindLevel tests hello world
func (s *TstDec) TestDecodePackErr1(c *check.C) {
	lRes, lSts := DecodePack([4]byte{0x0A, 0x50, 0x60, 0x60})
	c.Check(lSts, check.Equals, false)
	c.Check(lRes, check.Equals, [3]byte{0x0, 0x0, 0x0})
}

//TestFindLevel tests hello world
func (s *TstDec) TestDecodeLineInvalidLineLen(c *check.C) {
	lIn := make([]byte, MaxBytesPerEncLine+1)
	lRes, lErr := DecodeLine(lIn)
	c.Check(lErr, check.Equals, ErrInvalidLineLen)
	c.Check(lRes, check.IsNil)
}

//TestFindLevel tests hello world
func (s *TstDec) TestDecodeLineInvalidDataLen1(c *check.C) {
	lIn := []byte{UUCharStart + 2, UUCharStart, UUCharStart, UUCharStart, PCharCR, PCharNewline}
	lRes, lErr := DecodeLine(lIn)
	c.Check(lErr, check.Equals, ErrInvalidDataLen)
	c.Check(lRes, check.IsNil)
}

//TestFindLevel tests hello world
func (s *TstDec) TestDecodeLineInvalidData1(c *check.C) {
	lIn := []byte{0x61, PCharCR, PCharNewline}
	lRes, lErr := DecodeLine(lIn)
	c.Check(lErr, check.Equals, ErrInvalidData)
	c.Check(lRes, check.IsNil)
}

//TestFindLevel tests hello world
func (s *TstDec) TestDecodeLineInvalidData2(c *check.C) {
	lIn := []byte{UUCharStart + 2, UUCharStart, UUCharStart, UUCharStart - 2, UUCharStart, PCharCR, PCharNewline}
	lRes, lErr := DecodeLine(lIn)
	c.Check(lErr, check.Equals, ErrInvalidData)
	c.Check(lRes, check.IsNil)
}

//TestFindLevel tests hello world
func (s *TstDec) TestDecodeLine1(c *check.C) {
	lIn := []byte("::'1T<#HO+W=W=RYW:6MI<&5D:6$N;W)G#0H \r\n")
	lResBytes, _ := DecodeLine(lIn[:len(lIn)])
	lRes := string(lResBytes)
	c.Check(lRes, check.Equals, "http://www.wikipedia.org\r\n")
}
