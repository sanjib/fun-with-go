package main

import (
	"fmt"
	"strings"
	"time"
)

type Planet struct {
	Name             string
	Mass             int64
	Diameter         int
	Gravity          float64
	RotationPeriod   time.Duration
	HasAtmosphere    bool
	HasMagneticField bool
	Satellites       []string
	next, previous   *Planet
}

type CelestialBody struct {
	Name           string
	Mass           int64
	Diameter       int
	Gravity        float64
	RotationPeriod time.Duration
}

func uppercase(p *Planet) {
	p.Name = strings.ToUpper(p.Name)
}

func main() {
	var earth, jupiter Planet

	mars := Planet{
		Name:           "Mars",
		Mass:           642e15,
		Diameter:       6792,
		RotationPeriod: 24.7 * 60 * time.Minute,
		HasAtmosphere:  true,
		Satellites:     []string{"Phobos", "Deimos"},
		next:           &earth,
		previous:       &jupiter,
	}
	fmt.Printf("mars %v\n", mars)

	var sun, moon CelestialBody

	sun.Name = "Sun"
	moon.Name = "Moon"

	fmt.Println("sun == moon", sun == moon)

	uppercase(&mars)
	fmt.Println("mars", mars)
}
