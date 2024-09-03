package search_insert

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums1 := []int{1, 3, 5, 6}
	target1 := 5
	fmt.Println(searchInsert(nums1, target1)) // 输出: 2

	nums2 := []int{1, 3, 5, 6}
	target2 := 2
	fmt.Println(searchInsert(nums2, target2)) // 输出: 1

	nums3 := []int{1, 3, 5, 6}
	target3 := 7
	fmt.Println(searchInsert(nums3, target3)) // 输出: 4
}
