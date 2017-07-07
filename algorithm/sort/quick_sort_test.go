package sort

import (
	"testing"
	"fmt"
)

func Test_quickSort(t *testing.T) {
	array := []int{7, 3, 1, 5, 9}
	quickSort(array, 0, len(array)-1)
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
