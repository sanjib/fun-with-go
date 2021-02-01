package main

import (
	cryto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
)

func init() {
	var b [8]byte
	_, err := cryto_rand.Read(b[:])
	if err != nil {
		fmt.Println("cryto_rand.Read(b[:] failed:", err)
	}
	seed := int64(binary.LittleEndian.Uint64(b[:]))
	math_rand.Seed(seed)
}

func main() {
	fmt.Println("My 1st favorite number is", math_rand.Intn(100))
	fmt.Println("My 2nd favorite number is", math_rand.Intn(100))
	fmt.Println("My 3rd favorite number is", math_rand.Intn(100))
	fmt.Println("My 4th favorite number is", math_rand.Intn(100))
}
