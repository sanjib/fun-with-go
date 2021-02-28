package main

import "fmt"

func main() {
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for i, v := range friends {
		fmt.Printf("val %s, add %p, i add %p\n", v, &v, &friends[i])
	}

}
