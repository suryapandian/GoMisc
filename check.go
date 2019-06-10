package main

import (
	"fmt"
	"math"
	"math/rand"
)

var i, j int = 1, 2
//Global variable declaration

const Pizza = "PIZZA"
//Constants cannot be declared using the := syntax.

func add(x int, y int) int {
// func add(x, y int) int { - this also correct
	return x + y
}


//function parameters of same type
func swap(x, y string) (string, string) {
	return y, x
}

//naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var c, python, java = true, false, "no!"  //local variable declaration
	k := 3 //short hand function - inside function
	fmt.Println("My favorite number is", rand.Intn(10), c, python,java)
	fmt.Println(math.Pi) //will through error, all imported variables should be accessed with capital letter
	fmt.Println(add(42, 13))
	a, b := swap("hello1", "world")
	p := &a         // point to i
	fmt.Println(a,b)
	sum := 0 + 1 + j + k
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	x :=1.0
	n :=2.0
	fmt.Println("Aappa",math.Pow(x, n) > 100.00)
	if v := math.Pow(x, n); v > 100.00 {
		fmt.Sprint("www",v)
	}else {
		fmt.Println("OK",v)
		fmt.Sprint("www1",v)
	}
	sum2 := 1
	for ; sum2 < 1000; {
	//optional inti and increment statement
		sum2 += sum2
	}
	fmt.Println("aaa",sum2)

	//can drop the semicolons as well -> for is the while in Go
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	
	if sum < 0 {
		base := sum + 1
		fmt.Println("sss",base)
	}
	

}
