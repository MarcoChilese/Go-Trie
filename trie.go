package trie

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
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

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func (t Trie) PrefixSearch(prefix string) []string {
	if strings.Contains(prefix, "\n") {
		prefix = strings.ReplaceAll(prefix, "\n", "")
	}

	if !IsLetter(prefix) {
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

func BuildTrieFromDictionary(pathToDict string) *Trie {
	var gotrie = NewTrie()

	file, err := os.Open(pathToDict)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		gotrie.AddWord(word)
	}
	return gotrie
}
