package trie

type Couple struct {
	a *Node
	b string
}

func pop(slice *[]Couple) Couple {
	length := len(*slice)
	popped := (*slice)[length-1]
	*slice = append((*slice)[:length-1])
	return popped
}
