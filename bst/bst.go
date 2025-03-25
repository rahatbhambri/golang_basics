package main

import (
	"fmt"
)

type bstnode struct {
	val   int
	left  *bstnode
	right *bstnode
}

var root *bstnode

func insert(rnode *bstnode, val int) *bstnode {
	if rnode == nil {
		return &bstnode{
			val, nil, nil,
		}
	} else {
		if val < rnode.val {
			rnode.left = insert(rnode.left, val)
		} else {
			rnode.right = insert(rnode.right, val)
		}
	}
	return rnode
}

func inOrder(rnode *bstnode) {
	if rnode == nil {
		return
	}
	if rnode.left != nil {
		inOrder(rnode.left)
	}
	fmt.Println(rnode.val)

	if rnode.right != nil {
		inOrder(rnode.right)
	}
}

func main() {

	root = insert(root, 1)
	root = insert(root, 5)
	root = insert(root, 3)
	root = insert(root, 2)

	inOrder(root)

}
