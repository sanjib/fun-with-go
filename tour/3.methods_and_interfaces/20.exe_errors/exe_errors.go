package main

import "fmt"

// ErrNegativeSqrt type serves for error reporting.
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt calculates the square root of the given number.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	prev2, prev1 := 0.0, 0.0
	i := 1
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if z == prev1 || z == prev2 {
			fmt.Printf("took %d iterations\n", i)
			return z, nil
		}
		prev2 = prev1
		prev1 = z
		i++
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
