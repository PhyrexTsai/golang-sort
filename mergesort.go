package sort

func merge(x, y []int) []int {
	if len(x) == 0 {
		return y
	}
	if len(y) == 0 {
		return x
	}
	m := 0
	n := 0
	var arr []int
	for m < len(x) && n < len(y) {
		if x[m] <= y[n] {
			arr = append(arr, x[m])
			m++
		} else if x[m] > y[n] {
			arr = append(arr, y[n])
			n++
		}
	}
	if len(x) == m {
		for n < len(y) {
			arr = append(arr, y[n])
			n++
		}
	}
	if len(y) == n {
		for m < len(x) {
			arr = append(arr, x[m])
			m++
		}
	}
	return arr
}

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	var half int = len(arr) / 2
	var m []int
	var n []int
	for i := 0; i < len(arr); i++ {
		if i < half {
			m = append(m, arr[i])
		} else {
			n = append(n, arr[i])
		}
	}
	return merge(MergeSort(m), MergeSort(n))
}
