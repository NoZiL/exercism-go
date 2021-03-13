package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var res []string
	for i := 1; i < len(rhyme); i++ {
		res = append(res, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	if len(rhyme) > 0 {
		res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	return res
}
