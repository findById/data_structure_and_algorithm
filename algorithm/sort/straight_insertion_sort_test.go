package sort

import (
	"testing"
	"fmt"
)

func Test_straightInsertionSort(t *testing.T) {
	array := []int{7, 3, 1, 5, 9}
	straightInsertionSort(array)
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
