package main

import (
	"fmt"
	"time"
)

func main() {
	targetDay := time.Saturday
	fmt.Printf("When's %s?\n", targetDay)
	today := time.Now().Weekday()
	switch targetDay {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
