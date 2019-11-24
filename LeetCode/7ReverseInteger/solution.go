package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var result, sign int
	sign = 1
	if x < 0 {
		sign = -1
		x = x * sign
	}

	for x > 0 {
		result = (result * 10) + (x % 10)
		x = x / 10
	}
	result = result * sign
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	fmt.Println("REsult", result)
	return result
}

func main() {
	reverse(91238)
	reverse(-123)
	reverse(120)
	reverse(1534236469)

}
