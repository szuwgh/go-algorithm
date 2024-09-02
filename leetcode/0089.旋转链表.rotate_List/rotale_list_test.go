package rotate_List

import "testing"

func Test_rotateList(t *testing.T) {
	l1 := &linkNode{Value: 1}
	l1.Add(2)
	l1.Add(3)
	l1.Add(4)
	l1.Add(5)
	l := rotateList(l1, 2)
	l.Print()
}
