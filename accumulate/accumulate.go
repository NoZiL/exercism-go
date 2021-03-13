package accumulate

func Accumulate(array []string, transform func(string) string) []string {
	var res []string
	for _, elem := range array {
		res = append(res, transform(elem))
	}
	return res
}
