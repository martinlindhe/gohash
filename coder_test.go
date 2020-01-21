package gohash

import (
	"bytes"
	"strings"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

type expectedForms map[string]string

var (
	blank = ""
	fox   = "The quick brown fox jumps over the lazy dog"

	f                 = fuzz.New()
	iterationsPerAlgo = 100 // increase to fuzz properly, slows down tests
	expectedEncodings = map[string]expectedForms{
		"ascii85": {
			fox:   "<+ohcEHPu*CER),Dg-(AAoDo:C3=B4F!,CEATAo8BOr<&@=!2AA8c)",
			blank: ""},
		"base32": {
			fox:   "KRUGKIDROVUWG2ZAMJZG653OEBTG66BANJ2W24DTEBXXMZLSEB2GQZJANRQXU6JAMRXWO===",
			blank: ""},
		"base36": {
			fox:   "29T3UBYZNHH32O9X3PZVLJP1QA22WUN2QUO35NAEMPN0GTX0LXIYRF3QWWI0LZU165J",
			blank: ""},
		"base58": {
			fox:   "7DdiPPYtxLjCD3wA1po2rvZHTDYjkZYiEtazrfiwJcwnKCizhGFhBGHeRdx",
			blank: ""},
		"base64": {
			fox:   "VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZw==",
			blank: ""},
		"base91": {
			fox:   "nX^Iz?T1s!2t:aRn#o>vf>6C9#`##mlLK#_1:Wzv;RG!,a%q3Lc=Z",
			blank: ""},
		"bubblebabble": {
			fox:   "xihak-minod-besol-hopak-fypad-bumal-daril-lurad-binik-zovad-bepyl-hirol-bysod-barel-konal-domel-gipuk-hamok-somyl-pivad-bonuk-zanox",
			blank: "xexax"},
		"binary": {
			fox: "01010100 01101000 01100101 00100000 01110001 01110101 01101001 01100011 01101011 00100000 01100010" +
				" 01110010 01101111 01110111 01101110 00100000 01100110 01101111 01111000 00100000 01101010 01110101" +
				" 01101101 01110000 01110011 00100000 01101111 01110110 01100101 01110010 00100000 01110100 01101000" +
				" 01100101 00100000 01101100 01100001 01111010 01111001 00100000 01100100 01101111 01100111",
			blank: ""},
		"decimal": {
			fox: "84 104 101 32 113 117 105 99 107 32 98 114 111 119 110 32 102 111 120 32 106 117 109 112 115 32 111" +
				" 118 101 114 32 116 104 101 32 108 97 122 121 32 100 111 103",
			blank: ""},
		"hex": {
			fox:   "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
			blank: ""},
		"hexup": {
			fox:   "54686520717569636B2062726F776E20666F78206A756D7073206F76657220746865206C617A7920646F67",
			blank: ""},
		"octal": {
			fox: "0124 0150 0145 040 0161 0165 0151 0143 0153 040 0142 0162 0157 0167 0156 040 0146 0157 0170 040" +
				" 0152 0165 0155 0160 0163 040 0157 0166 0145 0162 040 0164 0150 0145 040 0154 0141 0172 0171 040 0144 0157 0147",
			blank: ""},
		"reverse": {
			fox:   "god yzal eht revo spmuj xof nworb kciuq ehT",
			blank: ""},
		"rot13": {
			fox:   "Gur dhvpx oebja sbk whzcf bire gur ynml qbt",
			blank: ""},
		"rot47": {
			fox:   "%96 BF:4< 3C@H? 7@I ;F>AD @G6C E96 =2KJ 5@8",
			blank: ""},
		"uu": {
			fox:   "begin 644 file.txt\nK5&AE('%U:6-K(&)R;W=N(&9O>\"!J=6UP<R!O=F5R('1H92!L87IY(&1O9P\n`\nend\n",
			blank: "begin 644 file.txt\n`\nend\n",
		},
		"z85": {
			fox:   "ra]?=ADL#9yAN8bz*c7ww]z]pyisxjB0byAwPw]nxK@r5vs0hwwn=8X",
			blank: ""},
	}
)

func TestEncodingDefines(t *testing.T) {
	for algo := range encoders {
		if _, ok := decoders[algo]; !ok {
			t.Error("algo not defined in decoders map", algo)
		}
		if _, ok := expectedEncodings[algo]; !ok {
			t.Error("algo lacks testcase in expectedEncodings map", algo)
		}
	}
	for algo := range decoders {
		if _, ok := encoders[algo]; !ok {
			t.Error("algo not defined in encoders map", algo)
		}
	}
}

func TestCalcExpectedEncodings(t *testing.T) {
	for algo, forms := range expectedEncodings {
		for form, hash := range forms {
			coder := NewCoder(algo)
			res, err := coder.Encode(strings.NewReader(form))
			assert.Equal(t, nil, err, algo)
			assert.Equal(t, hash, string(res), algo)
		}
	}
}

func TestCalcExpectedDecodings(t *testing.T) {
	for algo := range decoders {
		if forms, ok := expectedEncodings[algo]; ok {
			for clear, coded := range forms {
				coder := NewCoder(algo)
				res, err := coder.Decode(strings.NewReader(coded))
				if err != nil {
					t.Error(algo, coded, err)
				} else {
					assert.Equal(t, []byte(clear), res, algo)
				}
			}
		}
	}
}

func TestFuzzEncoders(t *testing.T) {
	// feed encoders with random data and decoding it, making sure we always get the same result back
	for algo := range expectedEncodings {
		for i := 0; i < iterationsPerAlgo; i++ {
			var rnd []byte
			f.Fuzz(&rnd)
			if algo == "z85" {
				rnd = stripZ85Padding(rnd)
			}
			coder := NewCoder(algo)
			enc, err := coder.Encode(bytes.NewReader(rnd))
			assert.Equal(t, nil, err, algo)

			if err == nil {
				dec, err := coder.Decode(bytes.NewReader(enc))
				assert.Equal(t, nil, err, algo)
				if err == nil {
					assert.Equal(t, string(rnd), string(dec), algo+" decode of "+string(enc))
				}
			}
		}
	}
}

func TestFuzzDecoders(t *testing.T) {
	// feed decoders with random data, looking for crashes
	for algo := range expectedEncodings {
		for i := 0; i < iterationsPerAlgo; i++ {
			rnd := ""
			f.Fuzz(&rnd)
			coder := NewCoder(algo)
			coder.Decode(strings.NewReader(rnd))
		}
	}
}

func TestEncodeZ85(t *testing.T) {
	res, err := encodeZ85(bytes.NewReader([]byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B}))
	assert.Equal(t, nil, err)
	assert.Equal(t, "HelloWorld", string(res))
}

