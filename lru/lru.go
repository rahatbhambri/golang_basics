package main

import (
	"fmt"
	"sample_go/dll"
)

var hm map[int]*dll.Node = make(map[int]*dll.Node)
var cap int = 3

func evict() {
	if len(hm) > cap {
		dl_node := dll.Head
		delete(hm, dl_node.Val[0])
		dll.Delete(0)
	}
}

func getAndSet() {
	var k int
	var v int

	fmt.Print("Enter key")
	fmt.Scan(&k)

	fmt.Print("Enter value")
	fmt.Scan(&v)

	v1 := []int{k, v}
	hm[k] = dll.AppendRight(v1)

}

func main() {

	i := 1
	for i <= 5 {
		getAndSet()
		evict()
		i += 1
	}

	fmt.Println(get(4))
	fmt.Println(get(5))
	fmt.Println(get(3))

	getAndSet()
	evict()

	fmt.Println(get(2))
	fmt.Println(get(4))
	fmt.Println(get(27))
	fmt.Println(get(3))

}

func get(k int) int {
	var ans int
	if v, pres := hm[k]; pres {
		ans = v.Val[1]
		dll.RelocateNode(v)
	} else {
		ans = -1
	}
	return ans
}
