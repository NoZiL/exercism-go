package grains

import "errors"

const maxSquares = 64

// Square returns how many grains were on a given square
func Square(n int) (uint64, error) {
	if n < 1 || n > maxSquares {
		return 0, errors.New("n must be number between 1 and 64")
	}
	return 1 << (n - 1), nil
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	return 1<<maxSquares - 1
}
