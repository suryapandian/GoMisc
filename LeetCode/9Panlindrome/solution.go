package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || ((x%10 == 0) && x > 0) {
		return false
	}
	var orig, rev int
	orig = x
	for orig > 0 {
		rev = (rev * 10) + (orig % 10)
		orig = orig / 10
		if rev == orig {
			return true
		}
	}
	return rev == x
}

func main() {
	fmt.Println(isPalindrome(5))
	fmt.Println(isPalindrome(50))
	fmt.Println(isPalindrome(72))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(11))

}
