package main

import "fmt"

// the below function is an example of named return and multi return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main() {
	var i, j int = 1, 2
	var name,age,isLovely = "Surya", 25, true
	fmt.Println(split(17))
	fmt.Println(i, j, c, python, java)
	fmt.Println("GET TO KNOW ME",name,age,isLovely)
}
