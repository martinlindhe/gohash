// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package md2

import (
	"fmt"
	"io"
	"testing"
)

type md2Test struct {
	out string
	in  string
}

var golden = []md2Test{
	{"8350e5a3e24c153df2275c9f80692773", ""},
	{"32ec01ec4a6dac72c0ab96fb34c0b5d1", "a"},
	{"da853b0d3f88d99b30283a69e6ded6bb", "abc"},
	{"ab4f496bfb2a530b219ff33031fe06b0", "message digest"},
	{"4e8ddff3650292ab5a4108c3aa47940b", "abcdefghijklmnopqrstuvwxyz"},
	{"da33def2a42df13975352846c30338cd", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"},
}

func TestGolden(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i]
		c := New()
		for j := 0; j < 3; j++ {
			if j < 2 {
				io.WriteString(c, g.in)
			} else {
				//fmt.Printf("Testing the first half\n")
				io.WriteString(c, g.in[0:len(g.in)/2])
				c.Sum(nil)
				//fmt.Printf("Testing the second half\n")
				io.WriteString(c, g.in[len(g.in)/2:])
			}
			s := fmt.Sprintf("%x", c.Sum(nil))
			if s != g.out {
				t.Fatalf("md2[%d](%s) = %s want %s", j, g.in, s, g.out)
			} else {
				fmt.Printf("md2[%d](%s) = %s want %s - Passed\n", j, g.in, s, g.out)
			}
			c.Reset()
		}
	}
}

func ExampleNew() {
	h := New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x", h.Sum(nil))
}

var bench = New()
var buf = makeBuf()

func makeBuf() []byte {
	b := make([]byte, 8<<10)
	for i := range b {
		b[i] = byte(i)
	}
	return b
}

func BenchmarkHash1K(b *testing.B) {
	b.SetBytes(1024)
	for i := 0; i < b.N; i++ {
		bench.Write(buf[:1024])
	}
}

func BenchmarkHash8K(b *testing.B) {
	b.SetBytes(int64(len(buf)))
	for i := 0; i < b.N; i++ {
		bench.Write(buf)
	}
}
