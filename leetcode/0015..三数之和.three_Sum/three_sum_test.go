package three_Sum

import (
	"fmt"
	"testing"
)

func Test_ThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
