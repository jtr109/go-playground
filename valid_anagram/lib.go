package validanagram

func isAnagram(s string, t string) bool {
	sCounts := getCharCounts(s)
	tCounts := getCharCounts(t)
	for i := 0; i < 26; i++ {
		if sCounts[i] != tCounts[i] {
			return false
		}
	}
	return true
}

func getCharCounts(s string) []int {
	counts := []int{}
	for i := 0; i < 26; i++ {
		counts = append(counts, 0)
	}
	for _, r := range s {
		counts[int(r)-int('a')]++
	}
	return counts
}
