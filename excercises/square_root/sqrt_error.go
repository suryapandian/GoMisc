package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return "caanot Sqrt negative number:"
}

func Sqrt(x ErrNegativeSqrt) ErrNegativeSqrt {
	var z,before ErrNegativeSqrt = 1.0,2.0
	for ; before!=z;{
		before = z
		z -= (z*z - x) / (2*z)
	}
	return z

}


func main() {
	var a,b ErrNegativeSqrt = 100.0,-4.0
	fmt.Println(Sqrt(a))
	fmt.Println(b)
}