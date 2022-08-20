package bst

import (
	"errors"
	"fmt"
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
	root *Node
}

func NewTree() *Tree {
	t := &Tree{
		root: nil,
	}
	return t
}

func (tr *Tree) Insert(value int) {
	tr.root = tr.insertRecursive(tr.root, value)
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
	if node == nil {
		return node
	}

	if value < node.value {
		// recursively
		node.left = tr.deleteRecursive(node.left, value)
	} else if value > node.value {
		node.right = tr.deleteRecursive(node.right, value)
	} else { // if key is same as root's key, then This is the node to be deleted
		// node with only one child or no child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		// node with two children: Get the inorder successor (smallest in the right subtree)
		node.value = tr.leftMostValue(node.right)
		// Delete the inorder successor
		node.right = tr.deleteRecursive(node.right, node.value)
	}

	return node
}
