package findtheuniquenumber

func FindUniq(arr []float32) float32 {
	var anchor float32
	if arr[0] == arr[1] {
		anchor = arr[0]
	} else if arr[0] == arr[2] {
		return arr[1]
	} else {
		return arr[0]
	}
	for _, n := range arr[2:] {
		if n != anchor {
			return n
		}
	}
	return 0
}
