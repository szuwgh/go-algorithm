package swap_nodes_in_pairs

import "testing"

func Test_SwapNodesInPairs(t *testing.T) {
	l1 := &linkNode{Value: 1}
	l1.Add(2)
	l1.Add(3)
	l1.Add(4)
	l := swapPairs(l1)
	l.Print()
}
