package main

import "fmt"

func main() {
	var m1 map[string]bool
	fmt.Printf("m1 %#v\n", m1)

	m2 := new(map[string]bool)
	fmt.Printf("m2 %#v\n", m2)

	m3 := make(map[string]bool)
	fmt.Printf("m3 %#v\n", m3)

	//m := map[string]int{}
	var m map[string]int
	v := m["one"]
	fmt.Println("v", v)
	if v, ok := m["one"]; ok == false {
		fmt.Println(v, ok)
	}

	m3["ack"] = true
	fmt.Printf("m3 %#v\n", m3)

	//s1 := make([]int, 11)
	//s1[10] = 10
	//fmt.Printf("s1 %#v\n", s1)
}
