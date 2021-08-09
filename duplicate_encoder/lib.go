// https://www.codewars.com/kata/54b42f9314d9229fd6000d9c/train/go

package duplicateencoder

import "strings"

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	cache := make(map[rune]int)
	res := []rune{}
	for _, r := range word {
		cache[r]++
	}
	for _, r := range word {
		if cache[r] > 1 {
			res = append(res, ')')
		} else {
			res = append(res, '(')
		}
	}
	return string(res)
}
