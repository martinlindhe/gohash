package gohash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"sort"

	"github.com/cxmcc/tiger"
	"github.com/dchest/blake256"
	"github.com/dchest/blake2b"
	"github.com/dchest/blake2s"
	"github.com/dchest/blake512"
	"github.com/dchest/siphash"
	"github.com/dchest/skein"
	"github.com/howeyc/crc16"
	"github.com/htruong/go-md2"
	"github.com/jzelinskie/whirlpool"
	"github.com/martinlindhe/crc24"
	"github.com/martinlindhe/gogost/gost341194"
	"github.com/mewpkg/hashutil/crc8"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// Calculator is used to calculate hash of input cleartext
type Calculator struct {
	data []byte
}

// NewCalculator creates a new Calculator
func NewCalculator(data []byte) *Calculator {

	return &Calculator{
		data: data,
	}
}

var (
	algos = map[string]int{
		// name, key size in bits
		"adler32":           32,
		"blake224":          224,
		"blake256":          256,
		"blake384":          384,
		"blake512":          512,
		"blake2b-256":       256,
		"blake2b-512":       512,
		"blake2s-256":       256,
		"crc8-atm":          8,
		"crc16-ccitt":       16,
		"crc16-ccitt-false": 16,
		"crc16-ibm":         16,
		"crc16-scsi":        16,
		"crc24-openpgp":     24,
		"crc32-ieee":        32,
		"crc32-castagnoli":  32,
		"crc32-koopman":     32,
		"crc64-iso":         64,
		"crc64-ecma":        64,
		"fnv1-32":           32,
		"fnv1a-32":          32,
		"fnv1-64":           64,
		"fnv1a-64":          64,
		"gost94":            256,
		"md2":               128,
		"md4":               128,
		"md5":               128,
		"ripemd160":         160,
		"sha1":              160,
		"sha224":            224,
		"sha256":            256,
		"sha384":            384,
		"sha512":            512,
		"sha512-224":        224,
		"sha512-256":        256,
		"sha3-224":          224,
		"sha3-256":          256,
		"sha3-384":          384,
		"sha3-512":          512,
		"shake128-256":      256,
		"shake256-512":      512,
		"siphash-2-4":       64,
		"skein512-256":      256,
		"skein512-512":      512,
		"tiger192":          192,
		"whirlpool":         512,
	}

	hashers = map[string]func(*[]byte) *[]byte{
		"adler32":           adler32Sum,
		"blake224":          blake224Sum,
		"blake256":          blake256Sum,
		"blake384":          blake384Sum,
		"blake512":          blake512Sum,
		"blake2b-256":       blake2b256Sum,
		"blake2b-512":       blake2b512Sum,
		"blake2s-256":       blake2s256Sum,
		"crc8-atm":          crc8AtmSum,
		"crc16-ccitt":       crc16CcittSum,
		"crc16-ccitt-false": crc16CcittFalseSum,
		"crc16-ibm":         crc16IbmSum,
		"crc16-scsi":        crc16ScsiSum,
		"crc24-openpgp":     crc24OpenPGPSum,
		"crc32-ieee":        crc32IEEESum,
		"crc32-castagnoli":  crc32CastagnoliSum,
		"crc32-koopman":     crc32KoopmanSum,
		"crc64-iso":         crc64ISOSum,
		"crc64-ecma":        crc64ECMASum,
		"fnv1-32":           fnv1_32Sum,
		"fnv1a-32":          fnv1a32Sum,
		"fnv1-64":           fnv1_64Sum,
		"fnv1a-64":          fnv1a64Sum,
		"gost94":            gost94Sum,
		"md2":               md2Sum,
		"md4":               md4Sum,
		"md5":               md5Sum,
		"ripemd160":         ripemd160Sum,
		"sha1":              sha1Sum,
		"sha224":            sha224Sum,
		"sha256":            sha256Sum,
		"sha384":            sha384Sum,
		"sha512":            sha512Sum,
		"sha512-224":        sha512_224Sum,
		"sha512-256":        sha512_256Sum,
		"sha3-224":          sha3_224Sum,
		"sha3-256":          sha3_256Sum,
		"sha3-384":          sha3_384Sum,
		"sha3-512":          sha3_512Sum,
		"shake128-256":      shake128_256Sum,
		"shake256-512":      shake256_512Sum,
		"siphash-2-4":       siphash2_4Sum,
		"skein512-256":      skein512_256Sum,
		"skein512-512":      skein512_512Sum,
		"tiger192":          tiger192Sum,
		"whirlpool":         whirlpoolSum,
	}
)

