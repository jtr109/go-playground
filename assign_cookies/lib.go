package assigncookies

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	i := len(g) - 1
	for j := len(s) - 1; j >= 0; j-- {
		for i >= 0 && g[i] > s[j] {
			i--
		}
		if i < 0 {
			break
		}
		count++
		i--
	}
	return count
}
