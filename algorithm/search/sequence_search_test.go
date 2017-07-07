package search

import (
	"fmt"
	"testing"
)

func Test_sequenceSearch(t *testing.T) {
	array := []int{1, 3, 5, 7, 8, 9}
	key := 5
	fmt.Println(sequenceSearch(array, key))
}