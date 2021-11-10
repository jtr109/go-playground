// https://leetcode.com/problems/evaluate-reverse-polish-notation/

package evaluatereversepolishnotation

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, r := range tokens {
		switch r {
		case "+":
			i, j := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, i+j)
		case "-":
			i, j := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, i-j)
		case "*":
			i, j := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, i*j)
		case "/":
			i, j := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, i/j)
		default:
			i, _ := strconv.Atoi(r)
			stack = append(stack, i)
		}
	}
	return stack[0]
}
