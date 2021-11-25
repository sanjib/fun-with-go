package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type Capper struct {
	writer io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	upper := strings.ToUpper(string(p))
	return c.writer.Write([]byte(upper))
}

func main() {
	capper := Capper{os.Stdout}
	_, err := capper.Write([]byte("hello"))
	if err != nil {
		log.Println(err)

	}
}
