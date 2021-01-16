package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter type is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc safely increments SafeCounter with locks
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the value for given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
