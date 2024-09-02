package remove_nth_node_from_end_of_list

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	//1->2->3->4->5
	newNode := &LinkNode{1, nil}
	newNode.Add(2)
	newNode.Add(3)
	newNode.Add(4)
	newNode.Add(5)
	node := removeNthFromEnd(newNode, 2)
	node.Print()

}
