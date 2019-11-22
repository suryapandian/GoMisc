package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(str string) (longestPalindrome string) {
	fmt.Println("string",str)
	if len(str) > 0 {
		longestPalindrome = str[:1]
	}
	for i, rune := range str {
		characterOccurence := i
		character := string(rune)
		fmt.Println("character",character,str[characterOccurence+1:])
		numberOfOccurences := Count(str[characterOccurence+1:], character)
		fmt.Println("numberOfOccurences", numberOfOccurences)
		for _, w:= range(numberOfOccurences) {
			nthOccurence := characterOccurence + w + 2
			fmt.Println("nthOccurence", nthOccurence,i)
			if i < nthOccurence  {
				characterOccurence = nthOccurence - 1
				possiblePalindrome := str[i:nthOccurence]
				fmt.Println("Palindrome", possiblePalindrome)

				if (len(possiblePalindrome) >  len(longestPalindrome)) && isPanlindrome(str[i:nthOccurence]) {
					fmt.Println("Palindrome", possiblePalindrome)
					longestPalindrome = biggerString(longestPalindrome, possiblePalindrome)
				}

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

func Count(s, substr string) (result []int) {
	fmt.Println("sss",len(substr))
	// special case
	if len(substr) == 0 {
		return 
	}
	// if len(substr) == 1 {
	// 	result = append(result, 0)
	// 	return 
	// }
	for {
		i := strings.Index(s, substr)
		if i == -1 {
			return 
		}
		result = append(result, i)
		s = s[i+len(substr):]
	}
	return
}

func main() {
	longestPalindrome("")
	longestPalindrome("a")
	longestPalindrome("babad")
	longestPalindrome("bb")
	//longestPalindrome("mwwfjysbkebpdjyabcfkgprtxpwvhglddhmvaprcvrnuxifcrjpdgnktvmggmguiiquibmtviwjsqwtchkqgxqwljouunurcdtoeygdqmijdympcamawnlzsxucbpqtuwkjfqnzvvvigifyvymfhtppqamlgjozvebygkxawcbwtouaankxsjrteeijpuzbsfsjwxejtfrancoekxgfyangvzjkdskhssdjvkvdskjtiybqgsmpxmghvvicmjxqtxdowkjhmlnfcpbtwvtmjhnzntxyfxyinmqzivxkwigkondghzmbioelmepgfttczskvqfejfiibxjcuyevvpawybcvvxtxycrfbcnpvkzryrqujqaqhoagdmofgdcbhvlwgwmsmhomknbanvntspvvhvccedzzngdywuccxrnzbtchisdwsrfdqpcwknwqvalczznilujdrlevncdsyuhnpmheukottewtkuzhookcsvctsqwwdvfjxifpfsqxpmpwospndozcdbfhselfdltmpujlnhfzjcgnbgprvopxklmlgrlbldzpnkhvhkybpgtzipzotrgzkdrqntnuaqyaplcybqyvidwcfcuxinchretgvfaepmgilbrtxgqoddzyjmmupkjqcypdpfhpkhitfegickfszermqhkwmffdizeoprmnlzbjcwfnqyvmhtdekmfhqwaftlyydirjnojbrieutjhymfpflsfemkqsoewbojwluqdckmzixwxufrdpqnwvwpbavosnvjqxqbosctttxvsbmqpnolfmapywtpfaotzmyjwnd")
	//longestPalindrome("flsuqzhtcahnyickkgtfnlyzwjuiwqiexthpzvcweqzeqpmqwkydhsfipcdrsjkefehhesubkirhalgnevjugfohwnlhbjfewiunlgmomxkafuuokesvfmcnvseixkkzekuinmcbmttzgsqeqbrtlwyqgiquyylaswlgfflrezaxtjobltcnpjsaslyviviosxorjsfncqirsjpkgajkfpoxxmvsyynbbovieoothpjgncfwcvpkvjcmrcuoronrfjcppbisqbzkgpnycqljpjlgeciaqrnqyxzedzkqpqsszovkgtcgxqgkflpmrikksaupukdvkzbltvefitdegnlmzeirotrfeaueqpzppnsjpspgomyezrlxsqlfcjrkglyvzvqakhtvfmeootbtbwfhqucbnuwznigoyatvkocqmbtqghybwrhmyvvuchjpvjckiryvjfxabezchynfxnpqaeampvaapgmvoylyutymdhvhqfmrlmzkhuhupizqiujpwzarnszrexpvgdmtoxvjygjpmiadzdcxtggwamkbwrkeplesupagievwsaaletcuxtpsxmbmeztcylsjxvhzrqizdmgjfyftpzpgxateopwvynljzffszkzzqgofdlwyknqfruhdkvmvrrjpijcjomnrjjubfccaypkpfokohvkqndptciqqiscvmpozlyyrwobeuazsawtimnawquogrohcrnmexiwvjxgwhmtpykqlcfacuadyhaotmmxevqwarppknoxthsmrrknu")
}
