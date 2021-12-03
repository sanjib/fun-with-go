// Usage: go run server_kill.go -file="./pid_file.txt"
package main

import (
	"flag"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	filename := flag.String("file", "", "Filename containing pid")
	flag.Parse()

	err := killServer(*filename)
	if err != nil {
		log.Println(err)
	}
}

func killServer(pidFile string) error {
	bytes, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "could not read file")
	}

	pid, err := strconv.Atoi(string(bytes))
	if err != nil {
		return errors.Wrap(err, "could not convert pid string to int")
	}
	log.Printf("killing pid: %d softly\n", pid)

	return nil
}
