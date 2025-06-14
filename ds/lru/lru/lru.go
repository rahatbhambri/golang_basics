package main

import (
	"fmt"
	pkg_ll "sample_go/ds/lru/dll"
)

/*
Store format: hm -> {k : *node}
DLL : Node -> [k, v]

get -> check if key present in hashmap
	Yes -> Get the value
			Relocate the node to the end of list
	No ->  Return -1

Set -> Set the key with a specific value
Check if key is already present in the map
	Yes -> Replace the value with newer value and relocate the node
 No -> Check if lru is at capacity
		Yes-> Evict LRU key using head of DLL
		No -> Set a new key pair in hashmap and appendRight

Eviction -> remove key from hashmap, Change pointers, move head
*/

// interface cache implements get and set
// struct lru_cache which implements this interface
// 		fields -> cap,

type cache interface {
	Set(l []int)
	get(k int)
}

type lru_cache struct {
	cap int
	dll *pkg_ll.DLL
	hm  map[int]*pkg_ll.Node
}

func NewLRU(cap int) *lru_cache {
	return &lru_cache{
		cap: cap,
		dll: &pkg_ll.DLL{},
		hm:  make(map[int]*pkg_ll.Node),
	}
}

func (lru *lru_cache) Set(l []int) {
	k, v := l[0], l[1]
	if node, pres := lru.hm[k]; pres {
		lru.dll.RelocateNode(node)
		node.Val = []int{k, v}
	} else {
		lru.hm[k] = lru.dll.AppendRight(l)
	}

	// Eviction logic
	if len(lru.hm) > lru.cap {
		dl_node := lru.dll.Head
		delete(lru.hm, dl_node.Val[0])
		lru.dll.Delete(lru.dll.Head)
	}
}

func (lru *lru_cache) get(k int) int {
	hm := lru.hm
	var ans int
	if v, pres := hm[k]; pres {
		fmt.Println(pres)
		ans = v.Val[1]
		lru.dll.RelocateNode(hm[k])
	} else {
		ans = -1
	}
	return ans
}

func (lru *lru_cache) PrintCache() {
	lru.dll.PrintList()
}

func main() {
	cache := NewLRU(3)
	var k int
	var v int
	for i := 1; i <= 5; i += 1 {
		fmt.Print("Enter key")
		fmt.Scan(&k)

		fmt.Print("Enter value")
		fmt.Scan(&v)

		cache.Set([]int{k, v})
		cache.PrintCache()
	}

	fmt.Println(cache.get(4))
	fmt.Println(cache.get(1))
	fmt.Println(cache.get(314))
	fmt.Println(cache.get(24))
	fmt.Println(cache.get(9))
	fmt.Println(cache.get(3))
}
