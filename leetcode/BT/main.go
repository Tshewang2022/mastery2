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
	if curr.val > val {
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

// printing of binary node
func inorder(n *Node) {
	if n == nil {
		return
	}
	inorder(n.left)
	fmt.Print(n.val, " ")
	inorder(n.right)
}

func preorder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.val, " ")
	preorder(n.left)
	inorder(n.right)
}

func postorder(n *Node) {
	if n == nil {
		return
	}
	postorder(n.left)
	postorder(n.right)
	fmt.Print(n.val, " ")
}
func main() {
	// tree := &Head{}
	// tree.binaryTree(9)
	// tree.binaryTree(56)
	// tree.binaryTree(4)
	// // tree.insertNode(9)
	// preorder(tree.root)
	// // inorder(tree.root)
	// // postorder(tree.root)
	// // tree.insertNode(90)
	// // tree.printNode()
	// fmt.Println()

	// ---- isBalanced checks ----

	//     2
	//    / \
	//   1   3
	// balanced := &TreeNode{val: 2,
	// 	left:  &TreeNode{val: 1},
	// 	right: &TreeNode{val: 3},
	// }

	//   1
	//    \
	//     2
	//      \
	//       3
	skewed := &TreeNode{val: 1,
		right: &TreeNode{val: 2,
			right: &TreeNode{val: 3},
		},
	}

	// root looks fine (height 3 vs 3), but each side is a chain
	//         1
	//        / \
	//       2   5
	//      /     \
	//     3       6
	//    /         \
	//   4           7

	// fmt.Println("balanced:", isBalanced(balanced)) // true
	fmt.Println("sweked:", isBalanced(skewed))

}

// needs to get down really to the fundemental property
// how does the args passing methods works and how does the conditions executed
// by value or by references, bridging the gap between impl and concepts
// most think we know, but in practical impl we dont really know;
// how does the function calls and code are being executed in the computer level
// this will really helps streamline the thought process with code logic in computer level

// today i will train like a champion;
