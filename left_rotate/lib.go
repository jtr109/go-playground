package leftrotate

func leftRotate(s string, k int) string {
	runes := []rune(s)
	for i, j := 0, k-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	for i, j := k, len(runes)-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	for i, j := 0, len(runes)-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}
