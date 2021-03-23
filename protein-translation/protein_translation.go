package protein

import "errors"

var ErrStop error = errors.New("ErrStop")
var ErrInvalidBase error = errors.New("ErrInvalidBase")

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(rna string) ([]string, error) {
	if rna == "" {
		return []string{}, nil
	}

	protein, err := FromCodon(rna[:3])
	if err != nil {
		return nil, err
	}

	proteins, err := FromRNA(rna[3:])
	if err == ErrStop {
		err = nil
	}

	return append([]string{protein}, proteins...), err
}
