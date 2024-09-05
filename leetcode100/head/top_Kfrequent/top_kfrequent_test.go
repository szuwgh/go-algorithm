package topK

import (
	"fmt"
	"testing"
)

func TestKFrequent(t *testing.T) {
	// 示例 1
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := 2
	fmt.Println(topKFrequent(nums1, k1)) // 输出: [1 2]

	// 示例 2
	nums2 := []int{1}
	k2 := 1
	fmt.Println(topKFrequent(nums2, k2)) // 输出: [1]
}
