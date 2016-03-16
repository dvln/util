package endecode

import (
	"encoding/base64"
	"testing"
)

// testing for endecode package, see:
//   https://groups.google.com/forum/#!topic/golang-nuts/yTm2kfKniqU
func TestEncodeDecode(t *testing.T) {
	// binary data
	bin := []byte("testing base64..")
	// base64 standard encoding
	e64 := base64.StdEncoding
	// encode
	enc := Encode(nil, bin, e64)
	// decode
	dec := Decode(nil, enc, e64)
	// results
	if string(bin) != string(dec) {
		t.Error("Base64 helper: Decoded value does not equal original encoded value")
	}
	if len(bin) != 16 {
		t.Error("Base64 helper: length of original string (bin) isn't the expected 16")
	}
	if len(enc) != 24 {
		t.Error("Base64 helper: encoded byte string not the expected length of 24")
	}
	if string(enc) != "dGVzdGluZyBiYXNlNjQuLg==" {
		t.Error("Base64 helper: encoded base64 string not set as expected")
	}
	if len(dec) != 16 {
		t.Error("Base64 helper: length of decoded string (dec) isn't the expected 16")
	}
}
