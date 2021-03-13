package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func KindFromSides(a, b, c float64) Kind {
	switch {
	case a <= 0, math.IsNaN(a), math.IsInf(a, 1),
		b <= 0, math.IsNaN(b), math.IsInf(b, 1),
		c <= 0, math.IsNaN(c), math.IsInf(c, 1),
		a > b+c, b > a+c, c > a+b:
		return NaT
	case a == b && a == c:
		return Equ
	case a == b, a == c, b == c:
		return Iso
	default:
		return Sca
	}
}
