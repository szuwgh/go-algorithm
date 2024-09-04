package maxSubArray

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("最大子数组和:", maxSubArray(nums)) // 输出: 6
}
