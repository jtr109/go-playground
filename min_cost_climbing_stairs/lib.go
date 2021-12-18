// https://leetcode.com/problems/min-cost-climbing-stairs/

package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	minCost := []int{0, 0}
	for i := 2; i < len(cost); i++ {
		c1 := minCost[i-2] + cost[i-2]
		c2 := minCost[i-1] + cost[i-1]
		if c1 < c2 {
			minCost = append(minCost, c1)
		} else {
			minCost = append(minCost, c2)
		}
	}
	c1 := minCost[len(cost)-2] + cost[len(cost)-2]
	c2 := minCost[len(cost)-1] + cost[len(cost)-1]
	if c1 < c2 {
		return c1
	} else {
		return c2
	}
}
