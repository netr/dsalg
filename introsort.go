package dsalg

import (
	"math"
)

type IntroSorter struct {
	arr []int
	len int
}

func NewIntroSorter(arrLen int) IntroSorter {
	return IntroSorter{
		arr: make([]int, arrLen),
		len: 0,
	}
}

// Sort initializes the sorting algorithm
func (is *IntroSorter) Sort() {
	// Initialise the depthLimit as 2*log(length(data))
	depthLimit := (2 * math.Floor(math.Log(float64(is.len))/
		math.Log(2)))

	is.sortData(0, is.len-1, int(depthLimit))
}

func (is *IntroSorter) Reset(arrLen int) {
	is.arr = make([]int, arrLen)
	is.len = 0
}

func (is *IntroSorter) Arr() []int {
	return is.arr
}

func (is *IntroSorter) Insert(temp int) {
	is.arr[is.len] = temp
	is.len++
}

func (is *IntroSorter) Set(data []int) {
	is.arr = data
	is.len = len(is.arr)
}

func (is *IntroSorter) swap(i int, j int) {
	temp := is.arr[i]
	is.arr[i] = is.arr[j]
	is.arr[j] = temp
}

func (is *IntroSorter) maxHeap(idx, heapLen, begin int) {
	temp := is.arr[begin+idx-1]
	child := 0

	for idx <= heapLen/2 {
		child = 2 * idx
		if child < heapLen && is.arr[begin+child-1] < is.arr[begin+child] {
			child++
		}

		if temp >= is.arr[begin+child-1] {
			break
		}

		is.arr[begin+idx-1] = is.arr[begin+child-1]
		idx = child
	}

	is.arr[begin+idx-1] = temp
}

func (is *IntroSorter) heapify(begin, end, heapLen int) {
	for i := heapLen / 2; i >= 1; i-- {
		is.maxHeap(i, heapLen, begin)
	}
}

func (is *IntroSorter) heapSort(begin, end int) {
	heapLen := end - begin
	// Build heap (rearrange array)
	is.heapify(begin, end, heapLen)
	// One by one extract an element from heap
	for i := heapLen; i >= 1; i-- {
		// Move current root to end
		is.swap(begin, begin+i)
		// call maxHeap() on the reduced heap
		is.maxHeap(1, i, begin)
	}
}

func (is *IntroSorter) insertionSort(left, right int) {
	for i := left; i <= right; i++ {
		key := is.arr[i]
		j := i
		// Move elements of arr[0..i-1], that are
		// greater than the key, to one position ahead
		// of their current position
		for j > left && is.arr[j-1] > key {
			is.arr[j] = is.arr[j-1]
			j--
		}
		is.arr[j] = key
	}
}

func (is *IntroSorter) findPivot(a, b, c int) int {
	max := math.Max(math.Max(float64(is.arr[a]), float64(is.arr[b])), float64(is.arr[c]))
	min := math.Min(math.Min(float64(is.arr[a]), float64(is.arr[b])), float64(is.arr[c]))
	median := int(max) ^ int(min) ^ is.arr[a] ^ is.arr[b] ^ is.arr[c]
	if median == is.arr[a] {
		return a
	}
	if median == is.arr[b] {
		return b
	}
	return c
}

// This function takes the last element as pivot, places
// the pivot element at its correct position in sorted
// array, and places all smaller (smaller than pivot)
// to the left of the pivot
// and greater elements to the right of the pivot
func (is *IntroSorter) partition(low, high int) int {
	pivot := is.arr[high]

	// Index of smaller element
	i := low - 1
	for j := low; j <= high-1; j++ {
		// If the current element is smaller
		// than or equal to the pivot
		if is.arr[j] <= pivot {
			// increment index of smaller element
			i++
			is.swap(i, j)
		}
	}
	is.swap(i+1, high)
	return (i + 1)
}

// The main function that implements Introsort
// low  --> Starting index,
// high  --> Ending index,
// depthLimit  --> recursion level
func (is *IntroSorter) sortData(begin, end, depthLimit int) {
	if end-begin > 16 {
		if depthLimit == 0 {
			// if the recursion limit is
			// occurred call heap sort
			is.heapSort(begin, end)
			return
		}
		depthLimit--
		pivot := is.findPivot(begin, begin+((end-begin)/2)+1, end)
		is.swap(pivot, end)
		// p is partitioning index,
		// arr[p] is now at right place
		p := is.partition(begin, end)
		// Separately sort elements before
		// partition and after partition
		is.sortData(begin, p-1, depthLimit)
		is.sortData(p+1, end, depthLimit)
	} else {
		// if the data set is small,
		// call insertion sort
		is.insertionSort(begin, end)
	}
}
