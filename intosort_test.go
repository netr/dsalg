package dsalg

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewIntroSorter(t *testing.T) {
	nums := []int{2, 10, 24, 2, 10, 11, 27,
		4, 2, 4, 28, 16, 9, 8,
		28, 10, 13, 24, 22, 28,
		0, 13, 27, 13, 3, 23,
		18, 22, 8, 8}
	so := NewIntroSorter(len(nums))

	for i := 0; i < len(nums); i++ {
		so.DateAppend(nums[i])
	}

	so.Sort()
	assert.Equal(t, []int{0, 2, 2, 2, 3, 4, 4, 8, 8, 8, 9, 10, 10, 10, 11, 13, 13, 13, 16, 18, 22, 22, 23, 24, 24, 27, 27, 28, 28, 28}, so.Arr())
}

func TestNewIntroSorter_ManyItems(t *testing.T) {
	var nums []int
	for i := 0; i < 100000; i++ {
		nums = append(nums, i)
	}

	shuffled := make([]int, len(nums))
	copy(shuffled, nums[:])

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })

	so := NewIntroSorter(len(shuffled))

	for i := 0; i < len(shuffled); i++ {
		so.DateAppend(shuffled[i])
	}

	so.Sort()
	assert.Equal(t, nums, so.Arr())
}

func BenchmarkNewIntroSorter(b *testing.B) {
	// BenchmarkNewIntroSorter-8   	     100	  10643609 ns/op	 1646650 B/op	       2 allocs/op
	var nums []int
	for i := 0; i < 500000; i++ {
		nums = append(nums, i)
	}
	so := NewIntroSorter(0)

	for i := 0; i < b.N; i++ {
		shuffled := make([]int, len(nums))
		copy(shuffled, nums[:])

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })

		so.Reset(len(shuffled))

		for i := 0; i < len(shuffled); i++ {
			so.DateAppend(shuffled[i])
		}

		so.Sort()
	}
}
