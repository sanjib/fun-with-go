package main

import "fmt"

func main() {
	f := 123.4567
	fmt.Printf("%.2f\n", f)    // 123.46
	fmt.Printf("%10f\n", f)    // 123.456700
	fmt.Printf("%10.2f\n", f)  //     123.46
	fmt.Printf("%010.2f\n", f) // 0000123.46
	fmt.Printf("%+10.2f\n", f) //    +123.46
}
