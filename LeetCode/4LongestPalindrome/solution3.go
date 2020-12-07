package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(str string) (longestPalindrome string) {
	s := str
	if len(str) == 0 {
		return
	}
	longestPalindrome = str[:1]
	for outerIndex, rune := range str {
		character := string(rune)
		innerIndex := 0
		for {
			occurence := strings.Index(s, character)
			if occurence == -1 {
				break
			}
			nthOccurence := outerIndex + occurence + innerIndex
			//fmt.Printf ("nthOccurence %v character %v string %v outerIndex %v Occurence %v\n", nthOccurence,character,s,outerIndex,occurence)
			if outerIndex < nthOccurence && len(str) > nthOccurence {
				possiblePalindrome := str[outerIndex : nthOccurence+1]
				//fmt.Println("possiblePalindrome", possiblePalindrome)
				if (len(possiblePalindrome) > len(longestPalindrome)) && isPanlindrome(possiblePalindrome) {
					longestPalindrome = biggerString(longestPalindrome, possiblePalindrome)
				}

			}

			s = s[occurence+1:]
			innerIndex = occurence + 1 + innerIndex
		}
		s = str[outerIndex+1:]
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
	longestPalindrome("ccc")
	longestPalindrome("bb")
	longestPalindrome("")
	longestPalindrome("a")
	longestPalindrome("babad")
	longestPalindrome("mwwfjysbkebpdjyabcfkgprtxpwvhglddhmvaprcvrnuxifcrjpdgnktvmggmguiiquibmtviwjsqwtchkqgxqwljouunurcdtoeygdqmijdympcamawnlzsxucbpqtuwkjfqnzvvvigifyvymfhtppqamlgjozvebygkxawcbwtouaankxsjrteeijpuzbsfsjwxejtfrancoekxgfyangvzjkdskhssdjvkvdskjtiybqgsmpxmghvvicmjxqtxdowkjhmlnfcpbtwvtmjhnzntxyfxyinmqzivxkwigkondghzmbioelmepgfttczskvqfejfiibxjcuyevvpawybcvvxtxycrfbcnpvkzryrqujqaqhoagdmofgdcbhvlwgwmsmhomknbanvntspvvhvccedzzngdywuccxrnzbtchisdwsrfdqpcwknwqvalczznilujdrlevncdsyuhnpmheukottewtkuzhookcsvctsqwwdvfjxifpfsqxpmpwospndozcdbfhselfdltmpujlnhfzjcgnbgprvopxklmlgrlbldzpnkhvhkybpgtzipzotrgzkdrqntnuaqyaplcybqyvidwcfcuxinchretgvfaepmgilbrtxgqoddzyjmmupkjqcypdpfhpkhitfegickfszermqhkwmffdizeoprmnlzbjcwfnqyvmhtdekmfhqwaftlyydirjnojbrieutjhymfpflsfemkqsoewbojwluqdckmzixwxufrdpqnwvwpbavosnvjqxqbosctttxvsbmqpnolfmapywtpfaotzmyjwnd")
	longestPalindrome("flsuqzhtcahnyickkgtfnlyzwjuiwqiexthpzvcweqzeqpmqwkydhsfipcdrsjkefehhesubkirhalgnevjugfohwnlhbjfewiunlgmomxkafuuokesvfmcnvseixkkzekuinmcbmttzgsqeqbrtlwyqgiquyylaswlgfflrezaxtjobltcnpjsaslyviviosxorjsfncqirsjpkgajkfpoxxmvsyynbbovieoothpjgncfwcvpkvjcmrcuoronrfjcppbisqbzkgpnycqljpjlgeciaqrnqyxzedzkqpqsszovkgtcgxqgkflpmrikksaupukdvkzbltvefitdegnlmzeirotrfeaueqpzppnsjpspgomyezrlxsqlfcjrkglyvzvqakhtvfmeootbtbwfhqucbnuwznigoyatvkocqmbtqghybwrhmyvvuchjpvjckiryvjfxabezchynfxnpqaeampvaapgmvoylyutymdhvhqfmrlmzkhuhupizqiujpwzarnszrexpvgdmtoxvjygjpmiadzdcxtggwamkbwrkeplesupagievwsaaletcuxtpsxmbmeztcylsjxvhzrqizdmgjfyftpzpgxateopwvynljzffszkzzqgofdlwyknqfruhdkvmvrrjpijcjomnrjjubfccaypkpfokohvkqndptciqqiscvmpozlyyrwobeuazsawtimnawquogrohcrnmexiwvjxgwhmtpykqlcfacuadyhaotmmxevqwarppknoxthsmrrknu")

}
