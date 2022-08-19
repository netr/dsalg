package dsalg

import (
	"math"
)

func HeapSort(arr []int64) []int64 {
	if arr == nil {
		return nil
	}
	n := len(arr)

	for i := math.Max(0, (float64(n)/2)-1); i >= 0; i-- {
		sink(arr, n, int(i))
	}
	for i := n - 1; i >= 0; i-- {
		swap(arr, 0, i)
		sink(arr, i, 0)
	}

	return arr
}
func sink(arr []int64, arrLen int, idx int) []int64 {
	for {
		left := 2*idx + 1
		right := 2*idx + 2
		largest := idx

		if right < arrLen && arr[right] > arr[largest] {
			largest = right
		}
		if left < arrLen && arr[left] > arr[largest] {
			largest = left
		}

		if largest != idx {
			arr = swap(arr, largest, idx)
			idx = largest
		} else {
			break
		}
	}

	return arr
}

func swap(arr []int64, largest int, idx int) []int64 {
	tmp := arr[idx]
	arr[idx] = arr[largest]
	arr[largest] = tmp
	return arr
}
