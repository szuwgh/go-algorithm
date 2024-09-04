package productExceptSelf

import (
	"fmt"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	result1 := productExceptSelf(nums1)
	fmt.Println(result1) // 输出: [24, 12, 8, 6]

	nums2 := []int{-1, 1, 0, -3, 3}
	result2 := productExceptSelf(nums2)
	fmt.Println(result2) // 输出: [0, 0, 9, 0, 0]
}
