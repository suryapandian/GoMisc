package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "50552"
    fmt.Printf("Solution to %s is %d \n", str, Solution(str))
    str = "10101"
    fmt.Printf("Solution to %s is %d \n", str, Solution(str))
    str = "88"
    fmt.Printf("Solution to %s is %d \n", str, Solution(str))
    str = "79401312"
    fmt.Printf("Solution to %s is %d \n", str, Solution(str))
    str = "0604040291"
    fmt.Printf("Solution to %s is %d \n", str, Solution(str))
}

func Solution(S string) int {
    n := len(S)
    var max int
    for i := 0; i < n-1; i++ {
        fmt.Println(i, n)
        num, _ := strconv.Atoi(S[i : i+2])
        fmt.Println(num)
        if num > max {
            max = num
        }
    }
    return max
}
