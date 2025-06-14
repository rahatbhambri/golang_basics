package dll

import (
	"fmt"
	"log"
)

type Node struct {
	Val  []int
	next *Node
	prev *Node
}

type DLL struct {
	Head *Node
	Tail *Node
}

func (dll *DLL) AppendRight(Val []int) *Node {
	Head, Tail := dll.Head, dll.Tail
	if Head == nil {
		Head = &Node{Val, nil, nil}
		Tail = Head
	} else {
		nn := &Node{Val, nil, Tail}
		Tail.next = nn
		Tail = nn
	}

	dll.Head, dll.Tail = Head, Tail
	return Tail
}

func (dll *DLL) AppendLeft(Val []int) {
	Head, Tail := dll.Head, dll.Tail

	log.Println(Tail)
	// empty list
	if Head == nil {
		Head = &Node{Val, nil, nil}
		Tail = Head
	} else {
		nn := &Node{Val, Head, nil}
		Head.prev = nn
		Head = nn
	}
	dll.Head, dll.Tail = Head, Tail
}

func (dll *DLL) PrintList() {

	temp := dll.Head
	for {
		if temp != nil {
			fmt.Println(temp.Val, " ")
			temp = temp.next
		} else {
			break
		}
	}
	println()
}

func (dll *DLL) Delete(node *Node) {
	Head, Tail := dll.Head, dll.Tail

	P, N := node.prev, node.next

	if node == Head {
		nxt := Head.next
		Head.next = nil
		nxt.prev = nil
		Head = nxt

		if node == Tail {
			Tail = nil
		}

		dll.Head, dll.Tail = Head, Tail
		return
	}

	if node == Tail {
		pv := Tail.prev
		Tail.prev = nil
		pv.next = nil
		Tail = pv

		dll.Head, dll.Tail = Head, Tail
		return
	}

	node.next = nil
	node.prev = nil

	P.next = N
	N.prev = P

	dll.Head, dll.Tail = Head, Tail

}

func (dll *DLL) RelocateNode(node *Node) {
	Head, Tail := dll.Head, dll.Tail

	// Moves the node to the tail of linked list

	if node == Tail {
		return
	}

	if node == Head {
		Head = Head.next
		Head.prev = nil
	} else {
		P, N := node.prev, node.next
		P.next = N
		N.prev = P
	}
	node.prev = Tail
	Tail.next = node
	node.next = nil
	Tail = node

	dll.Head, dll.Tail = Head, Tail
}
