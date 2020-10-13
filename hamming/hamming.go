package hamming

import (
	"errors"
)

// Distance is a function that returns the hamming distance between two given strings
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		err = errors.New("length differs")
		return
	}
	for i, runeA := range a {
		runeB := rune(b[i])
		switch {
		case !isRuneDNA(runeA), !isRuneDNA(runeB):
			err = errors.New("is not DNA")
			return
		case runeA != runeB:
			distance++
		}
	}
	return
}

func isRuneDNA(r rune) (test bool) {
	switch r {
	case 'C', 'A', 'G', 'T':
		test = true
	}
	return
}
