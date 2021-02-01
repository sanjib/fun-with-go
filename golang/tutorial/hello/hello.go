package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	names := []string{"Sanjib", "Tasnima", "Qarnine"}
	msgs, err := greetings.Hellos(names)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(msgs)
}
