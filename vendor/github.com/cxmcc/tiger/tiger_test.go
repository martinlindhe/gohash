package tiger

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"unsafe"
)

type Test struct {
	out string
	in  string
}

var golden = []Test{
	{"3293ac630c13f0245f92bbb1766e16167a4e58492dde73f3", ""},
	{"77befbef2e7ef8ab2ec8f93bf587a7fc613e247f5f247809", "a"},
	{"2aab1484e8c158f2bfb8c5ff41b57a525129131c957b5f93", "abc"},
	{"d981f8cb78201a950dcf3048751e441c517fca1aa55a29f6", "message digest"},
	{"1714a472eee57d30040412bfcc55032a0b11602ff37beee9", "abcdefghijklmnopqrstuvwxyz"},
	{"0f7bf9a19b9c58f2b7610df7e84f0ac3a71c631e7b53f78e", "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq"},
	{"8dcea680a17583ee502ba38a3c368651890ffbccdc49a8cc", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"},
	{"1c14795529fd9f207a958f84c52f11e887fa0cabdfd91bfd", "12345678901234567890123456789012345678901234567890123456789012345678901234567890"},
	{"cdf0990c5c6b6b0bddd63a75ed20e2d448bf44e15fde0df4", strings.Repeat("A", 1024)},
	{"89292aee0f82842abc080c57b3aadd9ca84d66bf0cae77aa", strings.Repeat("A", 1025)},
}

func TestGolden(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i]
		c := New()
		buf := make([]byte, len(g.in)+4)
		for j := 0; j < 7; j++ {
			if j < 2 {
				io.WriteString(c, g.in)
			} else if j == 2 {
				io.WriteString(c, g.in[0:len(g.in)/2])
				c.Sum(nil)
				io.WriteString(c, g.in[len(g.in)/2:])
			} else if j > 2 {
				// test unaligned write
				buf = buf[1:]
				copy(buf, g.in)
				c.Write(buf[:len(g.in)])
			}
			s := fmt.Sprintf("%x", c.Sum(nil))
			if s != g.out {
				t.Fatalf("tiger[%d](%s) = %s want %s", j, g.in, s, g.out)
			}
			c.Reset()
		}
	}
}

type WriteTest struct {
	out int
	in  string
}

var writeTestVectors = []WriteTest{
	{0, ""},
	{1, "A"},
	{2, "AA"},
	{10, strings.Repeat("A", 10)},
	{1024, strings.Repeat("A", 1024)},
	{1025, strings.Repeat("A", 1025)},
	{0, ""},
}

func TestWriteReturnsCorrectSize(t *testing.T) {
	c := New()
	for i := 0; i < len(writeTestVectors); i++ {
		v := writeTestVectors[i]
		b := []byte(v.in)
		length, err := c.Write(b[:len(v.in)])
		if length != v.out {
			t.Fatalf("Write() = %d want %d", length, v.out)
		}
		if err != nil {
			t.Fatalf("Write(%s) failed.", v.in)
		}
	}
}

func ExampleNew() {
	h := New()
	io.WriteString(h, "It's the eye of the tiger, it's the thrill of the fight")
	io.WriteString(h, "Rising up to the challenge of our rival!")
	fmt.Printf("%x", h.Sum(nil))
	// Output: a7bbad36cc17918e399ae8ee893e4595e4d24e1639fe822c
}

func ExampleNew2() {
	h := New2()
	io.WriteString(h, "It's the eye of the tiger, it's the thrill of the fight")
	io.WriteString(h, "Rising up to the challenge of our rival!")
	fmt.Printf("%x", h.Sum(nil))
	// Output: c86695c2a639506682de2c12c2d23b61a12db78ea1ee1001
}

var bench = New()
var buf = make([]byte, 8192+1)
var sum = make([]byte, bench.Size())

func benchmarkSize(b *testing.B, size int, unaligned bool) {
	b.SetBytes(int64(size))
	buf := buf
	if unaligned {
		if uintptr(unsafe.Pointer(&buf[0]))&(unsafe.Alignof(uint32(0))-1) == 0 {
			buf = buf[1:]
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:size])
		bench.Sum(sum[:0])
	}
}

func BenchmarkHash8Bytes(b *testing.B) {
	benchmarkSize(b, 8, false)
}

func BenchmarkHash1K(b *testing.B) {
	benchmarkSize(b, 1024, false)
}

func BenchmarkHash8K(b *testing.B) {
	benchmarkSize(b, 8192, false)
}

func BenchmarkHash8BytesUnaligned(b *testing.B) {
	benchmarkSize(b, 8, true)
}

func BenchmarkHash1KUnaligned(b *testing.B) {
	benchmarkSize(b, 1024, true)
}

func BenchmarkHash8KUnaligned(b *testing.B) {
	benchmarkSize(b, 8192, true)
}
