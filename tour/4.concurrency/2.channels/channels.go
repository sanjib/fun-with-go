package main

import (
	"fmt"
	"time"
)

func populateSlice(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = 1
	}
}

type chanSum struct {
	sum   int
	label string
}

func sum(s []int, c chan chanSum, label string) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- chanSum{sum, label}
}

// Execution data on Windows 10
// Base speed: 1.99 GHz
// Logical processors: 8
// Memory: 15.9 GB
//
// Slice len: 1_000_000_000
// {125000000 sum8} {125000000 sum7} {125000000 sum4} {125000000 sum2} {125000000 sum1} {125000000 sum5} {125000000 sum6} {125000000 sum3}
// total = 1000000000
// exe time: 5.1941479 secs
//
// Slice len: 2_000_000_000
// {250000000 sum8} {250000000 sum5} {250000000 sum6} {250000000 sum7} {250000000 sum4} {250000000 sum1} {250000000 sum2} {250000000 sum3}
// total = 2000000000
// exe time: 30.2281804 secs
//
// Slice len: 3_000_000_000
// {375000000 sum8} {375000000 sum7} {375000000 sum6} {375000000 sum5} {375000000 sum2} {375000000 sum4} {375000000 sum3} {375000000 sum1}
// total = 3000000000
// exe time: 64.8062659 secs
//
// Slice len: 4_000_000_000
// {500000000 sum7} {500000000 sum8} {500000000 sum5} {500000000 sum3} {500000000 sum6} {500000000 sum4} {500000000 sum2} {500000000 sum1}
// total = 4000000000
// exe time: 101.584404 secs
func main() {
	start := time.Now()
	// numCPU := runtime.NumCPU()

	s := make([]int, 1_000_000_000)
	c := make(chan chanSum)
	populateSlice(s)
	// fmt.Println(s)

	chunk1 := (len(s) / 8) * 1
	chunk2 := (len(s) / 8) * 2
	chunk3 := (len(s) / 8) * 3
	chunk4 := (len(s) / 8) * 4
	chunk5 := (len(s) / 8) * 5
	chunk6 := (len(s) / 8) * 6
	chunk7 := (len(s) / 8) * 7

	go sum(s[:chunk1], c, "sum1")
	go sum(s[chunk1:chunk2], c, "sum2")
	go sum(s[chunk2:chunk3], c, "sum3")
	go sum(s[chunk3:chunk4], c, "sum4")
	go sum(s[chunk4:chunk5], c, "sum5")
	go sum(s[chunk5:chunk6], c, "sum6")
	go sum(s[chunk6:chunk7], c, "sum7")
	go sum(s[chunk7:], c, "sum8")
	sum1 := <-c
	sum2 := <-c
	sum3 := <-c
	sum4 := <-c
	sum5 := <-c
	sum6 := <-c
	sum7 := <-c
	sum8 := <-c

	fmt.Printf("%v %v %v %v %v %v %v %v\ntotal = %d\nexe time: %v secs",
		sum1, sum2, sum3, sum4, sum5, sum6, sum7, sum8,
		sum1.sum+sum2.sum+sum3.sum+sum4.sum+sum5.sum+sum6.sum+sum7.sum+sum8.sum,
		time.Since(start).Seconds(),
	)
}
