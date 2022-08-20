package bst

import (
	"errors"
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

// NewNode creates a new node with empty left/right nodes
func NewNode(value int) *Node {
	n := &Node{value: value, left: nil, right: nil}
	return n
}

type Tree struct {
	root        *Node
	insertIndex int
}

func NewTree() *Tree {
	t := &Tree{
		root: nil,
	}
	return t
}

// Insert using recursion: O(n^2)
func (tr *Tree) Insert(value ...int) {
	for _, v := range value {
		tr.root = tr.insertRecursive(tr.root, v)
	}
}

func (tr *Tree) insertRecursive(node *Node, value int) *Node {
	if node == nil {
		node = NewNode(value)
		return node
	}

	if value < node.value {
		node.left = tr.insertRecursive(node.left, value)
	} else if value > node.value {
		node.right = tr.insertRecursive(node.right, value)
	}

	return node
}

// PrintInOrderTraversal is a pretty print converted from: https://stackoverflow.com/a/50650932
func (tr *Tree) PrintInOrderTraversal() {
	if tr.root == nil {
		return
	}
	fmt.Printf("%d\n", tr.root.value)
	tr.inOrderTraversalRecursive(tr.root, "")
	fmt.Println("")
}

func (tr *Tree) inOrderTraversalRecursive(node *Node, prefix string) {
	if node != nil {
		hasLeft := node.left != nil
		hasRight := node.right != nil

		if !hasRight && !hasLeft {
			return
		}

		fmt.Print(prefix)
		if hasLeft && hasRight {
			fmt.Print("├── ")
		} else if !hasLeft && hasRight {
			fmt.Print("└── ")
		}

		if hasRight {
			printStrand := hasLeft && hasRight && (node.right.right != nil || node.right.left != nil)
			newPrefix := prefix
			if printStrand {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
			fmt.Printf("%d\n", node.right.value)
			tr.inOrderTraversalRecursive(node.right, newPrefix)
		}
		if hasLeft {
			if hasRight {
				fmt.Print(prefix)
			}
			fmt.Printf("└── %d\n", node.left.value)
			tr.inOrderTraversalRecursive(node.left, prefix+"    ")
		}
	}
}

var ErrValueNotFound = errors.New("value not found")

func (tr *Tree) Search(value int) (*Node, error) {
	if val := tr.searchRecursive(tr.root, value); val != nil {
		return val, nil
	}

	return nil, ErrValueNotFound
}

func (tr *Tree) searchRecursive(node *Node, value int) *Node {
	if node == nil || node.value == value {
		return node
	}

	if value < node.value {
		return tr.searchRecursive(node.left, value)
	}

	return tr.searchRecursive(node.right, value)
}

func (tr *Tree) Delete(value int) *Node {
	return tr.deleteRecursive(tr.root, value)
}

func (tr *Tree) leftMostValue(node *Node) int {
	mv := node.value
	for node.left != nil {
		mv = node.left.value
		node = node.left
	}
	return mv
}

func (tr *Tree) deleteRecursive(node *Node, value int) *Node {
	// fell off or empty tree
	if node == nil {
		return node
	}

	if value < node.value {
		// * start recursion
		node.left = tr.deleteRecursive(node.left, value)
	} else if value > node.value {
		// * start recursion
		node.right = tr.deleteRecursive(node.right, value)
	} else { // * match found, delete it.

		// ! with nodes that have one child, all we need to do is replace the node with its child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		// ! with nodes with two children, we need to traverse the right tree, and find the left most nodes value
		// we replace the current node's value with the left most's value since it will be less than the right node's value and still be in order
		node.value = tr.leftMostValue(node.right)

		// ! now traverse and delete the node we just swapped
		node.right = tr.deleteRecursive(node.right, node.value)
	}

	return node
}

// InsertFast is an O(n) implementation from: https://www.geeksforgeeks.org/archives/3042
func (tr *Tree) InsertFast(value ...int) *Node {
	tr.root = tr.fastConstruct(value)
	return tr.root
}

func (tr *Tree) fastConstruct(values []int) *Node {
	tr.insertIndex = 0
	return tr.fastConstructUntil(values, values[0], math.MinInt, math.MaxInt, len(values))
}

func (tr *Tree) fastConstructUntil(values []int, key, min, max, size int) *Node {
	if tr.insertIndex >= size {
		return nil
	}

	var root *Node

	if key > min && key < max {
		root = NewNode(key)
		tr.insertIndex++

		if tr.insertIndex < size {
			root.left = tr.fastConstructUntil(values, values[tr.insertIndex], min, key, size)
		}
		if tr.insertIndex < size {
			root.right = tr.fastConstructUntil(values, values[tr.insertIndex], key, max, size)
		}
	}

	return root
}
