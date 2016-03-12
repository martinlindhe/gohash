package gohash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expectedEncodings = map[string]expectedForms{
		"ascii85": expectedForms{
			fox:   "<+ohcEHPu*CER),Dg-(AAoDo:C3=B4F!,CEATAo8BOr<&@=!2AA8c)",
			blank: ""},
		"base32": expectedForms{
			fox:   "KRUGKIDROVUWG2ZAMJZG653OEBTG66BANJ2W24DTEBXXMZLSEB2GQZJANRQXU6JAMRXWO===",
			blank: ""},
		/*"base36": expectedForms{
		fox:   "xx",
		blank: ""},*/
		"base58": expectedForms{
			fox:   "7DdiPPYtxLjCD3wA1po2rvZHTDYjkZYiEtazrfiwJcwnKCizhGFhBGHeRdx",
			blank: ""},
		"base64": expectedForms{
			fox:   "VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZw==",
			blank: ""},
		"bubblebabble": expectedForms{
			fox:   "xihak-minod-besol-hopak-fypad-bumal-daril-lurad-binik-zovad-bepyl-hirol-bysod-barel-konal-domel-gipuk-hamok-somyl-pivad-bonuk-zanox",
			blank: "xexax"},
		"binary": expectedForms{
			fox:   "01010100 01101000 01100101 00100000 01110001 01110101 01101001 01100011 01101011 00100000 01100010 01110010 01101111 01110111 01101110 00100000 01100110 01101111 01111000 00100000 01101010 01110101 01101101 01110000 01110011 00100000 01101111 01110110 01100101 01110010 00100000 01110100 01101000 01100101 00100000 01101100 01100001 01111010 01111001 00100000 01100100 01101111 01100111",
			blank: ""},
		"decimal": expectedForms{
			fox:   "84 104 101 32 113 117 105 99 107 32 98 114 111 119 110 32 102 111 120 32 106 117 109 112 115 32 111 118 101 114 32 116 104 101 32 108 97 122 121 32 100 111 103",
			blank: ""},
		"hex": expectedForms{
			fox:   "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
			blank: ""},
		"hexup": expectedForms{
			fox:   "54686520717569636B2062726F776E20666F78206A756D7073206F76657220746865206C617A7920646F67",
			blank: ""},
		"octal": expectedForms{
			fox:   "0124 0150 0145 040 0161 0165 0151 0143 0153 040 0142 0162 0157 0167 0156 040 0146 0157 0170 040 0152 0165 0155 0160 0163 040 0157 0166 0145 0162 040 0164 0150 0145 040 0154 0141 0172 0171 040 0144 0157 0147",
			blank: ""},
		"z85": expectedForms{
			fox:   "ra]?=ADL#9yAN8bz*c7ww]z]pyisxjB0byAwPw]nxK@r5vs0hwwn=8X",
			blank: ""},
	}
)

func TestCalcExpectedEncodings(t *testing.T) {

	for algo, forms := range expectedEncodings {
		for form, hash := range forms {
			coder := NewCoder(algo)
			res, err := coder.Encode([]byte(form))
			if err != nil {
				t.Fatalf("ERROR algo fail %s: %s", algo, err)
			}
			assert.Equal(t, hash, res, algo)
		}
	}
}
