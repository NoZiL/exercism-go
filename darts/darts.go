package darts

import "math"

type circle struct {
	radius float64
	score  int
}

var (
	inner  = circle{1, 10}
	middle = circle{5, 5}
	outer  = circle{10, 1}
)

func Score(x, y float64) int {
	switch r := radius(x, y); {
	case r <= inner.radius:
		return inner.score
	case r <= middle.radius:
		return middle.score
	case r <= outer.radius:
		return outer.score
	default:
		return 0
	}
}

func radius(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
