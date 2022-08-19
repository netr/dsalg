package dsalg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	arr := &[]int{12, 11, 13, 5, 6, 7, 0, 2, 3, 4, -52, 230, 20210, 2004, 2201, 20, 6, 4, 686, 67, 6, 764}
	arrSize := len(*arr)
	s := MergeSort(arr, 0, arrSize-1)
	assert.Equal(t, []int{-52, 0, 2, 3, 4, 4, 5, 6, 6, 6, 7, 11, 12, 13, 20, 67, 230, 686, 764, 2004, 2201, 20210}, s)
}
