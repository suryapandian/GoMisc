import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var temp, longestString string
	for _, l := range s {
		v := string(l)
		if strings.ContainsAny(temp, v) {
			i := strings.Index(temp, v)
			temp = temp[i+1:] + v

		} else {
			temp = temp + v
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