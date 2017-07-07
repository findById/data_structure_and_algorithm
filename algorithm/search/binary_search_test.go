package search

import (
	"fmt"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	array := []int{1, 3, 5, 7, 8, 9}
	key := 5
	fmt.Println(binarySearch(array, key))
	fmt.Println(binarySearchRecursion(array, key, 0, len(array)-1))
}
