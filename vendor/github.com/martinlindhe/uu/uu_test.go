package uu

import (
	"strings"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

var (
	fuzzRounds   = 10000
	encodedLine1 = `K5&AE('%U:6-K(&)R;W=N(&9O>"!J=6UP<R!O=F5R('1H92!L87IY(&1O9P  `
	mode1        = "644"
	file1        = `stuff.txt`
	encoded1     = `begin ` + mode1 + ` ` + file1 + "\n" +
		encodedLine1 + "\n" +
		"`\n" +
		"end\n"
	clear1 = `The quick brown fox jumps over the lazy dog`

	encodedBlock2 = []string{
		`M5&AE('%U:6-K(&)R;W=N(&9O>"!J=6UP<R!O=F5R('1H92!L87IY(&1O9RP@`,
		`K5&AE('%U:6-K(&)R;W=N(&9O>"!J=6UP<R!O=F5R('1H92!L87IY(&1O9P  `,
	}
	clear2 = `The quick brown fox jumps over the lazy dog, The quick brown fox jumps over the lazy dog`

	file3         = `file.dat`
	mode3         = "0744"
	encodedBlock3 = []string{
		"M5&AE(&YA;64@(G5U96YC;V1I;F<B(&ES(&1E<FEV960@9G)O;2 B56YI>\"UT",
		"M;RU5;FEX(&5N8V]D:6YG(BP@:2YE+B!T:&4@:61E82!O9B!U<VEN9R!A('-A",
		`M9F4@96YC;V1I;F<@=&\@=')A;G-F97(@56YI>"!F:6QE<R!F<F]M(&]N92!5`,
		"M;FEX('-Y<W1E;2!T;R!A;F]T:&5R(%5N:7@@<WES=&5M(&)U=\"!W:71H;W5T",
		"M(&=U87)A;G1E92!T:&%T('1H92!I;G1E<G9E;FEN9R!L:6YK<R!W;W5L9\"!A",
		"3;&P@8F4@56YI>\"!S>7-T96US+@``",
	}

	encoded3 = `begin ` + mode3 + ` ` + file3 + "\n" +
		strings.Join(encodedBlock3, "\n") + "\n" +
		"`\n" +
		"end\n"
	clear3 = `The name "uuencoding" is derived from "Unix-to-Unix encoding",` +
		` i.e. the idea of using a safe encoding to transfer Unix files from one` +
		` Unix system to another Unix system but without guarantee that the` +
		` intervening links would all be Unix systems.`
)

func TestEncode(t *testing.T) {
	out, err := Encode([]byte(clear1), file1, mode1)
	assert.Equal(t, nil, err)
	assert.Equal(t, encoded1, string(out))

	out, err = Encode([]byte{}, file1, mode1)
	assert.Equal(t, nil, err)
	assert.Equal(t, "begin 644 stuff.txt\n`\nend\n", string(out))
}

func TestDecode(t *testing.T) {
	out, err := Decode([]byte(encoded1))
	assert.Equal(t, nil, err)
	assert.Equal(t, file1, out.Filename)
	assert.Equal(t, clear1, string(out.Data))
	assert.Equal(t, mode1, out.Mode)

	out, err = Decode([]byte(encoded3))
	assert.Equal(t, nil, err)
	assert.Equal(t, file3, out.Filename)
	assert.Equal(t, clear3, string(out.Data))
	assert.Equal(t, mode3, out.Mode)
}

func TestDecodeLine(t *testing.T) {
	out, err := DecodeLine(encodedLine1)
	assert.Equal(t, nil, err)
	assert.Equal(t, clear1, string(out))
}

func TestDecodeBlock(t *testing.T) {
	out, err := DecodeBlock([]string{encodedLine1})
	assert.Equal(t, nil, err)
	assert.Equal(t, clear1, string(out))

	out, err = DecodeBlock(encodedBlock2)
	assert.Equal(t, nil, err)
	assert.Equal(t, clear2, string(out))

	out, err = DecodeBlock(encodedBlock3)
	assert.Equal(t, nil, err)
	assert.Equal(t, clear3, string(out))
}

func TestFuzzDecode(t *testing.T) {
	f := fuzz.New()
	for i := 0; i < fuzzRounds; i++ {
		rnd := ""
		f.Fuzz(&rnd)
		Decode([]byte(rnd))
	}
}

func TestFuzzDecodeBlock(t *testing.T) {
	f := fuzz.New()
	for i := 0; i < fuzzRounds; i++ {
		rnd := ""
		f.Fuzz(&rnd)
		DecodeBlock([]string{rnd})
	}
}

func TestFuzzDecodeLine(t *testing.T) {
	f := fuzz.New()
	for i := 0; i < fuzzRounds; i++ {
		rnd := ""
		f.Fuzz(&rnd)
		DecodeLine(rnd)
	}
}
