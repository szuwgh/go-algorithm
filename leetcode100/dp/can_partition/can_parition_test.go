package can

import (
	"fmt"
	"testing"
)

func TestCanPartition(t *testing.T) {
	nums1 := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums1)) // 输出: true

	nums2 := []int{1, 2, 3, 5}
	fmt.Println(canPartition(nums2)) // 输出: false
}
