package main

import (
	"errors"
	"log"
)

func main() {
	res, err := divide(100.0, 0)
	if err != nil {
		log.Println(err)
	}
	log.Println("res of div", res)
}

func divide(x, y float32) (float32, error) {
	var res float32

	if y == 0 {
		return res, errors.New("err: can't divide by zero")
	}

	res = x / y
	return res, nil
}
