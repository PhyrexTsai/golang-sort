package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	var arr []int
	n := 10

	arr = make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}

	var result []int = make([]int, len(arr))
	copy(result, arr)
	sort.Ints(result)
	arr = SelectionSort(arr)
	for key, value := range result {
		if arr[key] != value {
			t.Error("Error, got: ", arr)
		}
	}
}
