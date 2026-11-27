package main

// 95. Unique binary search trees II
type rootNode struct {
	root *TreeNode
}

func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return []*TreeNode{}
	}

	return build(1, n)

}

func build(lo, hi int) []*TreeNode {

	if lo > hi {
		return []*TreeNode{nil}
	}

	var trees []*TreeNode
	for i := lo; i <= hi; i++ {
		leftTrees := build(lo, i-1)
		rightTree := build(i+1, hi)

		// i did not understand this, how it works, and the logic behind it
		for _, left := range leftTrees {
			for _, right := range rightTree {
				trees = append(trees, &TreeNode{val: i, left: left, right: right})
			}
		}
	}
	return trees

}

// classic recursion solutions with time complexity of O(n*2)
// func numTrees(n int) int {
// 	return countTrees(1, n)

// }

// func countTrees(lo, hi int) int {
// 	totalCount := 0
// 	if lo > hi {
// 		return 1
// 	}

// 	for i := lo; i <= hi; i++ {
// 		leftTreeCount := countTrees(lo, i-1)
// 		rightTreeCount := countTrees(i+1, hi)

// 		totalCount += leftTreeCount * rightTreeCount
// 	}
// 	return totalCount

// }

// now lets try with memoizations

func numTrees(n int) int {
	memo := make([]int, n+1)

	return count(n, memo)

}

func count(nodes int, memo []int) int {
	if nodes <= 1 {
		return 1
	}

	if memo[nodes] != 0 {
		return memo[nodes]
	}

	total := 0
	for root := 1; root <= nodes; root++ {
		total += count(root-1, memo) * count(nodes-root, memo)
	}

	memo[nodes] = total
	return total
}

func isValidBST(root *TreeNode) bool {
	// can be implemented both using queue and stack

	stack := []*TreeNode{}
	var prev *int
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil && curr.val <= *prev {
			return false
		}

		prev = &curr.val
		curr = curr.right
	}

	return true

}

// how to make every minute count
// what it takes be a champion and winner
func recoverTree(root *TreeNode) {
	// values of exactly two nodes of the tree were swapped by mistake, recover the tree without
	// changing its structure

}
