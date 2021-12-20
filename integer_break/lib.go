// https://leetcode.com/problems/integer-break/

package integerbreak

func newCache(n int) []int {
	result := []int{}
	for i := 0; i <= n; i++ {
		result = append(result, -1)
	}
	return result
}

func integerBreak(n int) int {
	cache := newCache(n)
	return dp(n, cache)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func dp(n int, cache []int) int {
	if cache[n] != -1 {
		return cache[n]
	}
	result := 1
	if n == 1 {
		cache[n] = result
		return result
	}
	for i := 1; i <= n/2; i++ {
		res := max(dp(i, cache), i) * max(dp(n-i, cache), n-i)
		if res > result {
			result = res
		}
	}
	cache[n] = result
	return result
}
