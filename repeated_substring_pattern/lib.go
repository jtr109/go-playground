// https://leetcode.com/problems/repeated-substring-pattern/

package repeatedsubstringpattern

func buildNext(pattern string) (next []int) {
	runes := []rune(pattern)
	j := -1
	next = append(next, j)
	for i := 1; i < len(runes); i++ {
		for j >= 0 && runes[j+1] != runes[i] {
			j = next[j]
		}
		if runes[j+1] == runes[i] {
			j++
		}
		next = append(next, j)
	}
	return
}

func repeatedSubstringPattern(s string) bool {
	next := buildNext(s)
	return next[len(s)-1] != -1 && len(s)%(len(s)-(next[len(s)-1]+1)) == 0
}
