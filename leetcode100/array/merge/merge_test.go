package merge

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	// 示例 1
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals1)) // 输出: [[1,6],[8,10],[15,18]]

	// 示例 2
	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals2)) // 输出: [[1,5]]
}
