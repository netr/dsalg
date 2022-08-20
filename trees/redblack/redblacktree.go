package redblack

// Comparison with AVL Tree:
// The AVL trees are more balanced compared to Red-Black Trees, but they may cause more rotations during insertion and deletion.
// So if your application involves frequent insertions and deletions, then Red-Black trees should be preferred.
//And if the insertions and deletions are less frequent and search is a more frequent operation, then AVL tree should be preferred over the Red-Black Tree.
// https://www.geeksforgeeks.org/red-black-tree-set-1-introduction-2/?ref=leftbar-rightbar

type Node struct {
	color byte
}

func NewNode(color byte) *Node {
	n := &Node{color: color}
	return n
}

type Tree struct {
	parent *Node
}

// every path from a node (including root) to any of its descendants NULL nodes has the same number of black nodes
// all leaf nodes are black nodes
const clrBlk = '0'

// can not have a red parent or red child
const clrRed = '1'

func NewTree() *Tree {
	t := &Tree{parent: NewNode(clrBlk)}
	return t
}
