package wfswordchange

// string array
// begin word
// end word
// -> hit
// hot .
// dot .
// dog .
// lot
// log
// -> cog
// output: number steps
// if not exists return 0

func steps(begin, end string, arr []string) int {
	return 1
}

type Word struct {
	content string
	related []*Word
	marked  bool
}

func createGraph(arr []string) []*Word {
	graph := []*Word{}
	for _, s := range arr {
		graph = append(graph, &Word{
			content: s,
			related: []*Word{},
		})
	}
	return graph
}

func buildGraph(graph []*Word) {
	for i := 0; i < len(graph); i++ {
		for j := i + 1; j < len(graph); j++ {
			w1 := graph[i]
			w2 := graph[j]
			if diffOneChar(w1, w2) {
				w1.related = append(w1.related, w2)
				w2.related = append(w2.related, w1)
			}
		}
	}
}

func diffOneChar(w1, w2 *Word) bool {
	s1 := w1.content
	s2 := w2.content
	if len(s1) != len(s2) {
		return false
	}
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff == 1
}

func search(begin, end *Word) int {
	step := 0
	words := []*Word{begin}
	for len(words) != 0 { // words empty
		step++
		nextWords := []*Word{} // TODO: refine
		for _, word := range words {
			for _, related := range word.related {
				if related.marked { // word is marked
					continue
				}
				if related == end {
					return step
				}
				related.marked = true
				nextWords = append(nextWords, related)
			}
		}
		words = nextWords
	}
	return 0
}
