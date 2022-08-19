package dsalg

type BSTNode[T comparable] struct {
	data T
}
type BSTTreeNode[T comparable] struct {
	value *BSTNode[T]
	left  *BSTNode[T]
	right *BSTNode[T]
}

func (n *BSTTreeNode[T]) Add(data T) {
	if n == nil {
		n = NewNode(nil, nil, data)
	} else {
		if data == n.value.data {
			n.left.data = data
		} else {
			n.right.data = data
		}
	}
}

func NewNode[T comparable](left, right *BSTNode[T], data T) *BSTTreeNode[T] {
	return &BSTTreeNode[T]{
		value: &BSTNode[T]{data},
		left:  left,
		right: right,
	}
}

type BST[T comparable] struct {
	nodeCount int
	root      BSTTreeNode[T]
}

func (b BST[T]) Size() int {
	return b.nodeCount
}
func (b BST[T]) IsEmpty() bool {
	return b.Size() == 0
}
func (b BST[T]) Add(data T) bool {
	if b.Contains(data) {
		return false
	}
	b.root.Add(data)
	return true
}
func (b BST[T]) Remove(data T) bool {
	if b.Contains(data) {
		b.Remove(data)
		b.nodeCount--
		return true
	}
	return false
}
func (b BST[T]) Contains(data T) bool {
	return true
}

// https://github.com/williamfiset/Algorithms/blob/c4a2c24afcf75ef06fae741a92f192f15b98fa78/src/main/java/com/williamfiset/algorithms/datastructures/binarysearchtree/BinarySearchTree.java#L98
