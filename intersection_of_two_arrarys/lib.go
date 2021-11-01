// https://leetcode.com/problems/intersection-of-two-arrays/

package intersectionoftwoarrarys

func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]int)
	set2 := make(map[int]int)
	result := []int{}
	for _, n := range nums1 {
		set1[n] = 0
	}
	for _, n := range nums2 {
		set2[n] = 0
	}
	for k := range set1 {
		if _, ok := set2[k]; ok {
			result = append(result, k)
		}
	}
	return result
}

// time complexity: O(N)
