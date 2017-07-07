package sort

import (
	"testing"
	"fmt"
)

func Test_bubbleSort(t *testing.T) {
	array := []int{7, 3, 1, 5, 9}
	bubbleSort(array)
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
