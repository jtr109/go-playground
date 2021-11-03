package tihuankonggelcof

func replaceSpace(s *[]rune) {
	left, right := len(*s)-1, len(*s)-1
	for _, b := range *s {
		if b == ' ' {
			for i := 0; i < 2; i++ {
				*s = append(*s, 0)
				right++
			}
		}
	}
	for left >= 0 {
		if (*s)[left] == ' ' {
			(*s)[right] = '0'
			right--
			(*s)[right] = '2'
			right--
			(*s)[right] = '%'
		} else {
			(*s)[right] = (*s)[left]
		}
		left--
		right--
	}
}
