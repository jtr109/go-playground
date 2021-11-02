// https://leetcode.com/problems/4sum-ii/

package foursumii

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// return occurrence of the sum of elements in num1 and num2
	count := 0
	cache := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			cache[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			count += cache[-n3-n4]
		}
	}
	return count
}
