package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func writeFileData(filename string) {
	// create file
	f, err := os.Create(filename)
	handleErr(err)

	// close file later
	defer f.Close()

	// file name
	fmt.Println("file name:", f.Name())

	// write string
	_, err = f.WriteString("this is some text\n")
	handleErr(err)

	// write byte
	_, err = f.Write([]byte("hello world\n"))
	handleErr(err)

	// write byte
	_, err = f.Write([]byte{'a', 'b', 'c', '\n'})
	handleErr(err)

	// sync forces data to be written
	err = f.Sync()
	handleErr(err)
}

func appendFileData(filename, data string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err)
	defer f.Close()

	_, err = f.WriteString(data)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func checkFileExistsForFileWrite(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	// write whole file
	foo1 := []byte("this is some sample data\n")
	err := ioutil.WriteFile("foo1.txt", foo1, 0644)
	handleErr(err)

	// write granular
	if !checkFileExistsForFileWrite("foo2.txt") {
		writeFileData("foo2.txt")
	} else {
		// truncate file to 10 bytes
		err := os.Truncate("foo2.txt", 10)
		handleErr(err)
		fmt.Println("trimmed file data")
	}

	// append to file
	appendFileData("foo1.txt", "This data was appended")
}
