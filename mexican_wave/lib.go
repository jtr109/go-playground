package mexicanwave

import (
	"fmt"
	"strings"
)

func wave(words string) []string {
	res := []string{}
	for i, r := range words {
		if r == ' ' {
			continue
		}
		newWords := fmt.Sprintf("%s%s%s", words[:i], strings.ToUpper(string(r)), words[i+1:])
		res = append(res, newWords)
	}
	return res
}
