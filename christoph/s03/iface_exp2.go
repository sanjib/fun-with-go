package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func readFromServer(in io.Reader) {
	var buf bytes.Buffer

	n, err := buf.ReadFrom(in)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Read", n, "bytes:", buf.String())
}

func main() {
	conn, err := net.Dial("tcp", "viaye.san:8080")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	readFromServer(conn)
}
