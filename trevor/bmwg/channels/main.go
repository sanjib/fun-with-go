package main

import (
	"github.com/sanjib/fake_channels/helpers"
	"log"
)

const numPool = 1000

func calVal(ci chan int) {
	rand := helpers.RandNum(numPool)
	ci <- rand
}

func main() {
	ci1 := make(chan int)
	defer close(ci1)

	go calVal(ci1)

	log.Println(<-ci1)
}
