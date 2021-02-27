package main

import (
	"fmt"
	"time"
)

type Planet2 struct {
	CelestialBody2
	Name             string
	HasAtmosphere    bool
	HasMagneticField bool
	Satellites       []string
	next, previous   *Planet2
}

type CelestialBody2 struct {
	Name           string
	Mass           int64
	Diameter       int
	Gravity        float64
	RotationPeriod time.Duration
}

func main() {
	mars := Planet2{}
	mars.CelestialBody2.Name = "Venus"
	mars.Name = "Mars"
	mars.CelestialBody2.Mass = 12184
	fmt.Printf("mars %v\n", mars)
	fmt.Printf("mars %+v\n", mars)

	pluto := Planet2{
		CelestialBody2: CelestialBody2{
			Name: "PLUTO",
		},
	}
	pluto.Name = "p.l.u.t.o."
	fmt.Println("pluto", pluto)
	fmt.Println("pluto name", pluto.CelestialBody2.Name, pluto.Name)
}
