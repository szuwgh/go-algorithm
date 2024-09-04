package Min

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	grid1 := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid1)) // 输出: 7

	grid2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(minPathSum(grid2)) // 输出: 12
}
