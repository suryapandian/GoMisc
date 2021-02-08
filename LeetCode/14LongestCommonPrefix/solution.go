import "fmt"

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	if n == 1 {
		return strs[0]
	}

	commonWord := strs[0]
	for i := 1; i < n; i++ {
		commonTill := commonPrefix(commonWord, strs[i])
		if commonTill == 0 {
			return ""
		}
		commonWord = commonWord[:commonTill]
	}
	return commonWord
}

func commonPrefix(str1, str2 string) (commonTill int) {
	for i, j := 0, 0; (i < len(str1)) && j < len(str2); i, j = i+1, j+1 {
		if str1[i] != str2[j] {
			break
		}
		commonTill++
	}
	return commonTill
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flow", "flower", "flig"}))
}
