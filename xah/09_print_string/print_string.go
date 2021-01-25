package main

import "fmt"

func main() {
	fmt.Println("abc"[0])         // 97
	fmt.Printf("%v\n", "abc"[0])  // 97
	fmt.Printf("%#v\n", "abc"[0]) // 0x61
	fmt.Printf("%q\n", "abc"[0])  // 'a'
	fmt.Println("---")

	fmt.Println("♥"[0])         // 226
	fmt.Println("♥"[1])         // 153
	fmt.Println("♥"[2])         // 165
	fmt.Printf("%v\n", "♥"[0])  // 226
	fmt.Printf("%#v\n", "♥"[0]) // oxe2
	fmt.Printf("%q\n", "♥"[0])  // 'a'
	fmt.Println("---")

	x := "♥\t😂"
	fmt.Printf("%s\n", x)         // ♥ 😂
	fmt.Printf("%q\n", x)         // "♥\t😂"
	fmt.Printf("%+q\n", x)        // "\u2665\t\U0001f602"
	fmt.Printf("% x\n", x)        // e2 99 a5 09 f0 9f 98 82
	fmt.Printf("%U\n", []rune(x)) // e2 99 a5 09 f0 9f 98 82

}
