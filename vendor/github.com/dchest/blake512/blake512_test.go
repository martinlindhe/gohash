// Written in 2011-2012 by Dmitry Chestnykh.
//
// To the extent possible under law, the author have dedicated all copyright
// and related and neighboring rights to this software to the public domain
// worldwide. This software is distributed without any warranty.
// http://creativecommons.org/publicdomain/zero/1.0/

package blake512

import (
	"bytes"
	"fmt"
	"hash"
	"testing"
)

func Test512C(t *testing.T) {
	// Test as in C program.
	var hashes = [][]byte{
		{
			0x97, 0x96, 0x15, 0x87, 0xf6, 0xd9, 0x70, 0xfa, 0xba, 0x6d, 0x24, 0x78, 0x04, 0x5d, 0xe6, 0xd1,
			0xfa, 0xbd, 0x09, 0xb6, 0x1a, 0xe5, 0x09, 0x32, 0x05, 0x4d, 0x52, 0xbc, 0x29, 0xd3, 0x1b, 0xe4,
			0xff, 0x91, 0x02, 0xb9, 0xf6, 0x9e, 0x2b, 0xbd, 0xb8, 0x3b, 0xe1, 0x3d, 0x4b, 0x9c, 0x06, 0x09,
			0x1e, 0x5f, 0xa0, 0xb4, 0x8b, 0xd0, 0x81, 0xb6, 0x34, 0x05, 0x8b, 0xe0, 0xec, 0x49, 0xbe, 0xb3,
		},
		{
			0x31, 0x37, 0x17, 0xd6, 0x08, 0xe9, 0xcf, 0x75, 0x8d, 0xcb, 0x1e, 0xb0, 0xf0, 0xc3, 0xcf, 0x9f,
			0xC1, 0x50, 0xb2, 0xd5, 0x00, 0xfb, 0x33, 0xf5, 0x1c, 0x52, 0xaf, 0xc9, 0x9d, 0x35, 0x8a, 0x2f,
			0x13, 0x74, 0xb8, 0xa3, 0x8b, 0xba, 0x79, 0x74, 0xe7, 0xf6, 0xef, 0x79, 0xca, 0xb1, 0x6f, 0x22,
			0xCE, 0x1e, 0x64, 0x9d, 0x6e, 0x01, 0xad, 0x95, 0x89, 0xc2, 0x13, 0x04, 0x5d, 0x54, 0x5d, 0xde,
		},
	}
	data := make([]byte, 144)

	h := New()
	h.Write(data[:1])
	sum := h.Sum(nil)
	if !bytes.Equal(hashes[0], sum) {
		t.Errorf("0: expected %X, got %X", hashes[0], sum)
	}

	// Try to continue hashing.
	h.Write(data[1:])
	sum = h.Sum(nil)
	if !bytes.Equal(hashes[1], sum) {
		t.Errorf("1(1): expected %X, got %X", hashes[1], sum)
	}

	// Try with reset.
	h.Reset()
	h.Write(data)
	sum = h.Sum(nil)
	if !bytes.Equal(hashes[1], sum) {
		t.Errorf("1(2): expected %X, got %X", hashes[1], sum)
	}
}

type blakeVector struct {
	out, in string
}

var vectors512 = []blakeVector{
	{"1f7e26f63b6ad25a0896fd978fd050a1766391d2fd0471a77afb975e5034b7ad2d9ccf8dfb47abbbe656e1b82fbc634ba42ce186e8dc5e1ce09a885d41f43451",
		"The quick brown fox jumps over the lazy dog"},
	{"7bf805d0d8de36802b882e65d0515aa7682a2be97a9d9ec1399f4be2eff7de07684d7099124c8ac81c1c7c200d24ba68c6222e75062e04feb0e9dd589aa6e3b7",
		"BLAKE"},
	{"a8cfbbd73726062df0c6864dda65defe58ef0cc52a5625090fa17601e1eecd1b628e94f396ae402a00acc9eab77b4d4c2e852aaaa25a636d80af3fc7913ef5b8",
		""},
	{"19bb3a448f4eef6f0b9374817e96c7c848d96f20c5a3e4b808173d97aede52cb396506ac20e174a1d53d9e51e443e7447855f2c9e8c6e4247fa8e4f54cda5897",
		"'BLAKE wins SHA-3! Hooray!!!' (I have time machine)"},
	{"8cd8a7bf2953dd236371a07a3c9e70325abd76922dcb434c68532760e536cf2a955fe8c40d90cb38506fcde30b47da8ee8835064e091427d854ce1dfad972634",
		"Go"},
	{"465d047d9695f258a47af7b94a03d903cb60ae1286f263aac8628774ee90828bea31fb7fe1d3385af364080a317115c8df8596c3c608d8de77b95bff702a3984",
		"HELP! I'm trapped in hash!"},
	{"68376fe303ee09c3a220ee330bccc9fa9fba6dc41741507f195f5457ffa75864076f71bc07e94620123ec24f70458c2ba3dd1fa31a7fefc036d430c962c0969b",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. Praesent et diam eget libero egestas mattis sit amet vitae augue. Nam tincidunt congue enim, ut porta lorem lacinia consectetur. Donec ut libero sed arcu vehicula ultricies a non tortor. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean ut gravida lorem. Ut turpis felis, pulvinar a semper sed, adipiscing id dolor. Pellentesque auctor nisi id magna consequat sagittis. Curabitur dapibus enim sit amet elit pharetra tincidunt feugiat nisl imperdiet. Ut convallis libero in urna ultrices accumsan. Donec sed odio eros. Donec viverra mi quis quam pulvinar at malesuada arcu rhoncus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In rutrum accumsan ultricies. Mauris vitae nisi at sem facilisis semper ac in est.`,
	},
	{"c805573523a7f386732329b6c001e6fe1e1d1842b8152d8f205b86078e571afbaaf4c560cf084fe297e05aac14ae4ded7fdfa2db461fb05d3add28de3f2293c3", // test with one padding byte
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. Praesent et diam eget libero egestas mat`,
	},
}

