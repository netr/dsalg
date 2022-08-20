package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	tree := NewTree()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(20)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(16)
	tree.Insert(100)
	tree.Insert(85)

	assert.NotNil(t, tree.root.value)
	assert.NotNil(t, tree.root.left)
	assert.NotNil(t, tree.root.right)

	tree.PrintInOrderTraversal()
}

func TestTree_Search_NotFound(t *testing.T) {
	tree := NewTree()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)

	_, err := tree.Search(7)
	assert.ErrorIs(t, err, ErrValueNotFound)
}

func TestTree_Search(t *testing.T) {
	tree := NewTree()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(2)

	v, err := tree.Search(6)
	assert.Nil(t, err)
	assert.Equal(t, 6, v.value)
}

func TestTree_Delete(t *testing.T) {
	tree := NewTree()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(2)

	v, err := tree.Search(6)
	assert.Nil(t, err)
	assert.Equal(t, 6, v.value)
	tree.PrintInOrderTraversal()

	del := tree.Delete(4)
	assert.NotNil(t, del)

	v, err = tree.Search(4)
	assert.NotNil(t, err)
	tree.PrintInOrderTraversal()
}
