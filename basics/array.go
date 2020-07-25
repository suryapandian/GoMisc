package main


import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	q := []int{2, 3, 5, 7, 11, 13} //slice
	fmt.Println(q)

	//When slicing, you may omit the high or low bounds to use their defaults instead.
	//The default is zero for the low bound and the length of the slice for the high bound.

	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:]
	fmt.Println(s)

	m := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(m)
}
