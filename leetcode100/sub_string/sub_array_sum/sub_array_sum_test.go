package main

import (
	"fmt"
	"testing"
)

func TestSubArraySum(t *testing.T) {
	nums1 := []int{1, 1, 1}
	k1 := 2
	fmt.Println(subarraySum(nums1, k1)) // 输出: 2

	nums2 := []int{1, 2, 3}
	k2 := 3
	fmt.Println(subarraySum(nums2, k2)) // 输出: 2
}
