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
