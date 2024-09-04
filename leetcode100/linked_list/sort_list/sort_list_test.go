package sort_list

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	// 测试用例
	testCases := [][]int{
		{4, 2, 1, 3},
		{-1, 5, 3, 4, 0},
		{},
	}

	for _, testCase := range testCases {
		head := arrayToList(testCase)
		sortedList := sortList(head)
		result := listToArray(sortedList)
		fmt.Println(result)
	}
}
