package gohash

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expectedHashes = map[string]expectedForms{
		"adler32": {
			fox:   "5bdc0fda",
			blank: "00000001"},
		"blake224": {
			fox:   "c8e92d7088ef87c1530aee2ad44dc720cc10589cc2ec58f95a15e51b",
			blank: "7dc5313b1c04512a174bd6503b89607aecbee0903d40a8a569c94eed"},
		"blake256": {
			fox:   "7576698ee9cad30173080678e5965916adbb11cb5245d386bf1ffda1cb26c9d7",
			blank: "716f6e863f744b9ac22c97ec7b76ea5f5908bc5b2f67c61510bfc4751384ea7a"},
		"blake384": {
			fox:   "67c9e8ef665d11b5b57a1d99c96adffb3034d8768c0827d1c6e60b54871e8673651767a2c6c43d0ba2a9bb2500227406",
			blank: "c6cbd89c926ab525c242e6621f2f5fa73aa4afe3d9e24aed727faaadd6af38b620bdb623dd2b4788b1c8086984af8706"},
		"blake512": {
			fox:   "1f7e26f63b6ad25a0896fd978fd050a1766391d2fd0471a77afb975e5034b7ad2d9ccf8dfb47abbbe656e1b82fbc634ba42ce186e8dc5e1ce09a885d41f43451",
			blank: "a8cfbbd73726062df0c6864dda65defe58ef0cc52a5625090fa17601e1eecd1b628e94f396ae402a00acc9eab77b4d4c2e852aaaa25a636d80af3fc7913ef5b8"},
		"blake2b-256": {
			fox:   "01718cec35cd3d796dd00020e0bfecb473ad23457d063b75eff29c0ffa2e58a9",
			blank: "0e5751c026e543b2e8ab2eb06099daa1d1e5df47778f7787faab45cdf12fe3a8"},
		"blake2b-512": {
			fox:   "a8add4bdddfd93e4877d2746e62817b116364a1fa7bc148d95090bc7333b3673f82401cf7aa2e4cb1ecd90296e3f14cb5413f8ed77be73045b13914cdcd6a918",
			blank: "786a02f742015903c6c6fd852552d272912f4740e15847618a86e217f71f5419d25e1031afee585313896444934eb04b903a685b1448b755d56f701afe9be2ce"},
		"blake2s-256": {
			fox:   "606beeec743ccbeff6cbcdf5d5302aa855c256c29b88c8ed331ea1a6bf3c8812",
			blank: "69217a3079908094e11121d042354a7c1f55b6482ca1a51e1b250dfd1ed0eef9"},
		"crc8-atm": {
			fox:   "c1",
			blank: "00"},
		"crc16-ccitt": {
			fox:   "9358",
			blank: "0000"},
		"crc16-ccitt-false": {
			fox:   "8fdd",
			blank: "ffff"},
		"crc16-ibm": {
			fox:   "5763",
			blank: "0000"},
		"crc16-scsi": {
			fox:   "b32b",
			blank: "0000"},
		"crc24-openpgp": {
			fox:   "3e7e22",
			blank: "000000"},
		"crc32-ieee": {
			fox:   "414fa339",
			blank: "00000000"},
		"crc32-castagnoli": {
			fox:   "22620404",
			blank: "00000000"},
		"crc32-koopman": {
			fox:   "e021db90",
			blank: "00000000"},
		"crc64-iso": {
			fox:   "4ef14e19f4c6e28e",
			blank: "0000000000000000"},
		"crc64-ecma": {
			fox:   "5b5eb8c2e54aa1c4",
			blank: "0000000000000000"},
		"fnv1-32": {
			fox:   "e9c86c6e",
			blank: "811c9dc5"},
		"fnv1a-32": {
			fox:   "048fff90",
			blank: "811c9dc5"},
		"fnv1-64": {
			fox:   "a8b2f3117de37ace",
			blank: "cbf29ce484222325"},
		"fnv1a-64": {
			fox:   "f3f9b7f5e7e47110",
			blank: "cbf29ce484222325"},
		"gost94": {
			fox:   "77b7fa410c9ac58a25f49bca7d0468c9296529315eaca76bd1a10f376d1f4294",
			blank: "ce85b99cc46752fffee35cab9a7b0278abb4c2d2055cff685af4912c49490f8d"},
		"gost94-cryptopro": {
			fox:   "9004294a361a508c586fe53d1f1b02746765e71b765472786e4770d565830a76",
			blank: "981e5f3ca30c841487830f84fb433e13ac1101569b9c13584ac483234cd656c0"},
		"streebog-256": {
			fox:   "3e7dea7f2384b6c5a3d0e24aaa29c05e89ddd762145030ec22c71a6db8b2c1f4",
			blank: "3f539a213e97c802cc229d474c6aa32a825a360b2a933a949fd925208d9ce1bb"},
		"streebog-512": {
			fox:   "d2b793a0bb6cb5904828b5b6dcfb443bb8f33efc06ad09368878ae4cdc8245b97e60802469bed1e7c21a64ff0b179a6a1e0bb74d92965450a0adab69162c00fe",
			blank: "8e945da209aa869f0455928529bcae4679e9873ab707b55315f56ceb98bef0a7362f715528356ee83cda5f2aac4c6ad2ba3a715c1bcd81cb8e9f90bf4c1c1a8a"},
		"md2": {
			fox:   "03d85a0d629d2c442e987525319fc471",
			blank: "8350e5a3e24c153df2275c9f80692773"},
		"md4": {
			fox:   "1bee69a46ba811185c194762abaeae90",
			blank: "31d6cfe0d16ae931b73c59d7e0c089c0"},
		"md5": {
			fox:   "9e107d9d372bb6826bd81d3542a419d6",
			blank: "d41d8cd98f00b204e9800998ecf8427e"},
		"ripemd160": {
			fox:   "37f332f68db77bd9d7edd4969571ad671cf9dd3b",
			blank: "9c1185a5c5e9fc54612808977ee8f548b2258d31"},
		"sha1": {
			fox:   "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12",
			blank: "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		"sha224": {
			fox:   "730e109bd7a8a32b1cb9d9a09aa2325d2430587ddbc0c38bad911525",
			blank: "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f"},
		"sha256": {
			fox:   "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592",
			blank: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		"sha384": {
			fox:   "ca737f1014a48f4c0b6dd43cb177b0afd9e5169367544c494011e3317dbf9a509cb1e5dc1e85a941bbee3d7f2afbc9b1",
			blank: "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b"},
		"sha512": {
			fox:   "07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6",
			blank: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
		"sha512-224": {
			fox:   "944cd2847fb54558d4775db0485a50003111c8e5daa63fe722c6aa37",
			blank: "6ed0dd02806fa89e25de060c19d3ac86cabb87d6a0ddd05c333b84f4"},
		"sha512-256": {
			fox:   "dd9d67b371519c339ed8dbd25af90e976a1eeefd4ad3d889005e532fc5bef04d",
			blank: "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a"},
		"sha3-224": {
			fox:   "d15dadceaa4d5d7bb3b48f446421d542e08ad8887305e28d58335795",
			blank: "6b4e03423667dbb73b6e15454f0eb1abd4597f9a1b078e3f5b5a6bc7"},
		"sha3-256": {
			fox:   "69070dda01975c8c120c3aada1b282394e7f032fa9cf32f4cb2259a0897dfc04",
			blank: "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a"},
		"sha3-384": {
			fox:   "7063465e08a93bce31cd89d2e3ca8f602498696e253592ed26f07bf7e703cf328581e1471a7ba7ab119b1a9ebdf8be41",
			blank: "0c63a75b845e4f7d01107d852e4c2485c51a50aaaa94fc61995e71bbee983a2ac3713831264adb47fb6bd1e058d5f004"},
		"sha3-512": {
			fox:   "01dedd5de4ef14642445ba5f5b97c15e47b9ad931326e4b0727cd94cefc44fff23f07bf543139939b49128caf436dc1bdee54fcb24023a08d9403f9b4bf0d450",
			blank: "a69f73cca23a9ac5c8b567dc185a756e97c982164fe25859e0d1dcc1475c80a615b2123af1f5f94c11e3e9402c3ac558f500199d95b6d3e301758586281dcd26"},
		"shake128-256": {
			fox:   "f4202e3c5852f9182a0430fd8144f0a74b95e7417ecae17db0f8cfeed0e3e66e",
			blank: "7f9c2ba4e88f827d616045507605853ed73b8093f6efbc88eb1a6eacfa66ef26"},
		"shake256-512": {
			fox:   "2f671343d9b2e1604dc9dcf0753e5fe15c7c64a0d283cbbf722d411a0e36f6ca1d01d1369a23539cd80f7c054b6e5daf9c962cad5b8ed5bd11998b40d5734442",
			blank: "46b9dd2b0ba88d13233b3feb743eeb243fcd52ea62b81b82b50c27646ed5762fd75dc4ddd8c0f200cb05019d67b592f6fc821c49479ab48640292eacb3b7c4be"},
		"siphash-2-4": {
			fox:   "0de4702506520059",
			blank: "d70077739d4b921e"},
		"skein512-256": {
			fox:   "b3250457e05d3060b1a4bbc1428bc75a3f525ca389aeab96cfa34638d96e492a",
			blank: "39ccc4554a8b31853b9de7a1fe638a24cce6b35a55f2431009e18780335d2621"},
		"skein512-512": {
			fox:   "94c2ae036dba8783d0b3f7d6cc111ff810702f5c77707999be7e1c9486ff238a7044de734293147359b4ac7e1d09cd247c351d69826b78dcddd951f0ef912713",
			blank: "bc5b4c50925519c290cc634277ae3d6257212395cba733bbad37a4af0fa06af41fca7903d06564fea7a2d3730dbdb80c1f85562dfcc070334ea4d1d9e72cba7a"},
		"tiger192": {
			fox:   "6d12a41e72e644f017b6f0e2f7b44c6285f06dd5d2c5b075",
			blank: "3293ac630c13f0245f92bbb1766e16167a4e58492dde73f3"},
		"whirlpool": {
			fox:   "b97de512e91e3828b40d2b0fdce9ceb3c4a71f9bea8d88e75c4fa854df36725fd2b52eb6544edcacd6f8beddfea403cb55ae31f03ad62a5ef54e42ee82c3fb35",
			blank: "19fa61d75522a4669b44e39c1d2e1726c530232130d407f89afee0964997f7a73e83be698b288febcf88e3e03c4f0757ea8964e59b63d93708b138cc42a66eb3"},
	}
)

func TestCalcExpectedHashes(t *testing.T) {
	for algo, forms := range expectedHashes {
		for form, hash := range forms {
			r := strings.NewReader(form)
			calc := NewCalculator(r)
			res, err := calc.Sum(algo)
			if err != nil {
				t.Error("FATAL algo", algo, "failed with", err)
			} else {
				assert.Equal(t, hash, hex.EncodeToString(res), algo+" '"+form+"'")
			}
		}
	}
}

func TestFuzzHashes(t *testing.T) {
	for algo := range expectedHashes {
		for i := 0; i < iterationsPerAlgo; i++ {
			var rnd []byte
			f.Fuzz(&rnd)
			r := bytes.NewReader(rnd)
			calc := NewCalculator(r)
			calc.Sum(algo)
		}
	}
}

func TestHashersAndAlgosDefines(t *testing.T) {
	for algo := range algos {
		if _, ok := hashers[algo]; !ok {
			t.Error("algo not defined in hashers map", algo)
		}
		if _, ok := expectedHashes[algo]; !ok {
			t.Error("algo lacks testcase in expectedHashes map", algo)
		}
	}
	for algo := range hashers {
		if _, ok := algos[algo]; !ok {
			t.Error("algo not defined in algos map", algo)
		}
	}
}

func BenchmarkHashes(b *testing.B) {
	for algo, forms := range expectedHashes {
		for form := range forms {
			b.Run(algo+" "+form, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					r := strings.NewReader(form)
					calc := NewCalculator(r)
					calc.Sum(algo)
				}
			})
		}
	}
}

func ExampleNewCalculator() {
	r := strings.NewReader("hello world")
	calc := NewCalculator(r)
	h, _ := calc.Sum("sha1")
	fmt.Printf("%x", h)
	// Output: 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
}

func BenchmarkCrc32IEEE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := strings.NewReader("dsfgklhkjsdhfgkjhsdkjfghljksdhfgjkhsdfgkjhjksdfhgkhsdfgksdfg")
		calc := NewCalculator(r)
		calc.Sum("crc32-ieee")
	}
}
