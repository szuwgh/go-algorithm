package reverse_nodes_in_k_group

import "testing"

func Test_reverse_nodes_in_k_group(t *testing.T) {
	l1 := &linkNode{Value: 1}
	l1.Add(2)
	l1.Add(3)
	l1.Add(4)
	l1.Add(5)
	l1.Add(6)
	l2 := reverseNodesInKGroup(l1, 3)
	l2.Print()
}
