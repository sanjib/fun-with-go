package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (s rot13Reader) Read(b []byte) (int, error) {
	n, err := s.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case b[i]+13 > 'z':
			b[i] = 'a' + (b[i] + 13 - 'z' - 1)
		case b[i] == ' ':
			break
		case b[i] == '!':
			break
		default:
			b[i] = b[i] + 13
		}
	}
	if err != nil {
		return n, err
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
