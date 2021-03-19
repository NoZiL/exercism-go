package reverse

func Reverse(str string) string {
	var res = []rune(str)
	for x, y := 0, len(res)-1; x < y; x, y = x+1, y-1 {
		res[x], res[y] = res[y], res[x]
	}
	return string(res)
}
