package uniquepaths

type cache struct {
	content map[int]map[int]int
}

func newCache() cache {
	return cache{content: make(map[int]map[int]int)}
}

func (c *cache) set(m, n, count int) {
	v, ok := c.content[m]
	if !ok {
		c.content[m] = make(map[int]int)
		v = c.content[m]
	}
	v[n] = count
}

func (c *cache) get(m, n int) (int, bool) {
	v, ok := c.content[m]
	if !ok {
		return 0, false
	}
	count, ok := v[n]
	if !ok {
		return 0, false
	}
	return count, true
}

func uniquePaths(m int, n int) int {
	c := newCache()
	return dp(m, n, &c)
}

func dp(m, n int, c *cache) int {
	if count, ok := c.get(m, n); ok {
		return count
	}
	if m == 1 || n == 1 {
		c.set(m, n, 1)
		return 1
	}
	result := dp(m-1, n, c) + dp(m, n-1, c)
	c.set(m, n, result)
	return result
}
