// https://leetcode.com/problems/happy-number/

package happynumber

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	numbers := make(map[int]int)
	for {
		digitSum := 0
		for i := n; i > 0; i /= 10 {
			digitSum += (i % 10) * (i % 10)
		}
		if _, ok := numbers[digitSum]; ok {
			break
		}
		numbers[digitSum] = 0
		n = digitSum
	}
	return n%10 == 1
}
