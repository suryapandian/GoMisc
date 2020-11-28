package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MissingDigit(str string) string {
	inputs := strings.Split(str, " ")
	var answer string

	isNumber1, lhs := getNumber(inputs[0])
	isNumber2, rhs := getNumber(inputs[2])
	isNumber3, ans := getNumber(inputs[4])
	operator := inputs[1]
	//fmt.Println(isNumber1,isNumber2,isNumber3)

	if !isNumber3 {
		//fmt.Println("this is ans", ans)
		switch operator {
		case "+":
			ans = lhs + rhs
			break
		case "-":
			ans = lhs - rhs
			break
		case "*":
			ans = lhs * rhs
			break
		case "/":
			ans = lhs / rhs
		}
		index := getIndexOf(inputs[4])
		//fmt.Println("isNumber3", ans, index, answer)
		answer = strconv.Itoa(ans)
		if index == 0 {
			return answer
		}
		answer = answer[index : index+1]

	}

	if !isNumber2 {
		//fmt.Println("this is rhs", rhs)
		switch operator {
		case "+":
			rhs = ans - lhs
			break
		case "-":
			rhs = lhs - ans
			break
		case "*":
			rhs = ans / lhs
			break
		case "/":
			rhs = lhs / ans
		}
		index := getIndexOf(inputs[2])
		//fmt.Println("isNumber2", index)
		answer = strconv.Itoa(rhs)
		if index == 0 {
			return answer
		}
		answer = answer[index : index+1]
	}

	if !isNumber1 {
		//fmt.Println("this is lhs", lhs)
		switch operator {
		case "+":
			lhs = ans - rhs
			break
		case "-":
			lhs = ans + rhs
			break
		case "*":
			lhs = ans / rhs
			break
		case "/":
			lhs = ans * rhs
		}
		index := getIndexOf(inputs[0])
		//fmt.Println("isNumber1", index)
		answer = strconv.Itoa(lhs)
		if index == 0 {
			return answer
		}
		answer = answer[index : index+1]
	}

	//fmt.Printf(" LHS %d \n operator %s \n RHS %d \n ANS %d \n ", lhs, operator ,rhs, ans )
	//fmt.Println("ANSWER!!!!", answer)
	// code goes here
	return answer

}

func getNumber(str string) (isNumber bool, number int) {
	number, err := strconv.Atoi(str)
	if err != nil {
		return false, 0
	}
	isNumber = true
	return
}

func getIndexOf(str string) int {
	//fmt.Println("Index input", str)
	for i, c := range str {
		if c == 'x' {
			return i
		}
	}
	return 0
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(MissingDigit(readline()))

}
