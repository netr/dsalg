package dsalg

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	hp := HeapSort([]int64{10, 4, 6, 4, 8, -13, 2, 3})
	assert.Equal(t, []int64{-13, 2, 3, 4, 4, 6, 8, 10}, hp)

	hp = HeapSort([]int64{10, 4, 6, 4, 8, -13, 2, 3})
	assert.NotEqual(t, []int64{10, 8, 6, 4, 4, 3, 2, -13}, hp)
}

func BenchmarkNewHeapSort(b *testing.B) {
	// BenchmarkNewHeapSort-8   	      99	  11483141 ns/op	  844250 B/op	       1 allocs/op
	var nums []int64
	var i int64
	for i = 0; i < 500000; i++ {
		nums = append(nums, i)
	}

	for i := 0; i < b.N; i++ {
		shuffled := make([]int64, len(nums))
		copy(shuffled, nums[:])

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })

		HeapSort(shuffled)
	}
}
