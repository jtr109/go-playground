// reverse integer

package reverseinteger

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func reverse(a int) int {
	current := a
	result := 0
	for current != 0 {
		digit := current % 10 // last digit
		current /= 10
		if a > 0 && (MaxInt-digit)/10 < result { // positive overflow
			return 0
		}
		if a < 0 && (MinInt-digit)/10 > result { // negative overflow
			return 0
		}
		result = result*10 + digit
	}
	return result
}
