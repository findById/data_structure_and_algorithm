package sort

func quickSort(array []int, start, end int) {
	if start < end {
		i, j := start, end
		key := array[(start+end)/2]
		for i <= j {
			for array[i] < key {
				i++
			}
			for array[j] > key {
				j--
			}
			if i <= j {
				array[i], array[j] = array[j], array[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort(array, start, j)
		}
		if end > i {
			quickSort(array, i, end)
		}
	}
}
