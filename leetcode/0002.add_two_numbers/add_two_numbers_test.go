package add_two_numbers

import (
	"linklist"
	"testing"
)

func Test_AddTwoNumbers(t *testing.T) {
	l1 := &linklist.LinkNode{Value: 2}
	l2 := &linklist.LinkNode{Value: 5}

	l1.Add(4)
	l1.Add(3)

	l2.Add(6)
	l2.Add(4)

	l3 := addTwoNumbers(l1, l2)
	l3.Print()
}
