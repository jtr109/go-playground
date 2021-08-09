package wiredstringcase

import "strings"

func toWeirdCase(str string) string {
	res := []string{}
	for _, s := range strings.Split(str, " ") {
		res = append(res, wordToWeirdCase(s))
	}
	return strings.Join(res, " ")
}

func wordToWeirdCase(str string) string {
	res := []rune{}
	for i, c := range str {
		if i%2 == 0 && c >= 'a' && c <= 'z' {
			res = append(res, c-'a'+'A')
		} else if i%2 == 1 && c >= 'A' && c <= 'Z' {
			res = append(res, c-'A'+'a')
		} else {
			res = append(res, c)
		}
	}
	return string(res)
}
