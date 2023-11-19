package main

type Node struct {
	child map[rune]*Node
	last  bool
}

func NewNode() *Node {
	return &Node{
		child: make(map[rune]*Node),
	}
}

type Trie struct {
	rootNode *Node
}

func Constructor() Trie {
	return Trie{
		rootNode: NewNode(),
	}
}

func (this *Trie) Insert(word string) {
	currentNode := this.rootNode
	for _, c := range word {
		if _, ok := currentNode.child[c]; !ok {
			currentNode.child[c] = NewNode()
		}
		currentNode = currentNode.child[c]
	}
	currentNode.last = true
}

func (this *Trie) Search(word string) bool {
	currentNode := this.rootNode
	for _, c := range word {
		if _, ok := currentNode.child[c]; !ok {
			return false
		}
		currentNode = currentNode.child[c]
	}

	return currentNode.last
}

func (this *Trie) StartsWith(prefix string) bool {
	currentNode := this.rootNode
	for _, c := range prefix {
		if _, ok := currentNode.child[c]; !ok {
			return false
		}
		currentNode = currentNode.child[c]
	}

	return true
}

func main() {
	t := Constructor()
	t.Insert("apple")
	t.Search("apple")
	t.StartsWith("app")
}