var vectors384 = []blakeVector{
	{"67c9e8ef665d11b5b57a1d99c96adffb3034d8768c0827d1c6e60b54871e8673651767a2c6c43d0ba2a9bb2500227406",
		"The quick brown fox jumps over the lazy dog"},
	{"f28742f7243990875d07e6afcff962edabdf7e9d19ddea6eae31d094c7fa6d9b00c8213a02ddf1e2d9894f3162345d85",
		"BLAKE"},
	{"c6cbd89c926ab525c242e6621f2f5fa73aa4afe3d9e24aed727faaadd6af38b620bdb623dd2b4788b1c8086984af8706",
		""},
	{"685c83708eb104ad2be3e9f109f5bbb9054fcf38a5ac5ef7e10fbd2ce7853983002dad66cd07c420795733bf97d6fddf",
		"'BLAKE wins SHA-3! Hooray!!!' (I have time machine)"},
	{"9fc155189e04b0325604324bbf11438cb20e05bb634f190ef3e234c6ef0140ec5563303125685aff5690ca1c5a9dbdd5",
		"Go"},
	{"e32f545290b8b5eb2f092cf565854e304e91b5b4b10d9534ed722bcf1da9617f7132a47e37001e63ba833cce5b98ea7d",
		"HELP! I'm trapped in hash!"},
	{"018a6daf616874e8a146906295893599d3f65ad9da9f870ae11c3884dcf3c81fe86f24725ce0853f98f0fb335fdc10ce",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. Praesent et diam eget libero egestas mattis sit amet vitae augue. Nam tincidunt congue enim, ut porta lorem lacinia consectetur. Donec ut libero sed arcu vehicula ultricies a non tortor. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean ut gravida lorem. Ut turpis felis, pulvinar a semper sed, adipiscing id dolor. Pellentesque auctor nisi id magna consequat sagittis. Curabitur dapibus enim sit amet elit pharetra tincidunt feugiat nisl imperdiet. Ut convallis libero in urna ultrices accumsan. Donec sed odio eros. Donec viverra mi quis quam pulvinar at malesuada arcu rhoncus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In rutrum accumsan ultricies. Mauris vitae nisi at sem facilisis semper ac in est.`,
	},
	{"4dc8aa947b13db364e2eb837ae7f561cb301a443e62c92e5e2c890145e83859b906435d30ae46ee6207c84812c2b4eee", // test with one padding byte
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. Praesent et diam eget libero egestas mat`,
	},
}

func testVectors(t *testing.T, hashfunc func() hash.Hash, vectors []blakeVector) {
	for i, v := range vectors {
		h := hashfunc()
		h.Write([]byte(v.in))
		res := fmt.Sprintf("%x", h.Sum(nil))
		if res != v.out {
			t.Errorf("%d: expected %q, got %q", i, v.out, res)
		}
	}
}

func Test512(t *testing.T) {
	testVectors(t, New, vectors512)
}

func Test384(t *testing.T) {
	testVectors(t, New384, vectors384)
}

func TestTwoWrites(t *testing.T) {
	b := make([]byte, 65)
	for i := range b {
		b[i] = byte(i)
	}
	h1 := New()
	h1.Write(b[:1])
	h1.Write(b[1:])
	sum1 := h1.Sum(nil)

	h2 := New()
	h2.Write(b)
	sum2 := h2.Sum(nil)

	if !bytes.Equal(sum1, sum2) {
		t.Errorf("Result of two writes differs from a single write with the same bytes")
	}
}

var bench = New()
var buf = make([]byte, 8<<10)

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

func BenchmarkFull64(b *testing.B) {
	b.SetBytes(64)
	tmp := make([]byte, 32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:64])
		bench.Sum(tmp[0:0])
	}
}
