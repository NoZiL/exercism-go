package isogram

import "unicode"

func IsIsogram(word string) bool {
	var m = map[rune]bool{}
	for _, v := range word {
		char := unicode.ToLower(v)
		switch {
		case char == '-', char == ' ':
		case m[char]:
			return false
		default:
			m[char] = true
		}
	}
	return true
}
