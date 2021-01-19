package main

import (
	cryto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
)

func seed() int64 {
	var b [8]byte
	if _, err := cryto_rand.Read(b[:]); err != nil {
		fmt.Println(err)
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}

func main() {
	math_rand.Seed(seed())

	for i := 0; i < 10; i++ {
		r := math_rand.Intn(1e3)
		fmt.Println("random", r)
	}
}
