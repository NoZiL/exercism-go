package hamming

import (
	"errors"
)

// Distance is a function that returns the hamming distance between two given strings
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return distance, errors.New("length differs")
	}
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return
}
