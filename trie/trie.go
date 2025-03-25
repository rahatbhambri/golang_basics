package trie

import "fmt"

type node struct {
	endsHere bool
	next     map[string]*node
}

var root = &node{}

func main() {
	words := []string{"abra", "cadabra", "alakazam"}

	for _, w := range words {
		insert(w)
	}
	PrintTrie(root)
}

func insert(word string) {
	var temp = root
	for _, c := range word {
		ch := string(c)

		if temp.next == nil {
			temp.next = make(map[string]*node)
		}
		_, pres := temp.next[ch]
		if !pres {
			temp.next[ch] = &node{}
		}
		temp = temp.next[ch]
	}
	temp.endsHere = true
}

func PrintTrie(node *node) {
	var temp = node
	for k, v := range temp.next {
		fmt.Println(k)
		PrintTrie(v)
	}
}
