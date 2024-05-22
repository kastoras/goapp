package util_test

import (
	"encoding/hex"
	"goapp/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -cover "./pkg/util"
func TestRandGenerator(t *testing.T) {
	rstring := util.RandString(10)

	dst := make([]byte, hex.DecodedLen(len(rstring)))

	//decode hex to verify that it is a correct hex representation string
	_, err := hex.Decode(dst, []byte(rstring))
	assert.Nil(t, err)
}

// go test -bench="./pkg/util"
func BenchmarkRandGeneratorHex(b *testing.B) {
	number := 10
	for i := 0; i < b.N; i++ {
		util.RandString(number)
	}
}
