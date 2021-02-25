package dice

import (
	crand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
)

func init() {
	b := make([]byte, 16)
	_, err := crand.Read(b)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(b))
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// Roll returns a random die number.
func Roll(size int) int {
	return rand.Intn(size) + 1
}
