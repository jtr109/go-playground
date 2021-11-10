// https://leetcode.com/problems/valid-parentheses/

package validparentheses

func isValid(s string) bool {
	stack := []rune{}
	for _, r := range s {
		switch r {
		case '(':
			stack = append(stack, r)
		case '[':
			stack = append(stack, r)
		case '{':
			stack = append(stack, r)
		case ')':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if last != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if last != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if last != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}
