package main

import "fmt"

func sum_numbers(a int32,b int32) int32{
	return a + b
}

func main() {
	var a, b int32
	fmt.Println("Sum of 64 + 82 = ", sum_numbers(64,82))
	fmt.Println("Enter two numbers to find their sum:")
	fmt.Scanf("%v\n%v", &a, &b)
	fmt.Println("Sum of the numbers are:", sum_numbers(a, b))
}
