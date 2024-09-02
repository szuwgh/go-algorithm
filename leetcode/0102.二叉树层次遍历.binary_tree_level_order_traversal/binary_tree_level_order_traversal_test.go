package binary_tree_level_order_traversal

import (
	"fmt"
	"testing"
)

func Test_levelOrderTraversal(t *testing.T) {
	n := &treeNode{3, nil, nil}
	n.Left = &treeNode{9, nil, nil}
	n.Right = &treeNode{20, nil, nil}
	n.Right.Left = &treeNode{15, nil, nil}
	n.Right.Right = &treeNode{7, nil, nil}

	res := levelOrderTraversal(n)
	for _, v := range res {
		fmt.Println(v)
	}
}
