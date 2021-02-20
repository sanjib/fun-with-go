package main

import (
	"fmt"
	"net/http"
)

var a int
var b float64 = 10
var c = 10
var d, e, f bool

var (
	g       int
	h       string
	i       int
	j, k, l bool
)

var c1 http.Cookie
var c2 *http.Cookie

func main() {
	var a int
	var b float64 = 10
	var c = 10
	var d, e, f bool

	var (
		g       int
		h       string
		i       int
		j, k, l bool
	)

	var c1 http.Cookie
	var c2 *http.Cookie

	fmt.Printf("a %#v\n", a)
	fmt.Printf("b %#v\n", b)
	fmt.Printf("c %#v\n", c)
	fmt.Printf("d %#v\n", d)
	fmt.Printf("e %#v\n", e)
	fmt.Printf("f %#v\n", f)
	fmt.Printf("g %#v\n", g)
	fmt.Printf("h %#v\n", h)
	fmt.Printf("i %#v\n", i)
	fmt.Printf("j %#v\n", j)
	fmt.Printf("k %#v\n", k)
	fmt.Printf("l %#v\n", l)

	fmt.Printf("c1 %#v\n", c1)
	fmt.Printf("c2 %#v\n", c2)
}
