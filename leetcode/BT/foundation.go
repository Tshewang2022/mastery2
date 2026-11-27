// need strong foundation, and still havent made good use  or has not got stronger
package main

import (
	"fmt"
	"slices"
)

type stacks struct {
	items *[]int
}

type queues struct {
	items *[]int
}

func basics() {

	// array=> slice,
	// queue, stack, linked list

	fmt.Println("hello golango")
}

// at very basic level computer programming is just
// your instruction => encoding(asscii || utf-8 encoding) -> binary

// so in terms of network=> binary ->  light signals(physical layers)

// so in very basic terms, how can we read from and write to the different locations

// Software at the most fundamental level is just a CRUD operations on memory
// all other is just noise for doing something efficiently and securely in the memory level

// grind the fundamentals

// add, update, delete
func arrOps() []int {
	var arr []int
	arr = append(arr, 1)
	arr = append(arr, 2)

	// update
	arr[0] = 2

	// delete
	arr = slices.Delete(arr, 1, 2)

	return arr
}

// linked list
// stack
// queue

// Linked list
type NodeStruct struct {
	Val  int
	Next *NodeStruct
}

type HeadNode struct {
	Head *NodeStruct
}

// linked list operation => insertNode, updateNode, deleteNode && readNode
func (root *NodeStruct) insertNode(t int) *NodeStruct {
	// prepare new node
	newNode := &NodeStruct{Val: t}

	// this is the head node, only if its root == nil
	if root == nil {
		root = newNode
	}

	// insert at the tail
	// navigating
	curr := root
	// this should navigating
	for curr.Next != nil {
		curr = curr.Next
	}
	root = newNode
	// insert at the head node
	return root
}

func (root *NodeStruct) updateNode(pos int, newVal int) bool {
	if pos < 0 {
		return false
	}

	curr := root
	index := 0

	for curr != nil {
		if index == pos {
			curr.Val = newVal
			return true
		}
		curr = curr.Next
		index++
	}
	return false
}

func (root *HeadNode) insertAtParticularPosition(pos, newVal int) bool {

	if pos < 0 {
		return false
	}
	newNode := &NodeStruct{Val: newVal}

	curr := root.Head
	index := 0
	if pos == 0 {
		newNode.Next = root.Head
		root.Head = newNode
		return true
	}

	for curr != nil && index < pos-1 {
		curr = curr.Next
		index++
	}

	if curr == nil {
		return false
	}

	newNode.Next = curr.Next
	curr.Next = newNode
	return true
}
