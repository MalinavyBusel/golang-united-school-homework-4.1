package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	symbs := strings.Split(input, "")

	for i, j := 0, len(symbs)-1; i < j; i, j = i+1, j-1 {
		symbs[i], symbs[j] = symbs[j], symbs[i]
	}

	output = strings.Join(symbs, "")
	return output
}
