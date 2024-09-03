package search_matrix

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	matrix1 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target1 := 3
	fmt.Println(SearchMatrix(matrix1, target1)) // 输出：true

	matrix2 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target2 := 13
	fmt.Println(SearchMatrix(matrix2, target2)) // 输出：false
}
