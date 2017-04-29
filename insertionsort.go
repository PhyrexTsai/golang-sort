package sort

func InsertionSort(arr []int) []int {
	var j, tmp int
	for i := 0; i < len(arr); i++ {
		tmp = arr[i]
		j = i - 1
		for j > -1 && tmp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp
	}
	return arr
}
