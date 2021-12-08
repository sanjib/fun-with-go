package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func handleErrFr(err error) {
	if err != nil {
		fmt.Println("err->", err)
	}
}

func getFileLen(filename string) int64 {
	stat, err := os.Stat(filename)
	handleErrFr(err)
	return stat.Size()
}

func main() {
	// read whole file
	fBytes, err := ioutil.ReadFile("sample.txt")
	handleErrFr(err)
	fmt.Println("foo1:", string(fBytes))

	// read in smaller chunks
	const buffSize = 20
	f, err := os.Open("sample.txt")
	handleErrFr(err)
	defer f.Close()
	buffer1 := make([]byte, buffSize)
	for {
		n, err := f.Read(buffer1)
		if err != nil {
			if err != io.EOF {
				handleErrFr(err)
			}
			break
		}
		fmt.Print("bytes read:", n, "; ")
		fmt.Println("content:", string(buffer1[:n]))
	}

	// get length of file
	fileSize := getFileLen("sample.txt")
	fmt.Println("file len/size", fileSize)

	// read at a specific index: last 10 bytes
	buffer2 := make([]byte, 10)
	_, err = f.ReadAt(buffer2, fileSize-int64(len(buffer2)))
	handleErrFr(err)
	fmt.Println("last 10 bytes:", string(buffer2))
}
