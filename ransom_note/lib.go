// https://leetcode.com/problems/ransom-note/

package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	chars := []int{}
	for i := 0; i < 26; i++ {
		chars = append(chars, 0)
	}
	for _, r := range magazine {
		chars[int(r)-int('a')]++
	}
	for _, r := range ransomNote {
		chars[int(r)-int('a')]--
	}
	for _, count := range chars {
		if count < 0 {
			return false
		}
	}
	return true
}
