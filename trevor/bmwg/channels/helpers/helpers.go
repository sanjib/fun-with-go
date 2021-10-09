package helpers

import (
	"math/rand"
	"time"
)

func RandNum(n int) int {
	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(n)
	return val
}
