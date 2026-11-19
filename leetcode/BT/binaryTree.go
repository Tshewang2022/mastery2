package main

import (
	"fmt"
	"strconv"
)

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

// func insertBalancedTree(root *TreeNode, val int) *TreeNode {
// 	newNode := &TreeNode{val: val}

// 	if root == nil {
// 		root = newNode
// 	}
// 	if val < root.val {
// 		root.left = insertBalancedTree(root.left, val)
// 		if isBalanced(root.left) == false {
// 			// performs rotations;
// 			rotation(root.left)
// 		}

// 	} else {
// 		root.right = insertBalancedTree(root.right, val)
// 		if isBalanced(root.right) == false {
// 			// performs rotations
// 			rotation(root.right)
// 		}
// 	}
// 	return root

// }

func findTargetSum(targetSum int, root *TreeNode) bool {
	// this is going to be another base case
	if root == nil {
		return false
	}

	targetSum -= root.val

	if root.left == nil && root.right == nil {
		return targetSum == 0
	}

	leftSum := findTargetSum(targetSum, root.left)
	rightSum := findTargetSum(targetSum, root.right)

	if leftSum == false && rightSum == false {
		return false
	}

	return true
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// just need to swap the children;
	current := root.left
	root.left = root.right
	root.right = current

	invertTree(root.left)
	invertTree(root.right)

	return root

}

func levelOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	// need to write the implmentation of queue
	// when the queue is not empty;
	q := &Queue[*TreeNode]{}
	q.Enqueue(root)

	for !q.IsEmpty() {
		root, _ := q.Dequeue()
		fmt.Print(root.val, " ")
		if root.left != nil {
			q.Enqueue(root.left)
		}

		if root.right != nil {
			q.Enqueue(root.right)
		}
	}

	fmt.Println()

}

func binaryTreePaths(root *TreeNode) []string {

	var result []string

	if root == nil {
		return result
	}

	if root.left == nil && root.right == nil {
		return []string{strconv.Itoa(root.val)}
	}

	for _, p := range binaryTreePaths(root.left) {
		result = append(result, strconv.Itoa(root.val)+"->"+p)
	}

	for _, p := range binaryTreePaths(root.right) {
		result = append(result, strconv.Itoa(root.val)+"->"+p)
	}

	return result

}

func returnSlice(val int) []int {
	if val != 0 {
		return []int{val}
	}
	return nil
}

func sumOfLeftLeaves(root *TreeNode) int {

	sum := 0
	if root == nil {
		return 0
	}

	if root.left != nil && root.left.left == nil && root.left.right == nil {
		sum += root.left.val
	}

	// this means it has reached leave node
	sum += sumOfLeftLeaves(root.right)
	sum += sumOfLeftLeaves(root.left)

	return sum

}

func findMode(root *TreeNode) []int {
	var mode []int
	var prev *TreeNode
	currentCount, maxCount := 0, 0

	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.left)

		if prev != nil && prev.val == node.val {
			currentCount++
		} else {
			currentCount = 1
		}

		if currentCount > maxCount {
			maxCount = currentCount
			mode = []int{node.val}

		} else if currentCount == maxCount {
			mode = append(mode, node.val)
		}

		prev = node
	}
	inorder(root)

	return mode
}

func getMinimumDifference(root *TreeNode) int {
	var minDiff int
	var prev *TreeNode
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.left)

		if prev != nil {
			diff := node.val - prev.val
			if diff < minDiff {
				minDiff = diff
			}
		}

		prev = node
		inorder(node.right)
	}
	inorder(root)

	return minDiff
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := height(root.left)
		right := height(root.right)

		// this is the main logic of the code;
		if left+right > maxDiameter {
			maxDiameter = left + right
		}

		if right > left {
			return right + 1
		}
		return left + 1
	}
	height(root)
	return maxDiameter
}

// need to find this algorithm myself

func findTilt(root *TreeNode) int {
	tiltSum := 0
	var tilt func(node *TreeNode) int
	tilt = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := tilt(node.left)
		rightSum := tilt(node.right)
		diff := abs(leftSum, rightSum)
		tiltSum += diff
		return leftSum + rightSum + node.val

	}
	tilt(root)
	return tiltSum

}
