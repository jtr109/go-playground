// https://www.codewars.com/kata/5277c8a221e209d3f6000b56/go

package validbraces

import "strings"

func ValidBraces(str string) bool {
	cache := []rune{}
	for _, b := range str {
		if strings.ContainsRune("([{", b) {
			cache = append(cache, b)
			continue
		}
		if len(cache) == 0 {
			return false
		}
		lastRune := cache[len(cache)-1]
		if matchPairs(lastRune, b) {
			cache = cache[:len(cache)-1]
		} else {
			return false
		}
	}
	return len(cache) == 0
}

func matchPairs(left rune, right rune) bool {
	switch right {
	case ')':
		return left == '('
	case ']':
		return left == '['
	case '}':
		return left == '{'
	default:
		return false
	}
}
