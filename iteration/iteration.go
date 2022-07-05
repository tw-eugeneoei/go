package iteration

import "strings"

// * In Go there are no "while", "do", "until" keywords
func UseRepeat(character string, count int) string {
	// result := ""
	// for i := 0; i < count; i++ {
	// 	result += character
	// }
	// return result

	return strings.Repeat(character, count)
}
