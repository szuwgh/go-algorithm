package maxProduct

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	nums1 := []int{2, 3, -2, 4}
	fmt.Println("最大乘积:", maxProduct(nums1)) // 输出: 6

	nums2 := []int{-2, 0, -1}
	fmt.Println("最大乘积:", maxProduct(nums2)) // 输出: 0
}
