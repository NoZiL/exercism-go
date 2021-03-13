package strand

import "strings"

func ToRNA(dna string) string {
	var res string
	for _, nucleotide := range strings.Split(dna, "") {
		switch nucleotide {
		case "G":
			res += "C"
		case "C":
			res += "G"
		case "T":
			res += "A"
		case "A":
			res += "U"
		}
	}
	return res
}
