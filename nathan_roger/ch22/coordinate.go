package main

import "fmt"

func main() {
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	curiosity := newLocationDMS("Bradbury", lat, long)
	fmt.Println("curiosity location", curiosity)
}

type location struct {
	name      string
	lat, long float64
}

func newLocationDMS(name string, lat coordinate, long coordinate) location {
	return location{name, lat.decimal(), long.decimal()}
}

func newLocationDec(name string, lat float64, long float64) location {
	return location{name, lat, long}
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
