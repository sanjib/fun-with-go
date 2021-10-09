package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	n1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(n1)
	log.Println(n1[6:9])

	s1 := []string{"one", "seven", "fish", "cat"}
	log.Println(s1)
	log.Println(s1[1:3])

}

func main4() {
	s1 := []string{
		"a", "b",
	}
	s1 = append(s1, "c", "d")
	log.Println(s1)

	s2 := []int{3, 2, 1, 4}
	sort.Ints(s2)
	log.Println(s2)
}

func main3() {
	u1 := make(map[string]User)
	u1["sanjib"] = User{
		FirstName: "Sanjib",
		LastName:  "Ahmad",
	}
	log.Println(u1)
}

func main2() {
	m1 := make(map[string]string)
	m1["dog"] = "Samson"
	m1["other-dog"] = "Cassie"
	log.Println(m1)

	m2 := make(map[string]int)
	m2["First"] = 1
	m2["Second"] = 2
	log.Println(m2)
}
