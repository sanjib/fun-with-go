package main

import (
	"errors"
	"fmt"
	"os"
)

func writeFileTest() error {
	file, err := os.OpenFile("foo.txt", os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return fmt.Errorf("san err openfile: %w", err)
	}
	n, err := file.Write([]byte("hello foo"))
	if err != nil {
		return fmt.Errorf("san err write: %w", err)
	}
	fmt.Printf("wrote %d bytes\n", n)
	return nil
}

func main() {
	err := writeFileTest()
	if err != nil {
		//fmt.Println("top-level err:", err)
		//fmt.Println("top-level err:", errors.Unwrap(err))

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("hey error is!")
		}

		if errors.Is(err, errors.New("san err openfile: ")) {
			fmt.Println("hey error foo")
		}
	}
}
