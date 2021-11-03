// https://leetcode.com/problems/reverse-string-ii/

package reversestringii

func reverseStr(s string, k int) string {
	runeSlice := []rune(s)
	for i := 0; 2*k*i < len(runeSlice); i++ {
		left := 2 * k * i
		right := left + k - 1
		if right > len(runeSlice)-1 {
			right = len(runeSlice) - 1
		}
		for left < right {
			runeSlice[left], runeSlice[right] = runeSlice[right], runeSlice[left]
			left++
			right--
		}
	}
	return string(runeSlice)
}
