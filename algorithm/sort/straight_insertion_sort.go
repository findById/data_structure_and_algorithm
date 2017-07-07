package sort

func straightInsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			j := i - 1
			temp := array[i]
			for j >= 0 && array[j] > temp {
				array[j+1] = array[j]
				j--
			}
			array[j+1] = temp
		}
	}
}
