package main

import (
	"fmt"
	"strconv"
)

func encode(input string) (output string) {
	counter := 1
	char := string(input[0])
	for i := 1; i < len(input); i++ {
		fmt.Println(char, string(input[i]))
		if char == string(input[i]) {
			counter = counter + 1
			continue
		} else {
			n := ""
			if counter != 1 {
				n = strconv.Itoa(counter)
			}
			output = output + n + char
			counter = 1
		}
		char = string(input[i])
	}
	n := ""
	if counter != 1 {
		n = strconv.Itoa(counter)
	}
	output = output + n + char
	return
}

func encodeStream(input chan rune) (output chan string) {

	counter := 1
	c := <-input
	for true {
		nextChar := <-input
		if nextChar == '*' {
			break
		}

		if c == nextChar {
			counter = counter + 1
			continue
		} else {
			n := ""
			if counter != 1 {
				n = strconv.Itoa(counter)
			}
			output <- n + string(c)
			counter = 1
		}
		c = <-input
	}
	//close(input)

	n := ""
	if counter != 1 {
		n = strconv.Itoa(counter)
	}
	output <- n + string(c)
	return
}

func main() {
	fmt.Println(encode("aaabaaacd"))
	fmt.Println(encode("aaa"))
	fmt.Println(encode("zzzzmmmp"))
	inputString := make(chan rune)
	go func() { inputString <- 'a' }()
	go func() { inputString <- 'b' }()
	go func() { inputString <- 'b' }()
	go func() { inputString <- '*' }()
	encodeStream(inputString)

}
