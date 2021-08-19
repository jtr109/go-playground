package maxpq

type MaxPQ []int

func (pq MaxPQ) swim(i int) {
	for i > 1 {
		if pq[i] > pq[i/2] {
			pq[i], pq[i/2] = pq[i/2], pq[i]
			i /= 2
		} else {
			break
		}
	}
}

func (pq MaxPQ) sink(i int) {
	for 2*i < len(pq) {
		var j int
		if 2*i+1 >= len(pq) {
			j = 2 * i
		} else if pq[2*i] > pq[2*i+1] {
			j = 2 * i
		} else {
			j = 2*i + 1
		}
		if pq[i] > pq[j] {
			break
		}
		pq[i], pq[j] = pq[j], pq[i]
		i = j
	}
}
