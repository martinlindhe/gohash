package uu

import (
	"gopkg.in/check.v1"
	"testing"
)

// Test hooks up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

//TstEnc is the test suit
type TstEnc struct{}

var _ = check.Suite(&TstEnc{})

//TestFindLevel tests hello world
func (s *TstEnc) TestConv3Bytes1(c *check.C) {
	lRes := ToUUPack([3]byte{0x14, 0x0F, 0xA8})
	c.Check(lRes, check.Equals, [4]byte{0x25, 0x20, 0x5E, 0x48})
}

//TestFindLevel tests hello world
func (s *TstEnc) TestConv3Bytes2(c *check.C) {
	lRes := ToUUPack([3]byte{0x17, 0x00, 0x00})
	c.Check(lRes, check.Equals, [4]byte{0x25, 0x50, 0x20, 0x20})
}

//TestFindLevel tests hello world
func (s *TstEnc) TestEncodeLineInvalidLen(c *check.C) {
	lIn := make([]byte, MaxBytesPerLine+1)
	lRes := EncodeLine(lIn)
	c.Check(lRes, check.IsNil)
}

//TestFindLevel tests hello world
func (s *TstEnc) TestEncodeLine3Bytes(c *check.C) {
	lIn := []byte{0x14, 0x0F, 0xA8}
	lRes := string(EncodeLine(lIn))
	c.Check(lRes, check.Equals, "#% ^H\r\n")
}

//TestFindLevel tests hello world
func (s *TstEnc) TestEncodeLine4Bytes(c *check.C) {
	lIn := []byte{0x14, 0x0F, 0xA8, 0x17}
	lRes := string(EncodeLine(lIn))
	c.Check(lRes, check.Equals, "$% ^H%P  \r\n")
}

//TestFindLevel tests hello world
func (s *TstEnc) TestEncodeLine1(c *check.C) {
	lIn := []byte("http://www.wikipedia.org\r\n")
	lRes := string(EncodeLine(lIn[:len(lIn)]))
	c.Check(lRes, check.Equals, "::'1T<#HO+W=W=RYW:6MI<&5D:6$N;W)G#0H \r\n")
}
