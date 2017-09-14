package base91

import "testing"

var samples = map[string]string{
	"1":                           "xA",
	"1234567890":                  "QztEml0o[2;(A",
	"abcdefghijklmnopqurstuvwxyz": "#G(Ic,5ph#77&xrmlrjg2]jTs%2<WF%qfB",
}

func TestEncode(t *testing.T) {
	for k, v := range samples {
		if val := EncodeToString([]byte(k)); val != v {
			t.Errorf("Incorrect encoding of `%s` got `%s`", k, val)
		}
	}
}

func TestDecode(t *testing.T) {
	for k, v := range samples {
		if val := DecodeToString([]byte(v)); val != k {
			t.Errorf("Incorrect decoding of `%s` got `%s`", v, val)
		}
	}
}
