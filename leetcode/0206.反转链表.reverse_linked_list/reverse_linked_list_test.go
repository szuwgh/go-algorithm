package reverse_linked_list

import (
	"testing"
)

func Test_reverseLinkedList(t *testing.T) {
	l1 := &linkNode{Value: 1}
	l1.Add(2)
	l1.Add(3)
	l1.Add(4)
	l1.Add(5)
	l := reverseLinkedList(l1)
	l.Print()
	// fmt.Print("\n")

	// l2 := &linkNode{Value: 1}
	// l2.Add(2)
	// l2.Add(3)
	// l2.Add(4)
	// l2.Add(5)

	// l3 := reverseLinkedList2(l2)
	// l3.Print()

}
