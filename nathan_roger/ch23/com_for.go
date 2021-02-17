package main

import "fmt"

func main() {
	//test1()
	test2()
}

func test2() {
	r1 := report{sol: 15}
	fmt.Println("days", r1.sol.days(1446))
	fmt.Println("days", r1.days(1446))
}

func test1() {
	bradbury := report{
		sol: 15,
		temperature: temperature{
			high: -1,
			low:  -78,
		},
		location: location{
			lat:  -4.5895,
			long: 137.4417,
		},
	}
	fmt.Printf("bradbury %+v\n", bradbury)
	fmt.Printf("ave temp %v\n", bradbury.temperature.average())
	fmt.Printf("ave temp %v\n", bradbury.average())
}

// report type is a struct
type report struct {
	sol
	temperature
	location
}

//func (r report) average() celsius {
//	return r.temperature.average()
//}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

// sol type is an int
type sol int

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

// temperature struct
type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// location type is a struct
type location struct {
	lat, long float64
}

func (l location) days(l2 location) int {
	return 5
}

// celsius type is a float64
type celsius float64
