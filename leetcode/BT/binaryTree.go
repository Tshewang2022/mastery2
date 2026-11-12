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
	fmt.Printf("enter node %d\n", root.val)

	// here it computes only the height;
	left_h := height(root.left)
	fmt.Printf("  node %d: left_h  = %d\n", root.val, left_h)

	right_h := height(root.right)
	fmt.Printf("  node %d: right_h = %d\n", root.val, right_h)

	h := max(left_h, right_h) + 1 
	// fmt.Printf("return node %d: height = %d\n", root.val, h)
	return h

}
