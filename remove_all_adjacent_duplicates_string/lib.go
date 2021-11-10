// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

package removealladjacentduplicatesstring

func removeDuplicates(s string) string {
	stack := []rune{}
	for _, r := range s {
		if len(stack) == 0 {
			stack = append(stack, r)
			continue
		}
		last := stack[len(stack)-1]
		if last == r {
			stack = stack[:len(stack)-1] // pop the last element
			continue
		}
		stack = append(stack, r)
	}
	return string(stack)
}
