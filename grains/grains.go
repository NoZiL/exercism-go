package grains

import "errors"

func Square(n int) (uint64, error) {
	switch {
	case n < 1:
		return 0, errors.New("must be positive number")
	case n > 64:
		return 0, errors.New("must be inferior than or equal to 64")
	case n == 1:
		return 1, nil
	}

	t, _ := Square(n - 1)
	return t * 2, nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		sq, _ := Square(i)
		sum += sq
	}
	return sum
}
