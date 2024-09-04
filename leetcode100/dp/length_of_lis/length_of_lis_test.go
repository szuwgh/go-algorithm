package lengthOfLIS

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	nums1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("Length of LIS:", lengthOfLIS(nums1)) // 输出: 4

	nums2 := []int{0, 1, 0, 3, 2, 3}
	fmt.Println("Length of LIS:", lengthOfLIS(nums2)) // 输出: 4

	nums3 := []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println("Length of LIS:", lengthOfLIS(nums3)) // 输出: 1
}
