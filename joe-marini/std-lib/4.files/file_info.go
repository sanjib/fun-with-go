package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("file exists?", checkFileExists("./sample.txt"))
	stat, err := os.Stat("./sample.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("stat mod time:", stat.ModTime())
	fmt.Println("stat mode:", stat.Mode())
	mode := stat.Mode()
	fmt.Println("mode is regular?", mode.IsRegular())
	fmt.Println("stat size:", stat.Size())
	fmt.Println("stat is dir:", stat.IsDir())
}

func checkFileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
