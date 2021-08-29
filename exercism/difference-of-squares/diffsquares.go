package diffsquares

func SquareOfSum(n int) int {
	// sum := 0
	// for i := 1; i < (n+1); i++ {
	// 	sum = sum + i
	// }
	// return sum*sum

	r := (n + 1) * n / 2
	return r * r
}

func SumOfSquares(n int) int {
	// sum := 0
	// for i := 1; i < (n + 1); i++ {
	// 	sum = sum + (i * i)
	// }
	// return sum
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
