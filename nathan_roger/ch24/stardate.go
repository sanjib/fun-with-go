package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("stardate now earth", stardate(time.Now()))

	//var s sol = 1422
	s := sol(1422)
	fmt.Println("stardate 1422 mars", stardate(s))
}

type stardater interface {
	YearDay() int
	Hour() int
}

func stardate(t stardater) float64 {
	y := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + y + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}
