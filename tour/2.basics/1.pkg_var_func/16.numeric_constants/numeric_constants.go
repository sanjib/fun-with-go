package main

import "fmt"

const (
	// Big denotes a big number
	Big = 1 << 100
	// Small denotes a  small number
	Small = Big >> 100
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Printf("%v %T\n", Small, Small)
	fmt.Printf("%v %T\n", needInt(Small), needInt(Small))
	fmt.Printf("%v %T\n", needFloat(Small), needFloat(Small))

	// fmt.Printf("%v %T\n", Big, Big)
	// fmt.Printf("%v %T\n", needInt(Big), needInt(Big))
	fmt.Printf("%v %T\n", needFloat(Big), needFloat(Big))
}
