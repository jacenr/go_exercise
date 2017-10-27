package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) String() string {
	return fmt.Sprintf("String cannot Sqrt negative number: %.3f", e)
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error cannot Sqrt negative number: %.3f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(ErrNegativeSqrt(-2))
	fmt.Printf("%v\n", ErrNegativeSqrt(-2))
}
