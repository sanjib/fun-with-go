package lib

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

// Seed aims to provide a true random number
func Seed() int64 {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		fmt.Println(err)
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}
