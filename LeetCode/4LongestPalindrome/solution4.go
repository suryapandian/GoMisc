package main

import (
	"bytes"
	"fmt"
)

func longestPalindrome(st string) (result string) {
	var longestPalindrome []byte
	str := []byte(st)
	s := str
	if len(str) == 0 {
		return
	}
	longestPalindrome = str[:1]
	for outerIndex, character := range str {
		innerIndex := 0
		for {
			occurence := bytes.IndexByte(s, character)
			if occurence == -1 {
				break
			}
			nthOccurence := outerIndex + occurence + innerIndex
			if outerIndex < nthOccurence && len(str) > nthOccurence {
				possiblePalindrome := str[outerIndex : nthOccurence+1]
				if (len(possiblePalindrome) > len(longestPalindrome)) && isPanlindrome(possiblePalindrome) {
					longestPalindrome = biggerString(longestPalindrome, possiblePalindrome)
				}

			}

			s = s[occurence+1:]
			innerIndex = occurence + 1 + innerIndex
		}
		s = str[outerIndex+1:]
	}
	fmt.Println("Answer", string(longestPalindrome))
	return string(longestPalindrome)
}

func biggerString(str1, str2 []byte) []byte {
	if len(str2) < len(str1) {
		return str1
	}
	return str2
}

func isPanlindrome(str []byte) bool {
	n := len(str)
	for i := 0; i <= n/2; i++ {
		if str[i] != str[(n-1)-i] {
			return false
		}
	}
	return true
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
