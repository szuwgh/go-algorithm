package merge_two_sorted_lists

import "testing"

func Test_MergeTwoSortedList(t *testing.T) {
	l1 := &LinkNode{Value: 1}
	l1.Add(2)
	l1.Add(4)

	l2 := &LinkNode{Value: 1}
	l2.Add(3)
	l2.Add(5)

	l3 := mergeTwoSortedList(l1, l2)
	l3.Print()
}
