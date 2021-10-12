package guessnumberhigherorlower

func guess(num int) int {
	return 1
}

func guessNumber(n int) int {
	return guessInRange(1, n+1)
}

func guessInRange(start int, end int) int {
	mid := (start + end) / 2
	switch guess(mid) {
	case -1:
		// The number I picked is lower than your guess
		return guessInRange(start, mid)
	case 1:
		// The number I picked is higher than your guess
		return guessInRange(mid, end)
	default:
		return mid
	}
}
