// https://leetcode.com/problems/minimum-window-substring/

package minimumwindowsubstring

func minWindow(s string, t string) string {
	tRuneCounts := make(map[rune]int) // 变量 `t` 字符计数
	for _, char := range t {
		tRuneCounts[char]++
	}
	sRunes := []rune(s)
	// 为 substring 维护一个 map 来记录 substring 中我们需要关注的字符的数量
	substringRuneCounts := make(map[rune]int)
	minSubstring := ""
	for i, j := 0, 0; j < len(s); j++ {
		substringRuneCounts[sRunes[j]]++
		// move forward the pointer `i`
		for i <= j {
			count, required := tRuneCounts[sRunes[i]]
			if required && substringRuneCounts[sRunes[i]] <= count {
				// if the character is required but the count in substring is
				// not greater than required, the pointer `i` cannot move forward
				break
			}
			substringRuneCounts[sRunes[i]]--
			i++
		}
		substringIsValid := true
		for char, count := range tRuneCounts {
			countInSubstring := substringRuneCounts[char]
			if countInSubstring < count {
				substringIsValid = false
				break
			}
		}
		if !substringIsValid {
			// the requirement of character counts are not sufficient
			continue
		}
		if len(minSubstring) == 0 || j-i+1 < len(minSubstring) {
			minSubstring = s[i : j+1]
		}
	}
	return minSubstring
}

// time complexity: O(N)
// space complexity: O(N)
