package reverse

func Reverse(input string) string {
	var output string
	for _, letter := range input {
		output = string(letter) + output
	}
	return output
}
