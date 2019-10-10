package main

import (
	"fmt"
	"time"
)


type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("ataaaa %v, %s",
		e.When, e.What)
}

func run() *MyError {
	a := &MyError{
		time.Now(),
		"it didn't work",
	}
	return a
}

func (e *MyError) String() string{
	return "surya"
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		fmt.Sprintf(err)
		fmt.Println(err.Error())
	}
	
}
