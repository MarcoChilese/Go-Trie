package trie

type Node struct {
	children []*Node
	end      bool
	value    string
}

func newNode() *Node {
	return &Node{children: make([]*Node, 26), end: false, value: ""}
}
