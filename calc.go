package gohash

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io"
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
	"github.com/martinlindhe/gogost/gost28147"
	"github.com/martinlindhe/gogost/gost34112012256"
	"github.com/martinlindhe/gogost/gost34112012512"
	"github.com/martinlindhe/gogost/gost341194"
	"github.com/mewpkg/hashutil/crc8"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// Calculator is used to calculate hash of input cleartext
type Calculator struct {
	reader io.Reader
}

// NewCalculator creates a new Calculator
func NewCalculator(reader io.Reader) *Calculator {
	return &Calculator{
		reader: reader,
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
		"gost94-cryptopro":  256,
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
		"streebog-256":      256,
		"streebog-512":      512,
		"tiger192":          192,
		"whirlpool":         512,
	}

	hashers = map[string]func(io.Reader) ([]byte, error){
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
		"gost94-cryptopro":  gost94CryptoproSum,
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
		"streebog-256":      streebog256Sum,
		"streebog-512":      streebog512Sum,
		"tiger192":          tiger192Sum,
		"whirlpool":         whirlpoolSum,
	}
)

// Sum returns the checksum
func (c *Calculator) Sum(algo string) ([]byte, error) {
	algo = resolveAlgoAliases(algo)
	if checksum, ok := hashers[algo]; ok {
		return checksum(c.reader)
	}
	return nil, fmt.Errorf("%s", "FATAL: unknown algo "+algo)
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

	// "skein256" is used by sphsum
	if s == "skein256" {
		return "skein512-256"
	}

	// "skein512" is used by sphsum
	if s == "skein512" {
		return "skein512-256"
	}

	// "tiger" is used by rhash, sphsum
	if s == "tiger" {
		return "tiger192"
	}

	// "gost" is used by rhash
	if s == "gost" {
		return "gost94"
	}

	// streebog is sometimes referred to as GOST-2012
	if s == "gost2012-256" {
		return "streebog-256"
	}
	if s == "gost2012-512" {
		return "streebog-512"
	}

	return s
}

func adler32Sum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	i := adler32.Checksum(buf.Bytes())
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return bs, nil
}

func blake224Sum(r io.Reader) ([]byte, error) {
	h := blake256.New224()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func blake256Sum(r io.Reader) ([]byte, error) {
	h := blake256.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func blake384Sum(r io.Reader) ([]byte, error) {
	h := blake512.New384()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func blake512Sum(r io.Reader) ([]byte, error) {
	h := blake512.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func blake2b256Sum(r io.Reader) ([]byte, error) {
	h := blake2b.New256()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func blake2b512Sum(r io.Reader) ([]byte, error) {
	h := blake2b.New512()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func blake2s256Sum(r io.Reader) ([]byte, error) {
	h := blake2s.New256()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func crc8AtmSum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	i := crc8.ChecksumATM(buf.Bytes())
	bs := make([]byte, 1)
	bs[0] = i
	return bs, nil
}

func crc16CcittSum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	i := crc16.ChecksumCCITT(buf.Bytes())
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return bs, nil
}

func crc16CcittFalseSum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	i := crc16.ChecksumCCITTFalse(buf.Bytes())
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return bs, nil
}

func crc16IbmSum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	i := crc16.ChecksumIBM(buf.Bytes())
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return bs, nil
}

func crc16ScsiSum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	i := crc16.ChecksumSCSI(buf.Bytes())
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return bs, nil
}

func crc24OpenPGPSum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	i := crc24.ChecksumOpenPGP(buf.Bytes())
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	bs = bs[1:4]
	return bs, nil
}

func crc32IEEESum(r io.Reader) ([]byte, error) {
	h := crc32.New(crc32.MakeTable(crc32.IEEE))
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func crc32CastagnoliSum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	tbl := crc32.MakeTable(crc32.Castagnoli)
	i := crc32.Checksum(buf.Bytes(), tbl)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return bs, nil
}

func crc32KoopmanSum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	tbl := crc32.MakeTable(crc32.Koopman)
	i := crc32.Checksum(buf.Bytes(), tbl)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return bs, nil
}

func crc64ISOSum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	tbl := crc64.MakeTable(crc64.ISO)
	i := crc64.Checksum(buf.Bytes(), tbl)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, i)
	return bs, nil
}

func crc64ECMASum(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	tbl := crc64.MakeTable(crc64.ECMA)
	i := crc64.Checksum(buf.Bytes(), tbl)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, i)
	return bs, nil
}

func fnv1_32Sum(r io.Reader) ([]byte, error) {
	h := fnv.New32()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func fnv1a32Sum(r io.Reader) ([]byte, error) {
	h := fnv.New32a()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func fnv1_64Sum(r io.Reader) ([]byte, error) {
	h := fnv.New64()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func fnv1a64Sum(r io.Reader) ([]byte, error) {
	h := fnv.New64a()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func gost94Sum(r io.Reader) ([]byte, error) {
	h := gost341194.New(gost341194.SboxDefault)
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func gost94CryptoproSum(r io.Reader) ([]byte, error) {
	h := gost341194.New(&gost28147.GostR3411_94_CryptoProParamSet)
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func streebog256Sum(r io.Reader) ([]byte, error) {
	h := gost34112012256.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func streebog512Sum(r io.Reader) ([]byte, error) {
	h := gost34112012512.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func md2Sum(r io.Reader) (digest []byte, err error) {
	h := md2.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func md4Sum(r io.Reader) (digest []byte, err error) {
	h := md4.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func md5Sum(r io.Reader) (digest []byte, err error) {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func ripemd160Sum(r io.Reader) (digest []byte, err error) {
	h := ripemd160.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func sha1Sum(r io.Reader) (digest []byte, err error) {
	h := sha1.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha224Sum(r io.Reader) (digest []byte, err error) {
	h := sha256.New224()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha256Sum(r io.Reader) (digest []byte, err error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha384Sum(r io.Reader) (digest []byte, err error) {
	h := sha512.New384()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha512Sum(r io.Reader) (digest []byte, err error) {
	h := sha512.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha512_224Sum(r io.Reader) (digest []byte, err error) {
	h := sha512.New512_224()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha512_256Sum(r io.Reader) (digest []byte, err error) {
	h := sha512.New512_256()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha3_224Sum(r io.Reader) (digest []byte, err error) {
	h := sha3.New224()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha3_256Sum(r io.Reader) (digest []byte, err error) {
	h := sha3.New256()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha3_384Sum(r io.Reader) (digest []byte, err error) {
	h := sha3.New384()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func sha3_512Sum(r io.Reader) (digest []byte, err error) {
	h := sha3.New512()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(digest[:0]), nil
}

func shake128_256Sum(r io.Reader) ([]byte, error) {
	res := make([]byte, 32) // 256 bits
	h := sha3.NewShake128()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	h.Read(res)
	return res, nil
}

func shake256_512Sum(r io.Reader) ([]byte, error) {
	res := make([]byte, 64) // 512 bits
	h := sha3.NewShake256()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	h.Read(res)
	return res, nil
}

func siphash2_4Sum(r io.Reader) ([]byte, error) {
	key := make([]byte, 16) // NOTE using empty key
	h := siphash.New(key)
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func skein512_256Sum(r io.Reader) ([]byte, error) {
	h := skein.NewHash(32)
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func skein512_512Sum(r io.Reader) ([]byte, error) {
	h := skein.NewHash(64)
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func tiger192Sum(r io.Reader) ([]byte, error) {
	h := tiger.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func whirlpoolSum(r io.Reader) ([]byte, error) {
	h := whirlpool.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
