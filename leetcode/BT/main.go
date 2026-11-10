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

// this is how the node works;

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

// linked list
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

// need to have a rotation;

// how can we construct binary tree;
func (node *Head) binaryTree(val int) {
	newNode := &Node{val: val}
	// root node;
	if node.root == nil {
		// root node;
		node.root = newNode
		return
	}

	curr := node.root
	if curr.val < val {
		if curr.left == nil {
			curr.left = newNode
			return
		}
		curr = curr.left
	} else {
		if curr.right == nil {
			curr.right = newNode
			return
		}
		curr = curr.right
	}
}

func main() {
	tree := &Head{}
	tree.insertNode(9)
	tree.insertNode(90)
	tree.printNode()
}
