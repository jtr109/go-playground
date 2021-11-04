// https://programmercarl.com/%E5%89%91%E6%8C%87Offer58-II.%E5%B7%A6%E6%97%8B%E8%BD%AC%E5%AD%97%E7%AC%A6%E4%B8%B2.html

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
