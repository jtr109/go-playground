package sqrtx

func mySqrt(x int) int {
	// find the largest number whose square is less than the number x
	gte := 0
	lte := x
	guess := x / 2
	for gte <= lte {
		mid := gte + (lte-gte)/2
		diff := mid*mid - x
		if diff == 0 {
			// gotha
			return mid
		} else if diff > 0 {
			// the result is between the least number and the mid one
			lte = mid - 1
		} else if diff < 0 {
			// the mid may be equal to or less than the result
			guess = mid
			gte = mid + 1
		}
	}
	return guess
}
