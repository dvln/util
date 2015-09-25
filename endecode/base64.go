// Base 64 encoding/decoding helper, see gonuts discussion:
//   https://groups.google.com/forum/#!topic/golang-nuts/yTm2kfKniqU

package endecode 

import ( 
	"encoding/base64" 
) 

// Encode makes it a bit easier to deal with base 64 encoding, see
// example code below.
func Encode(encBuf, bin []byte, e64 *base64.Encoding) []byte { 
	maxEncLen := e64.EncodedLen(len(bin)) 
	if encBuf == nil || len(encBuf) < maxEncLen { 
		encBuf = make([]byte, maxEncLen) 
	} 
	e64.Encode(encBuf, bin) 
	return encBuf[0:] 
} 

// Decode makes it a bit easier to deal with base 64 decoding, see
// example code below.
func Decode(decBuf, enc []byte, e64 *base64.Encoding) []byte { 
	maxDecLen := e64.DecodedLen(len(enc)) 
	if decBuf == nil || len(decBuf) < maxDecLen { 
		decBuf = make([]byte, maxDecLen) 
	} 
	n, err := e64.Decode(decBuf, enc) 
	_ = err 
	return decBuf[0:n] 
} 

// example code for encoding/decoding:
/* 
func main() { 
	// binary data 
	bin := []byte("testing base64..") 
	// base64 standard encoding 
	e64 := base64.StdEncoding 
	// encode 
	enc := Encode(nil, bin, e64) 
	// decode 
	dec := Decode(nil, enc, e64) 
	// results 
	fmt.Println("Equal:", string(bin) == string(dec)) 
	fmt.Println("bin:", len(bin), bin) 
	fmt.Println("enc:", len(enc), string(enc)) 
	fmt.Println("dec:", len(dec), dec) 
} 

Which, using the idiomatic Go slice, gives the following neat and 
tidy 
and intuitive results. 
bin: 16 [116 101 115 116 105 110 103 32 98 97 115 101 54 52 46 46] 
enc: 24 dGVzdGluZyBiYXNlNjQuLg== 
dec: 16 [116 101 115 116 105 110 103 32 98 97 115 101 54 52 46 46] 
*/

