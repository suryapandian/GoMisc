package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 1 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	before := 2.0
	for before != z {
		before = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(-100.0))
	fmt.Println(Sqrt(100))
}
