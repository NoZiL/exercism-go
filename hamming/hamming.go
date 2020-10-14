package hamming

import (
	"errors"
)

// Distance is a function that returns the hamming distance between two given strings
func Distance(a, b string) (int, error) {
	var distance int
	if len(a) != len(b) {
		return distance, errors.New("length differs")
	}
	for i, runeA := range a {
		runeB := rune(b[i])
		switch {
		case !isRuneDNA(runeA), !isRuneDNA(runeB):
			return distance, errors.New("is not DNA")
		case runeA != runeB:
			distance++
		}
	}
	return distance, nil
}

func isRuneDNA(r rune) (test bool) {
	switch r {
	case 'C', 'A', 'G', 'T':
		test = true
	}
	return
}
