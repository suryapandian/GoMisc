package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	before := 2.0
	for ; before!=z;{
		before = z
		z -= (z*z - x) / (2*z)
		fmt.Println("Close",z)
	}
	return z

}

func main() {
	fmt.Println(Sqrt(100))
}
