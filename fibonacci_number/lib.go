package fibonaccinumber

func fib(n int) int {
	cache := []int{}
	cache = append(cache, 0, 1)
	for i := 2; i <= n; i++ {
		cache = append(cache, cache[i-2]+cache[i-1])
	}
	return cache[n]
}
