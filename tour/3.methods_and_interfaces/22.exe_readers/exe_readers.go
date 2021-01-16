package main

import "golang.org/x/tour/reader"

// MyReader is an empty struct
type MyReader struct{}

// Read is a method of MyReader
func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
