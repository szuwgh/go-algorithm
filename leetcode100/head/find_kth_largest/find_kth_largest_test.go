package find_kth_larget

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	// 示例 1
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	fmt.Printf("数组 %v 中第 %d 大的元素是 %d\n", nums1, k1, findKthLargest(nums1, k1))

	// 示例 2
	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	fmt.Printf("数组 %v 中第 %d 大的元素是 %d\n", nums2, k2, findKthLargest(nums2, k2))
}
