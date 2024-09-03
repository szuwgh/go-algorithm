package find_min

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	nums1 := []int{3, 4, 5, 1, 2}
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	nums3 := []int{11, 13, 15, 17}

	fmt.Println(findMin(nums1)) // Output: 1
	fmt.Println(findMin(nums2)) // Output: 0
	fmt.Println(findMin(nums3)) // Output: 11
}
