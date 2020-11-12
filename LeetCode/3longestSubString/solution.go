import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(str string) int {
	var temp, longestString string
	for _, rune := range str {
		character := string(rune)
		if strings.ContainsAny(temp, character) {
			i := strings.Index(temp, character)
			temp = temp[i+1:] + character

		} else {
			temp = temp + character
		}
		if len(temp) > len(longestString) {
			longestString = temp
		}
	}
	return len(longestString)

}

func main() {
	lengthOfLongestSubstring("dvdf")
}

