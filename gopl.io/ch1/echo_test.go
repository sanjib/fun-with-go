// Experiment to measure the difference in running time between 
// our potentially inefficient versions and the one that uses 
// strings.Join.
// Usage: go test -bench=. echo_test.go
package main

import (
	"strings"
	"testing"
)

func BenchmarkEchoFor(b *testing.B) {
	args := []string{"foo", "bar", "baz"}
	for i := 0; i < b.N; i++ {
		echoFor(args)
	}
}

func BenchmarkEchoJoin(b *testing.B) {
	args := []string{"foo", "bar", "baz"}
	for i := 0; i < b.N; i++ {
		echoJoin(args)
	}
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
