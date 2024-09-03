package search_range

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums1 := []int{5, 7, 7, 8, 8, 10}
	target1 := 8
	fmt.Println(searchRange(nums1, target1)) // 输出: [3, 4]

	nums2 := []int{5, 7, 7, 8, 8, 10}
	target2 := 6
	fmt.Println(searchRange(nums2, target2)) // 输出: [-1, -1]

	nums3 := []int{}
	target3 := 0
	fmt.Println(searchRange(nums3, target3)) // 输出: [-1, -1]
}
