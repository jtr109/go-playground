// https://leetcode.com/problems/valid-perfect-square/

package validperfectsquare

func isPerfectSquare(num int) bool {
	gte := 0
	lte := num
	for gte <= lte {
		mid := gte + (lte-gte)/2
		diff := mid*mid - num
		if diff == 0 {
			return true
		} else if diff > 0 {
			lte = mid - 1
		} else {
			gte = mid + 1
		}
	}
	return false
}
