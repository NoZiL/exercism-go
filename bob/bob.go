package bob

import "strings"

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
	upper := strings.ToUpper(str)
	return str == upper && upper != strings.ToLower(str)
}
