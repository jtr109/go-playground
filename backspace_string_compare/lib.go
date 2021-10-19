// https://leetcode.com/problems/backspace-string-compare/

package backspacestringcompare

func backspaceCompare(s string, t string) bool {
	return processString(s) == processString(t)
}

func processString(s string) string {
	// 维护两个索引，一个索引记录当前光标的位置，一个索引记录最新的操作位置
	r := []rune(s)
	slowIndex := 0
	fastIndex := 0
	for fastIndex < len(r) {
		if r[fastIndex] != '#' {
			r[slowIndex] = r[fastIndex]
			slowIndex++
		} else if slowIndex != 0 {
			slowIndex--
		}
		fastIndex++
	}
	return string(r[:slowIndex])
}
