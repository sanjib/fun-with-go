package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var from string
	var to string

	flag.StringVar(&from, "from", "defFrom", "The unit to convert from")
	flag.StringVar(&to, "to", "defTo", "The unit to convert to")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please provide the value to convert")
		return
	}

	var num float64
	//_, err := fmt.Sscanf(args[0], "%f", &num)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	num, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	var ans float64
	switch {
	case from == "feet" && to == "meters":
		ans = num * 0.3048
	case from == "meters" && to == "feet":
		ans = num * 3.28084
	default:
		fmt.Println("Unknown unit")
		return
	}
	fmt.Printf("%.2f %s = %.2f %s\n", num, from, ans, to)
}
