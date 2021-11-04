// https://leetcode.com/problems/reverse-words-in-a-string/

package reversewordsinastring

func reverseWords(s string) string {
	runes := []rune(s)
	// distinct spaces
	i := 0
	j := 0
	for j < len(runes) {
		if runes[j] == ' ' {
			j++
			continue
		}
		if j > 0 && runes[j-1] == ' ' && i != 0 {
			runes[i] = ' '
			i++
		}
		runes[i] = runes[j]
		i++
		j++
	}
	runes = runes[:i]
	// reverse rune by rune
	for i, j := 0, len(runes)-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	for left, right := 0, 0; right < len(runes); {
		for right < len(runes) && runes[right] != ' ' {
			right++
		}
		for i, j := left, right-1; i < j; {
			runes[i], runes[j] = runes[j], runes[i]
			i++
			j--
		}
		left = right + 1
		right = left
	}
	return string(runes)
}
