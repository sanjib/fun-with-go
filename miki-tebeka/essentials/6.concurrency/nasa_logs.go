package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func verifyMD5(md5str, filename string, ch chan string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	md5file := fmt.Sprintf("%x", h.Sum(nil))
	ok := md5file == md5str
	ch <- fmt.Sprintf("ok:%t; file:%s; string:%s", ok, md5file, md5str)
}

func main() {
	ch := make(chan string)
	bytes, err := ioutil.ReadFile("./nasa-logs/md5sum.txt")
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	scanner.Split(bufio.ScanLines)
	procs := 0
	for scanner.Scan() {
		procs++
		l := strings.Fields(scanner.Text())
		md5str := l[0]
		filename := l[1]
		go verifyMD5(md5str, "./nasa-logs/"+filename, ch)
	}

	for i := 0; i < procs; i++ {
		log.Println(<-ch)
	}

}
