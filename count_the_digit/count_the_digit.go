package countthedigit

import (
	"fmt"
	"strings"
)

func NbDig(n int, d int) int {
	result := 0
	for k := 0; k <= n; k++ {
		result += strings.Count(fmt.Sprint(k*k), fmt.Sprint(d))
	}
	return result
}

func digits(n int) (result []int) {
	if n == 0 {
		result = append(result, 0)
		return
	}
	for ; n > 0; n /= 10 {
		result = append(result, n%10)
	}
	return
}
