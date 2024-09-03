package search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		result int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1, 3}, 3, 1},
		{[]int{3, 1}, 3, 0},
	}

	for _, tc := range testCases {
		output := search(tc.nums, tc.target)
		fmt.Printf("Input: nums = %v, target = %d\n", tc.nums, tc.target)
		fmt.Printf("Output: %d\n", output)
		fmt.Printf("Expected: %d\n\n", tc.result)
	}
}
