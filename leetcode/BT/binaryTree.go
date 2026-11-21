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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	merged := &TreeNode{val: root1.val + root2.val}
	merged.left = mergeTrees(root1.left, root2.right)
	merged.right = mergeTrees(root1.right, root2.right)

	return merged

}

// need to solve 4 problems today;

// reflection on what to improve and how to improve
// need to get into more basic things
// how does the data structure works, and what are there drawbacks;
// need to focus on more basic things;
// i think i dont know the boundry, should be able to tell if it can work
// or if it will fail, because of the building block
// thats where the experience and ability to visualize the problem and solutions
// will come in place
// like traversal methods=> bfs, dfs, recursions;
// what goes in, and what comes out;
// i understands, how it works and but not able to combine them to solve bigger problems
// so basically lacking the fundementalss
// traversal methods => dfs
// stack, queue, linked list, maps and hash table
// after having this solid foundations, we can then say how its fits into my solutions
// will be able to code it out more effectively and efficiently
// even the solution is wrong, i will be able to express my ideas into the code
// thats why i study more in details with regards to fundamentals
// i think i work well with primitives basics
// but around the mid level i struggle, i guess

// stack and queue implementations
// stack, queue is the data structure

// ability to understand slice(arrays) will help termendously,
// because under the hood, most of the complex data structure and algo, use itsr

type stack struct {
	// so we are using the slice
	items []int
}

func (s *stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	front := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return front, true
}

func (s *stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return false
	}
	return true

}

// queue data types;

type queue struct {
	items []int
}

func (q *queue) enqueue(val int) {

	q.items = append(q.items, val)

}

func (q *queue) dequeue() {
	if q.IsEmpty() {
		return
	}
	q.items = q.items[1:]
}

func (q *queue) IsEmpty() bool {
	if len(q.items) == 0 {
		return true
	}
	return false
}

// for the tree, we need understanding of the recursions
// traversal methods and types;
// maybe my understanding of linked list is limited;

// stack for integers
func levelOrderTraversal2(root *TreeNode) {
	q := &Queue[*TreeNode]{}
	q.Enqueue(root)
	if !q.IsEmpty() {
		root, _ := q.Dequeue()
		if root.left != nil {
			q.Enqueue(root.left)
		}

		if root.right != nil {
			q.Enqueue(root.right)
		}

		fmt.Println()

	}

}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	var averages []float64
	queue := []*TreeNode{root}

	// wanted to look for this solution;
	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {

			// so everything from here will depend on the context of the node
			node := queue[0]
			queue = queue[1:] // Dequeue front element

			levelSum += node.val

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.left != nil {
				queue = append(queue, node.right)
			}
		}

		averages = append(averages, float64(levelSum)/float64(levelSize))
	}

	return averages

}

func findTarget(root *TreeNode, k int) bool {
	var vals []int

	var collect func(node *TreeNode)
	collect = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.val)
		collect(node.left)
		collect(node.right)
	}
	collect(root)

	for i := 0; i < len(vals); i++ {
		for j := i + 1; j < len(vals); j++ {
			if vals[i]+vals[j] == k {
				return true
			}
		}
	}
	return false
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.val == val {
		return root
	}

	if val < root.val {
		return searchBST(root.left, val)
	}

	return searchBST(root.right, val)
}

func checkTree(root *TreeNode) bool {
	// questions?"":"";
	// give tree has only 3 nodes, 1 root, 2 children
	// if root.Val == root.left.val + root.right.val ?"true":"false"

	if root == nil {
		return false
	}

	sum := root.left.val + root.right.val
	if sum == root.val {
		return true
	}
	return false
}

func isCousins(root *TreeNode, x int, y int) bool {

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		foundX, foundY := false, false

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// Same parent => siblings, not cousins
			if node.left != nil && node.right != nil {
				l, r := node.left.val, node.right.val
				if (l == x && r == y) || (l == y && r == x) {
					return false
				}
			}

			if node.val == x {
				foundX = true
			}
			if node.val == y {
				foundY = true
			}

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

		// Both found on the same level (and not siblings)
		if foundX && foundY {
			return true
		}
		// Only one found on this level => different depths
		if foundX || foundY {
			return false
		}
	}

	return false
}

// to have more leverage, have to put more effort and think more
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}

	if isSame(root, subRoot) {
		return true
	}

	return isSubtree(root.left, subRoot) || isSubtree(root.right, subRoot)
}

func isSame(a, b *TreeNode) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.val != b.val {
		return false
	}
	return isSame(a.left, b.left) && isSame(a.right, b.right)

}

func increasingBST(root *TreeNode) *TreeNode {
	// so this should be implemented here;
	newNode := &TreeNode{}
	tail := newNode

	if root == nil {
		return nil
	}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.left)
		// attach node to tail, and move tail forward
		tail.right = &TreeNode{val: node.val}
		tail = tail.right

		inorder(node.right)
	}
	inorder(root)
	return newNode.right
}

func inorder2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	newTree := &TreeNode{val: root.val}
	newTree.left = inorder2(root.left)
	newTree.right = inorder2(root.right)

	return newTree

}
