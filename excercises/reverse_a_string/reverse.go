package reverse

func Reverse(text string) (result string){
	for _, v:= range text {
		result =  string(v) + result
	}
	return result
}