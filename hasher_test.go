package gohash

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexStringToBytes(t *testing.T) {

	assert.Equal(t, []byte("hej"), hexStringToBytes("68656a"))
}

func TestHashMd5(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("md5")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("541c57960bb997942655d14e3b9607f9")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}

func TestHashSha1(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha1")
	hasher.AllowedKeys("tom")
	hasher.ExpectedHash("9af7d87edaba03e23f6dbdaed29101ee1291c8a6")
	hasher.Length(5)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "motom", string(res))
}

func TestHashSha224(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha224")
	hasher.AllowedKeys("bexlopslmv")
	hasher.ExpectedHash("0136053cfe4adcb6b0130742b5abcdd086b79bc367a9e27eeedc4fe5")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "slop", string(res))
}

func TestHashSha256(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha256")
	hasher.AllowedKeys("aeltx")
	hasher.ExpectedHash("a0b04eaca76db465982af821d6a304c3b904a7f13d6a9704d135aa07b3f1f6c2")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "tex", string(res))
}

func TestHashSha384(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha384")
	hasher.AllowedKeys("bexlopslmv")
	hasher.ExpectedHash("35dc3a0822d431161309a0fe460786618c7a6fd3b0d2883e03fc21c851000ce2817bad134c1a72e179ada6fc207b6f15")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "slop", string(res))
}

func TestHashSha512(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha512")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("2b4df6c7b86d49a71c5ad6c1ffc2e85fde618c69400d2a0ccabb8dd12df4ae2584103b6947379742c0bc11a4e81ad4a3a832c11a734bf8ae5f8a8af9317a4c03")
	hasher.Length(4)

	assert.Equal(t, "amot", hasher.GetAllowedKeys()) // make sure they are sorted alphabetically

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestHashSha512_224(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha512-224")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("34278cf3456954a9133745d308a3ab1a9d8c82f3f7b73ff4ecb6977e")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestHashSha512_256(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha512_256") // verify this is handled as "sha512-256"
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("6f00138d3734506cd4b98adf20d9f6262e981e34ef56dfc8fb24d2daa89e7041")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestHashSha512Onion(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha512")
	hasher.AllowedKeys(AllowedOnion)
	hasher.Suffix(".onion")
	hasher.ExpectedHash("f07be23625ad049e9c44d9d2a8088d3902f5dbbd3f16a1469c34051d5987c5859fc1eeb0127764ad1ba1de4da51297002baaa1b41f3e259d54b135434d8851cc")
	hasher.Length(16)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "222222222222222f.onion", string(res))
}

func TestHashSha512OnionRandom(t *testing.T) {

	rand.Seed(123)

	hasher := NewHasher()
	hasher.Algo("sha512")
	hasher.AllowedKeys(AllowedOnion)
	hasher.Suffix(".onion")
	hasher.ExpectedHash("bbc3581fa536cf90d95b60d226495d38257d73e971b3193cc3fd532338caba7710966c5c91eddc8c1193e9cf401db94cb7c16205f064b6c45e3320d8c5d0b5f3")
	hasher.Length(16)

	res, err := hasher.FindRandom()
	assert.Equal(t, nil, err)
	assert.Equal(t, "2gl57brnwcjqmaua.onion", string(res))
}

func TestHashSha512OnionReverse(t *testing.T) {

	rand.Seed(123)

	hasher := NewHasher()
	hasher.Algo("sha512")
	hasher.AllowedKeys(AllowedOnion)
	hasher.Suffix(".onion")
	hasher.ExpectedHash("4e73702fa409f71f7a564276998b5c663e0617d301dc2f6f79ee4b58d18794eea8449e3a385360e774be22f970a7127a4117ba41a576cab2f46704fd0b6b29e0")
	hasher.Length(16)
	hasher.Reverse(true)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "zzzzzzzzzzzzzzww.onion", string(res))
}

// benchmarks given key length and print a prediction based on it
func BenchmarkSha1Speed(*testing.B) {

	hasher := NewHasher()

	hasher.Algo("sha1")
	hasher.AllowedKeys("580%(=QWI+qwi*Nn")
	hasher.ExpectedHash("0000000000000000000000000000000000000000")
	hasher.Length(5)

	hasher.FindSequential()
}
