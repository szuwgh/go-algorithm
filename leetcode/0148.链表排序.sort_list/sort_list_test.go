package sort_list

import "testing"

func Test_SortList(t *testing.T) {
	l1 := &linkNode{Value: 4}
	l1.Add(3)
	l1.Add(1)
	l1.Add(2)
	res := sortList(l1)
	res.Print()
}
