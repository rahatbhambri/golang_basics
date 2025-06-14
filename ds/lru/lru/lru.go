package main

import (
	"fmt"
	pkg_ll "sample_go/ds/lru/dll"
)

// Store format: hm -> {k : *node}
// DLL : Node -> [k, v]

// get -> check if key present in hashmap
// 	Yes -> Get the value
// 			Relocate the node to the end of list
// 	No ->  Return -1

// Set -> Set the key with a specific value
// Check if key is already present in the map
// 	Yes -> Replace the value with newer value and relocate the node
//  No -> Check if lru is at capacity
// 		Yes-> Evict LRU key using head of DLL
// 		No -> Set a new key pair in hashmap and appendRight

// Eviction -> remove key from hashmap, Change pointers, move head

func evict(dll *pkg_ll.DLL, hm map[int]*pkg_ll.Node, cap int) {
	if len(hm) > cap {
		dl_node := dll.Head
		delete(hm, dl_node.Val[0])
		dll.Delete(dll.Head)
	}
}

func Set(l []int, dll *pkg_ll.DLL, hm map[int]*pkg_ll.Node) {

	k, v := l[0], l[1]
	if node, pres := hm[k]; pres {
		dll.RelocateNode(node)
		node.Val = []int{k, v}
	} else {
		hm[k] = dll.AppendRight(l)
	}
}

func main() {

	dll := &pkg_ll.DLL{}

	var hm map[int]*pkg_ll.Node = make(map[int]*pkg_ll.Node)
	var cap int = 3
	var k int
	var v int
	for i := 1; i <= 5; i += 1 {
		fmt.Print("Enter key")
		fmt.Scan(&k)

		fmt.Print("Enter value")
		fmt.Scan(&v)

		Set([]int{k, v}, dll, hm)
		evict(dll, hm, cap)
		dll.PrintList()
	}

	fmt.Println(get(4, dll, hm))
	fmt.Println(get(5, dll, hm))
	fmt.Println(get(3, dll, hm))
	fmt.Println(get(2, dll, hm))
	fmt.Println(get(4, dll, hm))
	fmt.Println(get(27, dll, hm))
	fmt.Println(get(3, dll, hm))

}

func get(k int, dll *pkg_ll.DLL, hm map[int]*pkg_ll.Node) int {
	var ans int
	if v, pres := hm[k]; pres {
		fmt.Println(pres)
		ans = v.Val[1]
		dll.RelocateNode(hm[k])
	} else {
		ans = -1
	}
	return ans
}
