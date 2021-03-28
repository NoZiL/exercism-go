package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.Trim(remark, "\n\r \t")
	shouting := isShouting(remark)
	asking := strings.HasSuffix(remark, "?")
	switch {
	case shouting && asking:
		return "Calm down, I know what I'm doing!"
	case shouting:
		return "Whoa, chill out!"
	case asking:
		return "Sure."
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

func isShouting(str string) bool {
	hasUpper := false
	for _, v := range str {
		switch {
		case unicode.IsLower(v):
			return false
		case unicode.IsUpper(v):
			hasUpper = true
		}
	}
	return hasUpper
}
