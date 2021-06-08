package grains

import "errors"

const maxSquares = 64

// Square returns how many grains were on a given square
func Square(n int) (uint64, error) {
	switch {
	case n < 1:
		return 0, errors.New("must be positive number")
	case n > maxSquares:
		return 0, errors.New("must be inferior than or equal to 64")
	}
	acc := uint64(1)
	for i := 1; i < n; i++ {
		acc *= 2
	}
	return acc, nil
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	acc := uint64(1)
	sum := uint64(1)
	for i := 1; i <= maxSquares; i++ {
		acc *= 2
		sum += acc
	}
	return sum
}
