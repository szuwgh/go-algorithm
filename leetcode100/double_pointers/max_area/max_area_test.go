package double_pointers

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("最大水量为：", maxArea(height)) // 输出 49
}
