package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		for _, v := range divisors {
			if v != 0 && i%v == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
