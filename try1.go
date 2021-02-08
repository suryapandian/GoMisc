package main

import (
    "fmt"
)

func main() {
    str := "50552"
    fmt.Printf("Solution to %s is %d  %d\n", str, Solution(str), 1)
    str = "10101"
    fmt.Printf("Solution to %s is %d %d\n", str, Solution(str), 0)
    str = "88"
    fmt.Printf("Solution to %s is %d %d\n", str, Solution(str), 0)
    str = "79401312"
    fmt.Printf("Solution to %s is %d %d\n", str, Solution(str), 5)
    str = "0604040291"
    fmt.Printf("Solution to %s is %d %d\n", str, Solution(str), 3)
}

// func Solution(S string) int {
//     charMap := make(map[rune]int)
//     for _, s := range S {
//         charMap[s] += 1
//     }

//     var result int
//     numMap := make(map[int]rune)
//     for k, v := range charMap {
//         if _, ok := numMap[v]; !ok {
//             numMap[v] = k
//         } else {
//             for i := v - 1; i >= 0; i-- {
//                 result += 1
//                 if _, ok := numMap[i]; !ok {
//                     numMap[i] = k
//                     break
//                 }
//             }
//         }
//     }
//     return result
// }

func Solution(S string) int {
    charMap := make(map[rune]int)
    for _, s := range S {
        charMap[s] += 1
    }

    var result int
    numMap := make(map[int]rune)
    for k, v := range charMap {
        for i := v; i >= 0; i-- {
            if _, ok := numMap[i]; (i == 0) || !ok {
                numMap[i] = k
                break
            }
            result += 1
        }
    }
    return result
}
