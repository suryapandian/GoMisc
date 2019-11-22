package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(str string) (longestPalindrome string) {
	if len(str) > 0 {
		longestPalindrome = str[:1]
	}
	for i, rune := range str {
		characterOccurence := i
		character := string(rune)
		//numberOfOccurences := strings.Count(str[characterOccurence+1:], character)
		for true {
			if strings.Contains(str[characterOccurence+1:], character) {
				nthOccurence := characterOccurence + strings.Index(str[characterOccurence+1:], character) + 2
				if i < nthOccurence  {
					characterOccurence = nthOccurence - 1
					possiblePalindrome := str[i:nthOccurence]
					if (len(possiblePalindrome) >  len(longestPalindrome)) && isPanlindrome(str[i:nthOccurence]) {
						fmt.Println("Palindrome", possiblePalindrome)
						longestPalindrome = biggerString(longestPalindrome, possiblePalindrome)
					}

				}

			} else {
				break
			}
		}
	}
	fmt.Println("Answer", longestPalindrome)
	return longestPalindrome
}

func Reverse(text string) (result string) {
	for _, v := range text {
		result = string(v) + result
	}
	return result
}

func biggerString(str1, str2 string) string {
	if len(str2) < len(str1) {
		return str1
	}
	return str2
}

func isPanlindrome(str1 string) bool {
	return str1 == Reverse(str1)
}

func main() {
	longestPalindrome("")
	longestPalindrome("a")
	longestPalindrome("babad")
	longestPalindrome("mwwfjysbkebpdjyabcfkgprtxpwvhglddhmvaprcvrnuxifcrjpdgnktvmggmguiiquibmtviwjsqwtchkqgxqwljouunurcdtoeygdqmijdympcamawnlzsxucbpqtuwkjfqnzvvvigifyvymfhtppqamlgjozvebygkxawcbwtouaankxsjrteeijpuzbsfsjwxejtfrancoekxgfyangvzjkdskhssdjvkvdskjtiybqgsmpxmghvvicmjxqtxdowkjhmlnfcpbtwvtmjhnzntxyfxyinmqzivxkwigkondghzmbioelmepgfttczskvqfejfiibxjcuyevvpawybcvvxtxycrfbcnpvkzryrqujqaqhoagdmofgdcbhvlwgwmsmhomknbanvntspvvhvccedzzngdywuccxrnzbtchisdwsrfdqpcwknwqvalczznilujdrlevncdsyuhnpmheukottewtkuzhookcsvctsqwwdvfjxifpfsqxpmpwospndozcdbfhselfdltmpujlnhfzjcgnbgprvopxklmlgrlbldzpnkhvhkybpgtzipzotrgzkdrqntnuaqyaplcybqyvidwcfcuxinchretgvfaepmgilbrtxgqoddzyjmmupkjqcypdpfhpkhitfegickfszermqhkwmffdizeoprmnlzbjcwfnqyvmhtdekmfhqwaftlyydirjnojbrieutjhymfpflsfemkqsoewbojwluqdckmzixwxufrdpqnwvwpbavosnvjqxqbosctttxvsbmqpnolfmapywtpfaotzmyjwnd")
	longestPalindrome("flsuqzhtcahnyickkgtfnlyzwjuiwqiexthpzvcweqzeqpmqwkydhsfipcdrsjkefehhesubkirhalgnevjugfohwnlhbjfewiunlgmomxkafuuokesvfmcnvseixkkzekuinmcbmttzgsqeqbrtlwyqgiquyylaswlgfflrezaxtjobltcnpjsaslyviviosxorjsfncqirsjpkgajkfpoxxmvsyynbbovieoothpjgncfwcvpkvjcmrcuoronrfjcppbisqbzkgpnycqljpjlgeciaqrnqyxzedzkqpqsszovkgtcgxqgkflpmrikksaupukdvkzbltvefitdegnlmzeirotrfeaueqpzppnsjpspgomyezrlxsqlfcjrkglyvzvqakhtvfmeootbtbwfhqucbnuwznigoyatvkocqmbtqghybwrhmyvvuchjpvjckiryvjfxabezchynfxnpqaeampvaapgmvoylyutymdhvhqfmrlmzkhuhupizqiujpwzarnszrexpvgdmtoxvjygjpmiadzdcxtggwamkbwrkeplesupagievwsaaletcuxtpsxmbmeztcylsjxvhzrqizdmgjfyftpzpgxateopwvynljzffszkzzqgofdlwyknqfruhdkvmvrrjpijcjomnrjjubfccaypkpfokohvkqndptciqqiscvmpozlyyrwobeuazsawtimnawquogrohcrnmexiwvjxgwhmtpykqlcfacuadyhaotmmxevqwarppknoxthsmrrknu")
}
