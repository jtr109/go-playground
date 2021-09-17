package intersectionoftwoarraysii

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	nums1Map := make(map[int]int)
	for _, n := range nums1 {
		nums1Map[n] += 1
	}
	for _, n := range nums2 {
		if nums1Map[n] == 0 {
			continue
		}
		result = append(result, n)
		nums1Map[n] -= 1
	}
	return result
}
