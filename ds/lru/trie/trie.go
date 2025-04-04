package main

import "fmt"

type node struct {
	endsHere bool
	next     map[string]*node
}

var root = &node{}

func main() {
	words := []string{"batman", "bastard", "blacki", "banger", "alladin", "acrobatic"}

	for _, w := range words {
		insert(w)
	}
	// PrintTrie("", root)

	getWord("ba")

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

// Get all words in the trie
func PrintTrie(curr string, node *node) {
	var temp = node

	if node.endsHere {
		fmt.Println(curr)
	}
	for k, v := range temp.next {
		PrintTrie(curr+k, v)
	}
}

// Get all words starting with a particular prefix
func getWord(prefix string) []string {
	var temp = root
	var ans []string

	for _, c := range prefix {
		ch := string(c)

		if temp.next == nil {
			return ans
		}
		_, pres := temp.next[ch]
		if !pres {
			return ans
		}
		temp = temp.next[ch]
	}

	PrintTrie(prefix, temp)
	return ans
}
