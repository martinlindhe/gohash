package gohash

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexStringToBytes(t *testing.T) {

	assert.Equal(t, []byte("hej"), hexStringToBytes("68656a"))
}

func TestHashAdler32(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("adler32")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("026f0138")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}

func TestHashBlake224(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("blake224")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("c5d6e24c89a45385af97ae89c9edde904656d75e5a3582b1c9a390de")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}

func TestHashBlake256(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("blake256")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("91bff832dc57e964a521c660b6500ad04d565536fc5ccd98032bdcb1ebc9402c")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}

func TestHashBlake384(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("blake384")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("11a0ee2934bdd0f3c39ca0eee3b09287db24bc995df15d238da8d95f337ab39badcc6ca2dad0ba10cb49d32113f378b8")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}

func TestHashBlake512(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("blake512")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("3f0b354957782ac9f690683117d391bbd4d0b35061c21043e0915201a16fbf31a0dceac3d98b357a5624e93060df59e607b645a645f4bc944ef825aaf7022348")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}

func TestHashCrc32(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("crc32")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("0C68542E") // verify uppercase hex works
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}

/*
func TestHashCrc32c(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("crc32c")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("xxx")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}


func TestHashCrc32k(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("crc32c")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("b9fd977e")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}
*/

func TestHashMd2(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("md2")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("a8791e99e1a205db46c6cdef1f459108")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
}

func TestHashMd4(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("md4")
	hasher.AllowedKeys("holej")
	hasher.ExpectedHash("da3a901f9f5956d23553b9f00bc134a9")
	hasher.Length(3)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "hej", string(res))
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

func TestHashSha3_224(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha3-224")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("34f2791516dba95be06bf9bba074f771517cadbfe123a0b4af4b13af")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestHashSha3_256(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha3-256")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("bf74973956c0d73a25445523a3e5ac89f769be45bb9e3c5951e75af7007a94a1")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestHashSha3_384(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha3-384")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("13645c8aff8bb91beb80d1680861e9fe1ce76ca2a50c039922f4afbb4c61cf413b978fdd28eb846eeac4015f7b4bf3fd")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestHashSha3_512(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("sha3-512")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("14baea5d356a589fd2da38b66403716ed50f528fe500d2453fd409681894c840b72f9103dc58c1a8be666e5966ea16b24c5a091b63041598c5e93bad2378ee12")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestHashShake128(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("shake128-256")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("4d4a61f7f33e9d7166e9690239f5e64c6c1e0bb790b9e19bb7b2c466315463d7")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestHashShake256(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("shake256-512")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("1bdf1af0ab2a906727f0a5b99a30808fc41b6ebc150f516965e12a9f1067e50788905a3fc3bc8df24fcb5248e96a92a5f214aeac74bf2eb0dec6d070a5a5968b")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

/*
func TestHashSkein512_512(t *testing.T) {
// XXX dont work
	hasher := NewHasher()
	hasher.Algo("skein512-512")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("d678cbd86810bb6d7f376f76722c7bbda1602b19ec185d3b9faa5e49d97a98d5f69cc407c5252c229fc4c6407817d1ce60f84485c66bbdb2913d4706a27feb33")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}
*/
func TestHashWhirlpool(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("whirlpool")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("a42b7d2f481d91330c77855bcc935805ddb20e6096412e0d115918984711495336bad4938e7b74568c9b532adfae497819512efcdd21147b38ab8f6ac09bb5d1")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestHashTiger192(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("tiger192")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("7d4bb5a29d4bc1fb0e3070057f20c6498ba96872ce69ec25")
	hasher.Length(4)

	res, err := hasher.FindSequential()
	assert.Equal(t, nil, err)
	assert.Equal(t, "atom", string(res))
}

func TestRipemd160(t *testing.T) {

	hasher := NewHasher()
	hasher.Algo("ripemd160")
	hasher.AllowedKeys("mota")
	hasher.ExpectedHash("5bbb79e23ab1d617bd97ef9a3f289ddf5545a5c1")
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
