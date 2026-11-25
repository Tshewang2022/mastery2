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
	// check if this is valid binary tree

	// vaild bst, if root.left.val<root.val> root.right.val
	// if root.left.val < root.val && root.val < root.right.val {
	// 	return true
	// }else {
	// 	return false
	// }

	qeueu := []*TreeNode{root}

	if len(qeueu) > 0 {
		// perform the operations;
		// stick to what you know solution, when trying to construct the solutions;
		
	}

}
