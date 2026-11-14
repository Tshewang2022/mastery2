package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(height(root.left), height(root.right)) > 1 {
		return false
	}
	return isBalanced(root.left) && isBalanced(root.right)

}

// breaking down into small parts;
func abs(x, y int) int {
	if x >= y {
		return x - y
	}
	return y - x
}

func height(root *TreeNode) int {

	// how does this code dry run works?
	if root == nil {
		return 0
	}
	// here it computes only the height;
	left_h := height(root.left) + 1
	right_h := height(root.right) + 1

	h := max(left_h, right_h)
	return h

}

// working on the insert

func insertNode(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{val: val}
	}

	if val < root.val {
		root.left = insertNode(root.left, val)
	} else {
		root.right = insertNode(root.right, val)
	}
	return root

}

func inorders(n *TreeNode) {

	if n == nil {
		return
	}
	inorders(n.left)
	fmt.Print(n.val, " ")
	inorders(n.right)
}

// rotation; all this concepts of rotations are just the concepts of the swaping between
// state variables, simple as that;
func rightRotation(root *TreeNode) *TreeNode {
	current := root.left
	root.left = current.right
	current.right = root
	return root
}

func leftRotation(root *TreeNode) *TreeNode {
	current := root.right
	root.right = current.left
	current.left = root
	return root
}

func insertBalancedTree(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{val: val}

	if root == nil {
		root = newNode
	}
	if val < root.val {
		root.left = insertBalancedTree(root.left, val)
		if isBalanced(root.left) == false {
			// performs rotations;
			rotation(root.left)
		}

	} else {
		root.right = insertBalancedTree(root.right, val)
		if isBalanced(root.right) == false {
			// performs rotations
			rotation(root.right)
		}
	}
	return root

}
