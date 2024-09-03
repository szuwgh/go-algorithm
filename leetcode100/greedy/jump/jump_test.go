package jump

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	nums1 := []int{2, 3, 1, 1, 4}
	nums2 := []int{2, 3, 0, 1, 4}

	fmt.Println(jump(nums1)) // Output: 2
	fmt.Println(jump(nums2)) // Output: 2
}