func TestDecodeHexWithSpaces(t *testing.T) {
	res, err := decodeHex(strings.NewReader("48 4f 2a"))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{0x48, 0x4f, 0x2a}, res)
}

func TestDecodeHexWithSeparators(t *testing.T) {
	res, err := decodeHex(strings.NewReader("48;4f;2a"))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{0x48, 0x4f, 0x2a}, res)
}

func TestRecodeInputEncodeSingle(t *testing.T) {
	res, err := RecodeInput([]string{"base64"}, strings.NewReader("hello"), false, false)
	assert.Equal(t, nil, err)
	assert.Equal(t, "aGVsbG8=", string(res))
}

func TestRecodeInputDecodeSingle(t *testing.T) {
	res, err := RecodeInput([]string{"base64"}, strings.NewReader("aGVsbG8="), true, false)
	assert.Equal(t, nil, err)
	assert.Equal(t, "hello", string(res))
}

func TestRecodeInputDecodeChain(t *testing.T) {
	res, err := RecodeInput([]string{"hex", "base64"}, strings.NewReader("614756736247383d"), true, false)
	assert.Equal(t, nil, err)
	assert.Equal(t, "hello", string(res))
}

func TestDecodeOctal(t *testing.T) {
	res, err := decodeOctal(bytes.NewReader([]byte{}))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{}, res)

	res, err = decodeOctal(strings.NewReader("0377"))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{0xff}, res)
}

func TestDecodeBinary(t *testing.T) {
	res, err := decodeBinary(bytes.NewReader([]byte{}))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{}, res)

	res, err = decodeBinary(strings.NewReader("11111111"))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{0xff}, res)
}

func TestDecodeDecimal(t *testing.T) {
	res, err := decodeDecimal(bytes.NewReader([]byte{}))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{}, res)

	res, err = decodeDecimal(strings.NewReader("255"))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{0xff}, res)
}

func TestDecodeASCII85(t *testing.T) {
	s := "<~BOu!rDZ~>"

	res, err := decodeASCII85(strings.NewReader(s))
	assert.Equal(t, nil, err)
	assert.Equal(t, "hello", string(res))
}

func BenchmarkEncodeBinary(b *testing.B) {
	for n := 0; n < b.N; n++ {
		coder := NewCoder("binary")
		coder.Encode(bytes.NewReader([]byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B, 0x48, 0x4F, 0x2A, 0x48, 0x4F, 0x2A}))
	}
}

func BenchmarkEncodeDecimal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		coder := NewCoder("decimal")
		coder.Encode(bytes.NewReader([]byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B, 0x48, 0x4F, 0x2A, 0x48, 0x4F, 0x2A}))
	}
}

func BenchmarkEncodeOctal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		coder := NewCoder("octal")
		coder.Encode(bytes.NewReader([]byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B, 0x48, 0x4F, 0x2A, 0x48, 0x4F, 0x2A}))
	}
}

func BenchmarkEncodeASCII85(b *testing.B) {
	for n := 0; n < b.N; n++ {
		coder := NewCoder("ascii85")
		coder.Encode(strings.NewReader("zafdklsahdfkjlkajsgdfkhjgajshdgfjklagsdfasdfkjlhskdjgfjhagsdfjhgasjdgfkjh"))
	}
}

func BenchmarkEncodeBase32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		coder := NewCoder("base32")
		coder.Encode(strings.NewReader("zafdklsahdfkjlkajsgdfkhjgajshdgfjklagsdfasdfkjlhskdjgfjhagsdfjhgasjdgfkjh"))
	}
}

func BenchmarkEncodeBase64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		coder := NewCoder("base64")
		coder.Encode(strings.NewReader("zafdklsahdfkjlkajsgdfkhjgajshdgfjklagsdfasdfkjlhskdjgfjhagsdfjhgasjdgfkjh"))
	}
}
