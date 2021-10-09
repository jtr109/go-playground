package trie

type Tree struct {
	children *[]*Tree
	end      bool
}

type Trie struct {
	data *Tree
}

func index(c rune) int {
	return int(c) - int('a')
}

func initializeChildren() []*Tree {
	children := []*Tree{}
	for i := 0; i < 26; i++ {
		children = append(children, nil)
	}
	return children
}

func initializeTree() Tree {
	children := initializeChildren()
	return Tree{
		children: &children,
	}
}

func Constructor() Trie {
	children := initializeChildren()
	return Trie{
		data: &Tree{
			children: &children,
		},
	}
}

func (this *Trie) Insert(word string) {
	current := this.data
	for _, c := range word {
		if (*current.children)[index(c)] == nil {
			tree := initializeTree()
			(*current.children)[index(c)] = &tree
		}
		current = (*current.children)[index(c)]
	}
	current.end = true
}

func (this *Trie) Search(word string) bool {
	current := this.data
	for _, c := range word {
		if (*current.children)[index(c)] == nil {
			return false
		}
		current = (*current.children)[index(c)]
	}
	return current.end
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.data
	for _, c := range prefix {
		if (*current.children)[index(c)] == nil {
			return false
		}
		current = (*current.children)[index(c)]
	}
	return true
}
