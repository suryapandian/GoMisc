
package grandma

import "strings"
import "fmt"


func isYelling(str string) bool {
	strUpper := strings.ToUpper(str)
	isUpperCase := strUpper == str
	hasLetters := strings.ContainsAny(strUpper,"ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return (isUpperCase && hasLetters)
}

func isQuestion(str string) bool {
	return strings.HasSuffix(str, "?")
}

func Hey(remark string) string {
	fmt.Println("Remark to Grandma:", remark)
	remark = strings.TrimSpace(remark)
	response := "Whatever."
	if isQuestion(remark){
		if isYelling(remark){
			response = "Calm down, I know what I'm doing!"
		}else{
	 		response = "Sure."
		}
	}else{
		if isYelling(remark){
			response = "Whoa, chill out!"
		}
		if len(remark) == 0 {
			response ="Fine. Be that way!"
		}
	}
	fmt.Println("Response of Grandma:", response)
	return response
}
