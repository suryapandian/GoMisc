package main

import "fmt"
import "strings"

func letterCasePermutation(S string) []string {
	var results []string
	n := len(S)
	for i := 0; i < n; i++ {
		char := string(S[i])
		hasLetters := strings.ContainsAny(strings.ToUpper(char), "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		if hasLetters {
			results = append(results, S[0:i]+strings.ToUpper(char)+S[i+1:n])
			results = append(results, S[0:i]+strings.ToLower(char)+S[i+1:n])
		}

	}
	return results
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}
