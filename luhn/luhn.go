package luhn

import (
	"strconv"
	"strings"
)

func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	length := len(number)
	if length <= 1 {
		return false
	}
	var sum int
	for i := 0; i < length; i++ {
		n, err := strconv.Atoi(string(number[length-1-i]))
		if err != nil {
			return false
		}
		if i%2 == 1 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}
	return sum%10 == 0
}
