package main

type linkedlist struct {
	val  int
	next *linkedlist
}

// need more solid foundations of this type;
// ability to visualize and devise solutions
// as you climb up the ladders of solving the problems
// will be much more crucial and important to have very solid foundations

type head struct {
	root *linkedlist
}

func (h *head) insertNode(val int) {
	newNode := &linkedlist{val: val}

	if h.root == nil {
		h.root = newNode
		return
	}

	// let see how the dot connects;
	// inserting the node at last
	curr := h.root
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}
