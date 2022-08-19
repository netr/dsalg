package dsalg

func mergeSort(arr []int, left int, middle int, right int) []int {
	leftMid := middle - left + 1
	rightMid := right - middle

	L := make([]int, leftMid)
	R := make([]int, rightMid)

	for i := 0; i < leftMid; i++ {
		L[i] = arr[left+i]
	}
	for j := 0; j < rightMid; j++ {
		R[j] = arr[middle+1+j]
	}

	i := 0
	j := 0
	k := left

	for i < leftMid && j < rightMid {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	for i < leftMid {
		arr[k] = L[i]
		i++
		k++
	}
	for j < rightMid {
		arr[k] = R[j]
		j++
		k++
	}
	return arr
}

func MergeSort(arr *[]int, left int, right int) []int {
	if left >= right {
		return []int{}
	}
	middle := left + int((right-left)/2)
	MergeSort(arr, left, middle)
	MergeSort(arr, middle+1, right)
	return mergeSort(*arr, left, middle, right)
}
