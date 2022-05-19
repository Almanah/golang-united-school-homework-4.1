package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	result_runes := make([]rune, len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		result_runes = append(result_runes, runes[i])
	}
	output = string(result_runes)
	return output
}
