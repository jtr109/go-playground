package nthtribonaccinumber

func tribonacci(n int) int {
	a := 0
	b := 1
	c := 1
	for i := 0; i < n; i++ {
		a, b, c = b, c, a+b+c
	}
	return a
}
