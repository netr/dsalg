package dsalg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	hp := HeapSort([]int64{10, 4, 6, 4, 8, -13, 2, 3})
	assert.Equal(t, []int64{-13, 2, 3, 4, 4, 6, 8, 10}, hp)

	hp = HeapSort([]int64{10, 4, 6, 4, 8, -13, 2, 3})
	assert.NotEqual(t, []int64{10, 8, 6, 4, 4, 3, 2, -13}, hp)
}
