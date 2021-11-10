package validparentheses

func ValidParentheses(parens string) bool {
	mark := 0
	for _, p := range parens {
		switch p {
		case '(':
			mark++
		default:
			mark--
		}
		if mark < 0 {
			return false
		}
	}
	return mark == 0
}

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
