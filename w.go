package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(cx context.Context, j int) {

			for {
				select {
				case <-cx.Done():
					fmt.Println("stopping", j)
					return
				default:
					//fmt.Println("This is i", j)
				}
			}

		}(ctx, i)
	}

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}
