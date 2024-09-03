package can_jump

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	// 测试用例
	nums1 := []int{2, 3, 1, 1, 4}
	nums2 := []int{3, 2, 1, 0, 4}

	fmt.Println(canJump(nums1)) // 输出: true
	fmt.Println(canJump(nums2)) // 输出: false
}
