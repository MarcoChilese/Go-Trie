package trie

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

func ordChar(key string) int {
	value, _ := utf8.DecodeRuneInString(key)
	aVal, _ := utf8.DecodeRuneInString("a")

	return int(value) - int(aVal)
}

func (t Trie) AddWord(key string) {
	wordlen := len(key)
	current := t.root

	for i := 0; i < wordlen; i++ {
		position := ordChar(string(key[i]))

		if current.children[position] == nil {
			current.children[position] = newNode()
		}
		current = current.children[position]
		current.value = string(key[i])
	}
	current.end = true
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func (t Trie) PrefixSearch(prefix string) []string {
	if strings.Contains(prefix, "\n") {
		prefix = strings.ReplaceAll(prefix, "\n", "")
	}

	if isInt(prefix) {
		return nil
	}

	current := t.root

	for _, c := range prefix {
		current = current.children[ordChar(string(c))]
		if current == nil {
			return nil
		}
	}

	var found []string
	stack := []Couple{{a: current, b: prefix}}
	for len(stack) > 0 {
		el := pop(&stack)
		current, prefix := el.a, el.b

		if current.end {
			found = append(found, prefix)
		}

		for _, child := range current.children {
			if child == nil {
				continue
			}
			stack = append(stack, Couple{a: child, b: prefix + child.value})
		}
	}
	return found
}
