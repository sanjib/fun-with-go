package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// get temp dir
	tempDir := os.TempDir()
	fmt.Println("temp dir:", tempDir)

	// create temp file
	tempContent := []byte("some temp content for a temp file")
	tempFile, err := ioutil.TempFile(tempDir, "z_go_joemarini_temp_*.txt")
	if err != nil {
		log.Println(err)
	}
	_, err = tempFile.Write(tempContent)
	if err != nil {
		log.Println(err)
	}

	// read/print temp file
	tempContentRead, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		log.Println(err)
	}
	fmt.Println("temp content read back:", string(tempContentRead))

	// remove temp file
	fmt.Println("temp file name:", tempFile.Name())

	// create temp dir with ioutil
	tmpDir, err := ioutil.TempDir("", "z_go_joemarini_*")
	fmt.Println("tmp dir name", tmpDir)

	defer func() {
		err := tempFile.Close()
		if err != nil {
			log.Println(err)
		}
		err = os.Remove(tempFile.Name())
		if err != nil {
			log.Println(err)
		}
		err = os.RemoveAll(tmpDir)
		if err != nil {
			log.Println(err)
		}
	}()
}