// Sum returns the checksum
func (c *Calculator) Sum(algo string) *[]byte {

	algo = resolveAlgoAliases(algo)

	if checksum, ok := hashers[algo]; ok {
		return checksum(&c.data)
	}
	return nil
}

// AvailableHashes returns the available hash id's
func AvailableHashes() []string {

	res := []string{}

	for key := range hashers {
		res = append(res, key)
	}

	sort.Strings(res)
	return res
}

func resolveAlgoAliases(s string) string {

	if s == "crc32" {
		return "crc32-ieee"
	}
	if s == "crc32c" {
		return "crc32-castagnoli"
	}
	if s == "crc32k" {
		return "crc32-koopman"
	}

	// "skein256" is used in sphsum
	if s == "skein256" {
		return "skein512-256"
	}

	// "skein512" is used in sphsum
	if s == "skein512" {
		return "skein512-256"
	}

	// "tiger" is used by rhash, sphsum
	if s == "tiger" {
		return "tiger192"
	}

	return s
}

func adler32Sum(b *[]byte) *[]byte {
	i := adler32.Checksum(*b)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return &bs
}

func blake224Sum(b *[]byte) *[]byte {
	w := blake256.New224()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func blake256Sum(b *[]byte) *[]byte {
	w := blake256.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func blake384Sum(b *[]byte) *[]byte {
	w := blake512.New384()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func blake512Sum(b *[]byte) *[]byte {
	w := blake512.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func blake2b256Sum(b *[]byte) *[]byte {
	x := blake2b.Sum256(*b)
	res := x[:]
	return &res
}

func blake2b512Sum(b *[]byte) *[]byte {
	x := blake2b.Sum512(*b)
	res := x[:]
	return &res
}

func blake2s256Sum(b *[]byte) *[]byte {
	x := blake2s.Sum256(*b)
	res := x[:]
	return &res
}

func crc8AtmSum(b *[]byte) *[]byte {
	i := crc8.ChecksumATM(*b)
	bs := make([]byte, 1)
	bs[0] = i
	return &bs
}

func crc16CcittSum(b *[]byte) *[]byte {
	i := crc16.ChecksumCCITT(*b)
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return &bs
}

func crc16CcittFalseSum(b *[]byte) *[]byte {
	i := crc16.ChecksumCCITTFalse(*b)
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return &bs
}

func crc16IbmSum(b *[]byte) *[]byte {
	i := crc16.ChecksumIBM(*b)
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return &bs
}

func crc16ScsiSum(b *[]byte) *[]byte {
	i := crc16.ChecksumSCSI(*b)
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return &bs
}

func crc24OpenPGPSum(b *[]byte) *[]byte {
	i := crc24.ChecksumOpenPGP(*b)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	bs = bs[1:4]
	return &bs
}

func crc32IEEESum(b *[]byte) *[]byte {
	i := crc32.ChecksumIEEE(*b)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return &bs
}

func crc32CastagnoliSum(b *[]byte) *[]byte {
	tbl := crc32.MakeTable(crc32.Castagnoli)
	i := crc32.Checksum(*b, tbl)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return &bs
}

func crc32KoopmanSum(b *[]byte) *[]byte {
	tbl := crc32.MakeTable(crc32.Koopman)
	i := crc32.Checksum(*b, tbl)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return &bs
}

func crc64ISOSum(b *[]byte) *[]byte {
	tbl := crc64.MakeTable(crc64.ISO)
	i := crc64.Checksum(*b, tbl)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, i)
	return &bs
}

func crc64ECMASum(b *[]byte) *[]byte {
	tbl := crc64.MakeTable(crc64.ECMA)
	i := crc64.Checksum(*b, tbl)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, i)
	return &bs
}

func fnv1_32Sum(b *[]byte) *[]byte {
	w := fnv.New32()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func fnv1a32Sum(b *[]byte) *[]byte {
	w := fnv.New32a()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func fnv1_64Sum(b *[]byte) *[]byte {
	w := fnv.New64()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func fnv1a64Sum(b *[]byte) *[]byte {
	w := fnv.New64a()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func gost94Sum(b *[]byte) *[]byte {
	h := gost341194.New(gost341194.SboxDefault)
	h.Write(*b)
	res := h.Sum(nil)
	return &res
}

func md2Sum(b *[]byte) *[]byte {
	w := md2.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func md4Sum(b *[]byte) *[]byte {
	w := md4.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func md5Sum(b *[]byte) *[]byte {
	x := md5.Sum(*b)
	res := x[:]
	return &res
}

func ripemd160Sum(b *[]byte) *[]byte {
	w := ripemd160.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func sha1Sum(b *[]byte) *[]byte {
	x := sha1.Sum(*b)
	res := x[:]
	return &res
}

func sha224Sum(b *[]byte) *[]byte {
	x := sha256.Sum224(*b)
	res := x[:]
	return &res
}

func sha256Sum(b *[]byte) *[]byte {
	x := sha256.Sum256(*b)
	res := x[:]
	return &res
}

func sha384Sum(b *[]byte) *[]byte {
	x := sha512.Sum384(*b)
	res := x[:]
	return &res
}

func sha512Sum(b *[]byte) *[]byte {
	x := sha512.Sum512(*b)
	res := x[:]
	return &res
}

func sha512_224Sum(b *[]byte) *[]byte {
	x := sha512.Sum512_224(*b)
	res := x[:]
	return &res
}

func sha512_256Sum(b *[]byte) *[]byte {
	x := sha512.Sum512_256(*b)
	res := x[:]
	return &res
}

func sha3_224Sum(b *[]byte) *[]byte {
	x := sha3.Sum224(*b)
	res := x[:]
	return &res
}

func sha3_256Sum(b *[]byte) *[]byte {
	x := sha3.Sum256(*b)
	res := x[:]
	return &res
}

func sha3_384Sum(b *[]byte) *[]byte {
	x := sha3.Sum384(*b)
	res := x[:]
	return &res
}

func sha3_512Sum(b *[]byte) *[]byte {
	x := sha3.Sum512(*b)
	res := x[:]
	return &res
}

func shake128_256Sum(b *[]byte) *[]byte {
	res := make([]byte, 32)
	sha3.ShakeSum128(res, *b)
	return &res
}

func shake256_512Sum(b *[]byte) *[]byte {
	res := make([]byte, 64)
	sha3.ShakeSum256(res, *b)
	return &res
}

func siphash2_4Sum(b *[]byte) *[]byte {
	key := make([]byte, 16) // NOTE using empty key
	w := siphash.New(key)
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func skein512_256Sum(b *[]byte) *[]byte {
	w := skein.NewHash(32)
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func skein512_512Sum(b *[]byte) *[]byte {
	w := skein.NewHash(64)
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func tiger192Sum(b *[]byte) *[]byte {
	w := tiger.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}

func whirlpoolSum(b *[]byte) *[]byte {
	w := whirlpool.New()
	w.Write(*b)
	res := w.Sum(nil)
	return &res
}
