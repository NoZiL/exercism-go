package pangram

import "unicode"

func IsPangram(word string) bool {
	alpha := map[rune]bool{}
	for _, v := range word {
		if unicode.IsLetter(v) {
			alpha[unicode.ToLower(v)] = true
		}
	}
	return len(alpha) == 26
}
