// https://leetcode.com/problems/unique-binary-search-trees/
//
// 这一题要求的是数量，所以可以理解为将 n 个不同数值的节点组合成二叉搜索树，其中包含的数字具体是什么并不重要。
// 所以 n 个节点的排布可以分为几类：根节点只有一个 n-1 个节点的左（或右）子树；n-1 个节点按照各种模式排布在根节点的左右子树上。

package uniquebinarysearchtrees

func newCache(n int) []int {
	result := []int{}
	for i := 0; i <= n; i++ {
		result = append(result, -1)
	}
	return result
}

func numTrees(n int) int {
	cache := newCache(n)
	return dp(n, cache)
}

func dp(n int, cache []int) int {
	if cache[n] != -1 {
		return cache[n]
	}
	if n == 0 || n == 1 {
		cache[n] = 1
		return 1
	}
	result := 0
	for i := 0; i < n; i++ {
		result += dp(i, cache) * dp(n-1-i, cache)
	}
	cache[n] = result
	return result
}
