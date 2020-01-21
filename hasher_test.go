package gohash

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	allowedOnion = "abcdefghijklmnopqrstuvwxyz234567"
)

type expectedSeqRev struct {
	length      int
	allowedKeys string
	hash        string
	decoded     string
}

var (
	sequentialHasherTest = map[string]expectedSeqRev{
		"adler32":      {3, "holej", "026f0138", "hej"},
		"blake224":     {3, "holej", "c5d6e24c89a45385af97ae89c9edde904656d75e5a3582b1c9a390de", "hej"},
		"blake256":     {3, "holej", "91bff832dc57e964a521c660b6500ad04d565536fc5ccd98032bdcb1ebc9402c", "hej"},
		"blake384":     {3, "holej", "11a0ee2934bdd0f3c39ca0eee3b09287db24bc995df15d238da8d95f337ab39badcc6ca2dad0ba10cb49d32113f378b8", "hej"},
		"blake512":     {3, "holej", "3f0b354957782ac9f690683117d391bbd4d0b35061c21043e0915201a16fbf31a0dceac3d98b357a5624e93060df59e607b645a645f4bc944ef825aaf7022348", "hej"},
		"crc32-ieee":   {3, "holej", "0c68542e", "hej"},
		"md2":          {3, "holej", "a8791e99e1a205db46c6cdef1f459108", "hej"},
		"md4":          {3, "holej", "da3a901f9f5956d23553b9f00bc134a9", "hej"},
		"md5":          {3, "holej", "541c57960bb997942655d14e3b9607f9", "hej"},
		"ripemd160":    {4, "mota", "5bbb79e23ab1d617bd97ef9a3f289ddf5545a5c1", "atom"},
		"sha1":         {5, "tom", "9af7d87edaba03e23f6dbdaed29101ee1291c8a6", "motom"},
		"sha224":       {4, "bexlopslmv", "0136053cfe4adcb6b0130742b5abcdd086b79bc367a9e27eeedc4fe5", "slop"},
		"sha256":       {3, "aeltx", "a0b04eaca76db465982af821d6a304c3b904a7f13d6a9704d135aa07b3f1f6c2", "tex"},
		"sha384":       {4, "bexlopslmv", "35dc3a0822d431161309a0fe460786618c7a6fd3b0d2883e03fc21c851000ce2817bad134c1a72e179ada6fc207b6f15", "slop"},
		"sha512":       {4, "mota", "2b4df6c7b86d49a71c5ad6c1ffc2e85fde618c69400d2a0ccabb8dd12df4ae2584103b6947379742c0bc11a4e81ad4a3a832c11a734bf8ae5f8a8af9317a4c03", "atom"},
		"sha512-224":   {4, "mota", "34278cf3456954a9133745d308a3ab1a9d8c82f3f7b73ff4ecb6977e", "atom"},
		"sha512-256":   {4, "mota", "6f00138d3734506cd4b98adf20d9f6262e981e34ef56dfc8fb24d2daa89e7041", "atom"},
		"sha3-224":     {4, "mota", "34f2791516dba95be06bf9bba074f771517cadbfe123a0b4af4b13af", "atom"},
		"sha3-256":     {4, "mota", "bf74973956c0d73a25445523a3e5ac89f769be45bb9e3c5951e75af7007a94a1", "atom"},
		"sha3-384":     {4, "mota", "13645c8aff8bb91beb80d1680861e9fe1ce76ca2a50c039922f4afbb4c61cf413b978fdd28eb846eeac4015f7b4bf3fd", "atom"},
		"sha3-512":     {4, "mota", "14baea5d356a589fd2da38b66403716ed50f528fe500d2453fd409681894c840b72f9103dc58c1a8be666e5966ea16b24c5a091b63041598c5e93bad2378ee12", "atom"},
		"shake128-256": {4, "mota", "4d4a61f7f33e9d7166e9690239f5e64c6c1e0bb790b9e19bb7b2c466315463d7", "atom"},
		"shake256-512": {4, "mota", "1bdf1af0ab2a906727f0a5b99a30808fc41b6ebc150f516965e12a9f1067e50788905a3fc3bc8df24fcb5248e96a92a5f214aeac74bf2eb0dec6d070a5a5968b", "atom"},
		"tiger192":     {4, "mota", "7d4bb5a29d4bc1fb0e3070057f20c6498ba96872ce69ec25", "atom"},
		"whirlpool":    {4, "mota", "a42b7d2f481d91330c77855bcc935805ddb20e6096412e0d115918984711495336bad4938e7b74568c9b532adfae497819512efcdd21147b38ab8f6ac09bb5d1", "atom"},
	}
)

func TestSequentialHasher(t *testing.T) {
	for algo, rev := range sequentialHasherTest {
		hasher := NewHasher()
		hasher.Algo(algo)
		hasher.Length(rev.length)
		hasher.AllowedKeys(rev.allowedKeys)
		hasher.ExpectedHash(rev.hash)

		res, err := hasher.FindSequential()
		assert.Equal(t, nil, err)
		assert.Equal(t, rev.decoded, string(res))
	}
}

func TestHashSequential(t *testing.T) {
	hasher := NewHasher()
	hasher.Algo("sha512")
	hasher.AllowedKeys(allowedOnion)
	hasher.Suffix(".onion")
	hasher.ExpectedHash("f07be23625ad049e9c44d9d2a8088d3902f5dbbd3f16a1469c34051d5987c5859fc1eeb0127764ad1ba1de4da51297002baaa1b41f3e259d54b135434d8851cc")
	hasher.Length(16)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "222222222222222f.onion", string(res))
}

func TestHashRandom(t *testing.T) {
	rand.Seed(123)

	hasher := NewHasher()
	hasher.Algo("sha1")
	hasher.AllowedKeys(allowedOnion)
	hasher.Suffix(".xxx")
	hasher.ExpectedHash("58d16ee2c8214cad052194d68d31384c9f2e4e57")
	hasher.Length(16)

	res, err := hasher.FindRandom()
	assert.Equal(t, nil, err)
	assert.Equal(t, "aawiioowvgzolbqa.xxx", string(res))
}

func TestHashReverse(t *testing.T) {

	rand.Seed(123)

	hasher := NewHasher()
	hasher.Algo("sha512")
	hasher.AllowedKeys(allowedOnion)
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

	_, _ = hasher.FindSequential()
}
