package gohash

import (
	"testing"

	"github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var (
	f                 = fuzz.New()
	iterationsPerAlgo = 10 // increase to fuzz properly, slows down tests
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
			fox:   "01010100 01101000 01100101 00100000 01110001 01110101 01101001 01100011 01101011 00100000 01100010 01110010 01101111 01110111 01101110 00100000 01100110 01101111 01111000 00100000 01101010 01110101 01101101 01110000 01110011 00100000 01101111 01110110 01100101 01110010 00100000 01110100 01101000 01100101 00100000 01101100 01100001 01111010 01111001 00100000 01100100 01101111 01100111",
			blank: ""},
		"decimal": {
			fox:   "84 104 101 32 113 117 105 99 107 32 98 114 111 119 110 32 102 111 120 32 106 117 109 112 115 32 111 118 101 114 32 116 104 101 32 108 97 122 121 32 100 111 103",
			blank: ""},
		"hex": {
			fox:   "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
			blank: ""},
		"hexup": {
			fox:   "54686520717569636B2062726F776E20666F78206A756D7073206F76657220746865206C617A7920646F67",
			blank: ""},
		"octal": {
			fox:   "0124 0150 0145 040 0161 0165 0151 0143 0153 040 0142 0162 0157 0167 0156 040 0146 0157 0170 040 0152 0165 0155 0160 0163 040 0157 0166 0145 0162 040 0164 0150 0145 040 0154 0141 0172 0171 040 0144 0157 0147",
			blank: ""},
		"z85": {
			fox:   "ra]?=ADL#9yAN8bz*c7ww]z]pyisxjB0byAwPw]nxK@r5vs0hwwn=8X",
			blank: ""},
	}
)

func TestCalcExpectedEncodings(t *testing.T) {

	for algo, forms := range expectedEncodings {
		for form, hash := range forms {
			coder := NewCoder(algo)
			res, err := coder.Encode([]byte(form))
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
				res, err := coder.Decode(coded)
				assert.Equal(t, nil, err, algo)
				assert.Equal(t, []byte(clear), res, algo)
			}
		}
	}
}

func TestFuzzEncoders(t *testing.T) {

	for algo := range expectedEncodings {
		for i := 0; i < iterationsPerAlgo; i++ {
			var rnd []byte
			f.Fuzz(&rnd)
			coder := NewCoder(algo)
			coder.Encode(rnd)
		}
	}
}

func TestFuzzDecoders(t *testing.T) {

	for algo := range expectedEncodings {
		for i := 0; i < iterationsPerAlgo; i++ {
			rnd := ""
			f.Fuzz(&rnd)
			coder := NewCoder(algo)
			coder.Decode(rnd)
		}
	}
}

func TestEncodeZ85(t *testing.T) {

	res, err := encodeZ85([]byte{0x86, 0x4F, 0xD2, 0x6F, 0xB5, 0x59, 0xF7, 0x5B})
	assert.Equal(t, nil, err)
	assert.Equal(t, "HelloWorld", string(res))
}

func TestDecodeHexWithSpaces(t *testing.T) {

	res, err := decodeHex("48 4f 2a")
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{0x48, 0x4f, 0x2a}, res)
}
