package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(k chan bool, j int) {
			fmt.Println("This is i", j)
			for {
				select {
				case <-k:
					fmt.Println("stopping", j)
					return
				default:
				}
			}

		}(c, i)
	}

	time.Sleep(1 * time.Second)
	c <- true
	c <- true
	c <- true
	c <- true
	c <- true
	time.Sleep(5 * time.Second)
}
