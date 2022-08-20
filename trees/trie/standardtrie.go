package trie

const AlphabetSize = 26

type Node struct {
	isEndOfWord bool
	children    []*Node
}

func NewNode() *Node {
	n := &Node{
		isEndOfWord: false,
		children:    make([]*Node, AlphabetSize),
	}
	return n
}

type StandardTrie struct {
	root *Node
}

func NewStandardTrie() *StandardTrie {
	return &StandardTrie{
		root: NewNode(),
	}
}

func (st *StandardTrie) Insert(key string) {
	depth := 0
	ttlChars := len(key)
	index := 0
	crawler := st.root

	for depth = 0; depth < ttlChars; depth++ {
		index = int(key[depth]) - 'a'
		if crawler.children[index] == nil {
			crawler.children[index] = NewNode()
		}
		crawler = crawler.children[index]
	}

	crawler.isEndOfWord = true
}

func (st *StandardTrie) Search(key string) bool {
	depth := 0
	ttlChars := len(key)
	index := 0
	parent := st.root

	for depth = 0; depth < ttlChars; depth++ {
		index = int(key[depth]) - 'a'
		if parent.children[index] == nil {
			return false
		}
		parent = parent.children[index]
	}
	return parent.isEndOfWord
}
