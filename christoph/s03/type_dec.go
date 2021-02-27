package main

import "fmt"

type km float64
type mile float64

func howLongDoINeedToWalk(distance km) {
	fmt.Println("You need to walk", distance/5.0*60, "minutes for", distance, "km.")
}

func milesToKM(m mile) km {
	return km(m / 0.621371)
}

func main() {
	var distance km = 12
	howLongDoINeedToWalk(distance)

	var journeyToBeach mile = 220
	fmt.Println(milesToKM(journeyToBeach))
}
