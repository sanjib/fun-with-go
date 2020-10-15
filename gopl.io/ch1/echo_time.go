// Experiment to measure the difference in running time between 
// our potentially inefficient versions and the one that uses 
// strings.Join.
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	args := []string{"foo", "bar", "baz"}
	loopCount := 10_000_000

	// Execution time for echoFor
	start := time.Now()
	for i := 0; i < loopCount; i++ {
		echoFor(args)
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("echoFor: %f secs\n", secs)

	// Execution time for echoFor
	start = time.Now()
	for i := 0; i < loopCount; i++ {
		echoJoin(args)
	}
	secs = time.Since(start).Seconds()
	fmt.Printf("echoJoin: %f secs\n", secs)
}

func echoFor(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	//fmt.Println(s)
}

func echoJoin(args []string) {
	strings.Join(args, " ")
	//fmt.Println(strings.Join(args, " "))
}
