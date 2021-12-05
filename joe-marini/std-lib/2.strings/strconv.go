package main

import (
	"log"
	"strconv"
)

func main() {
	sampleint := 100
	samplestr := "250"

	log.Println(string(sampleint))
	log.Println(strconv.Itoa(sampleint))
	log.Println(strconv.Atoi(samplestr))

	log.Println(strconv.ParseBool("true"))
	log.Println(strconv.ParseFloat("123", 64))
	log.Println(strconv.ParseFloat("123.11", 64))
	log.Println(strconv.ParseInt("123", 10, 64))
	log.Println(strconv.ParseUint("123", 10, 64))

	log.Println(strconv.FormatBool(true))
	log.Println(strconv.FormatFloat(123.4567, 'E', -2, 64))
	log.Println(strconv.FormatInt(123, 10))
	log.Println(strconv.FormatUint(123, 10))
}
