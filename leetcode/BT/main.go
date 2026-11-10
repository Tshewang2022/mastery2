package main

import "fmt"

// linked list, thats how it works
type Node struct {
	val   int
	left  *Node
	right *Node
}
type Head struct {
	root *Node
}

func (node *Head) insertNode(val int) {
	newNode := &Node{
		val: val,
	}
	// node.root == the first node of the list
	if node.root == nil {
		node.root = newNode
		return
	}
	// left = previous, right = next
	// head insert
	newNode.right = node.root
	node.root.left = newNode

	// so now head node is newNode
	node.root = newNode
}

func (node *Head) middleInsert(val int, position int) {

	if position <= 0 || node.root == nil {
		node.insertNode(val)
		return
	}
	if node.root == nil {
		return
	}

	current := node.root

	// as soon as this loop exit, insert it there;
	for i := 0; i < position-1 && current.right != nil; i++ {
		current = current.right

	}
	newNode := &Node{val: val}
	newNode.left = current
	newNode.right = current.right

	if current.right != nil {
		current.right.left = newNode
	}
	current.right = newNode
}

func (node *Head) printNode() {
	current := node.root

	for current != nil {
		fmt.Print(current.val, " ")
		current = current.right
	}
	// len is only for the arrays and map
	fmt.Println()

}

func main() {
	tree := &Head{}
	tree.insertNode(9)
	tree.insertNode(90)
	tree.printNode()
}
