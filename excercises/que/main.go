package main

import (
	"fmt"
	"que.com/que"
)

func main() {
	queue := que.Queue{}
	queue.Add(1)
	queue.Add(8)
	queue.Add(5)
	queue.Delete()
	fmt.Println("Final queue", queue)
}
