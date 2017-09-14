package hexutils

import (
	"gopkg.in/check.v1"
	"testing"
)

// Test hooks up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

//MySuite is the test suit
type MySuite struct{}

var _ = check.Suite(&MySuite{})

//TestToUInt8 tests ToUInt8
func (s *MySuite) TestToUInt8(c *check.C) {
	lVal, lErr := ToUInt8([2]byte{byte('0'), byte('1')})
	c.Check(lErr, check.IsNil)
	c.Check(lVal, check.Equals, uint8(0x01))
}

//TestToUInt8Invalid tests invalid input to ToUInt8
func (s *MySuite) TestToUInt8Invalid(c *check.C) {
	lVal, lErr := ToUInt8([2]byte{byte('S'), byte('1')})
	c.Check(lErr, check.Equals, ErrNotHex)
	c.Check(lVal, check.Equals, uint8(0x0))
}

//TestToUInt16 tests ToUInt16
func (s *MySuite) TestToUInt16(c *check.C) {
	lVal, lErr := ToUInt16([4]byte{byte('0'), byte('1'), byte('0'), byte('0')})
	c.Check(lErr, check.IsNil)
	c.Check(lVal, check.Equals, uint16(0x0100))
}

//TestToUInt16Invalid tests invalid inputs to ToUInt16
func (s *MySuite) TestToUInt16Invalid(c *check.C) {
	lVal, lErr := ToUInt16([4]byte{byte('S'), byte('1'), byte('0'), byte('0')})
	c.Check(lErr, check.Equals, ErrNotHex)
	c.Check(lVal, check.Equals, uint16(0x0))
}
