package rotate

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	rotate(nums1, k1)
	fmt.Println(nums1) // 输出: [5, 6, 7, 1, 2, 3, 4]

	nums2 := []int{-1, -100, 3, 99}
	k2 := 2
	rotate(nums2, k2)
	fmt.Println(nums2) // 输出: [3, 99, -1, -100]
}
