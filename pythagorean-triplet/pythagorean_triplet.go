package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	var res []Triplet
	for c := min + 2; c <= max; c++ {
		c2 := c * c
		for b := min + 1; b < c; b++ {
			b2 := b * b
			for a := min; a < b; a++ {
				if a*a+b2 == c2 {
					res = append(res, Triplet{a, b, c})
				}
			}
		}
	}
	return res
}

func Sum(sum int) []Triplet {
	var res []Triplet
	for _, v := range Range(1, sum) {
		if v[0]+v[1]+v[2] == sum {
			res = append(res, v)
		}
	}
	return reverse(res)
}

func reverse(res []Triplet) []Triplet {
	for i := len(res)/2 - 1; i >= 0; i-- {
		opp := len(res) - 1 - i
		res[i], res[opp] = res[opp], res[i]
	}
	return res
}
