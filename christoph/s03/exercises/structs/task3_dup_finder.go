package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lines = make(map[string]bool)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename to scan")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \r\n\t")
		if line == "" {
			continue
		}
		if lines[line] {
			//fmt.Printf("%q\n", line)
			fmt.Println(line)
			continue
		}
		lines[line] = true
	}
}
