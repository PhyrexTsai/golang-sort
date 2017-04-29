package sort

import "math/rand"

func QuickSort(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, arr)
		return sliceCopy
	}

	m := arr[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range arr {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}
