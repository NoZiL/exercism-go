package darts

import "math"

type circle struct {
	radius float64
	score  int
}

var circles = []circle{
	{1, 10},
	{5, 5},
	{10, 1},
}

func Score(x, y float64) int {
	r := radius(x, y)
	for _, c := range circles {
		if r <= c.radius {
			return c.score
		}
	}
	return 0
}

func radius(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
