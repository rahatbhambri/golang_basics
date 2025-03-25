package dll

type Node struct {
	Val  []int
	next *Node
	prev *Node
}

var Head *Node
var Tail *Node

// func main() {

// 	appendLeft(4)
// 	printList()
// 	appendLeft(5)
// 	printList()
// 	appendLeft(6)
// 	printList()
// 	appendLeft(7)
// 	printList()
// 	appendRight(10)
// 	appendRight(9)
// 	appendRight(8)
// 	printList()
// 	insert(2, 20)
// 	insert(2, 40)
// 	insert(2, 60)

// 	println("Deletion!")
// 	printList()
// 	delete(2)

// 	printList()
// 	delete(2)

// 	printList()
// 	delete(2)
// 	printList()
// }

func AppendRight(Val []int) *Node {
	if Head == nil {
		Head = &Node{Val, nil, nil}
		Tail = Head
	} else {
		nn := &Node{Val, nil, Tail}
		Tail.next = nn
		Tail = nn
	}
	return Tail
}

func AppendLeft(Val []int) {
	if Head == nil {
		Head = &Node{Val, nil, nil}
		Tail = Head
	} else {
		nn := &Node{Val, nil, Tail}
		Tail.next = nn
		Tail = nn
	}
}

func GetNodes(pos int) (*Node, *Node, *Node) {
	c := 0
	temp := Head

	for {
		if pos != c {
			temp = temp.next
			c += 1
		} else {
			break
		}
	}

	var curr_pre *Node
	var curr_nxt *Node

	if temp.prev != nil {
		curr_pre = temp.prev
	} else {
		curr_pre = nil
	}

	if temp.next != nil {
		curr_nxt = temp.next
	} else {
		curr_nxt = nil
	}

	return curr_pre, temp, curr_nxt
}

func Insert(pos int, Val []int) {

	curr_pre, temp, _ := GetNodes(pos)
	nn := &Node{Val, temp, curr_pre}
	temp.prev = nn
	curr_pre.next = nn
}

func PrintList() {

	temp := Head

	for {
		if temp != nil {
			print(temp.Val, " ")
			temp = temp.next
		} else {
			break
		}
	}
	println()
}

func Delete(pos int) {
	curr_pre, temp, curr_nxt := GetNodes(pos)

	if pos == 0 {
		Head = Head.next
	}

	if curr_pre != nil {
		curr_pre.next = curr_nxt
	}

	if curr_nxt != nil {
		curr_nxt.prev = curr_pre
	}

	temp.next = nil
	temp.prev = nil
}

func RelocateNode(node *Node) {

	if node == Tail {
		return
	}

	curr_pre := node.prev
	curr := node
	curr_nxt := node.next

	if curr_pre == nil {
		Head = curr_nxt
	}

	curr.prev = Tail
	curr.next = nil
	Tail.next = node
	Tail = curr

	if curr_pre != nil {
		curr_pre.next = curr_nxt
	}

	if curr_nxt != nil {
		curr_nxt.prev = curr_pre
	}

}
