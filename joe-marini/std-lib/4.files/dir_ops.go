package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// create dir
	err := os.Mkdir("testdir2", 755)
	if err != nil {
		log.Println(err)
	}

	// create nested dir
	err = os.MkdirAll("foo2/bar/baz", 755)
	if err != nil {
		log.Println(err)
	}

	// remove single
	err = os.Remove("testdir")
	if err != nil {
		log.Println(err)
	}

	// remove all / multiple
	err = os.RemoveAll("foo")
	if err != nil {
		log.Println(err)
	}

	// get current working dir
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("pwd:", pwd)

	// get dir of current process
	exeDir, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("exe dir:", exeDir)

	// get contents of dir
	fileInfo, err := ioutil.ReadDir(".")
	if err != nil {
		log.Println(err)
	}
	for _, fi := range fileInfo {
		if fi.IsDir() {
			fmt.Print("Dir\t")
		} else {
			fmt.Print("File\t")
		}
		fmt.Println(fi.Name())
	}

}
