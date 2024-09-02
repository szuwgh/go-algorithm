package two_sum

import (
	"fmt"
	"testing"
)

func Test_twoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
