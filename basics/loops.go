package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// If syntatically like for
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//For like if

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//Switch
	today := time.Now().Weekday()
	fmt.Println("When's Saturday?",today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//Switch without condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
