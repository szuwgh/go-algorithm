package binary_tree_preorder_traversal

import (
	"fmt"
	"testing"
)

func Test_BinaryTreePreorderTraversal(t *testing.T) {
	n := &treeNode{3, nil, nil}
	n.Left = &treeNode{9, nil, nil}
	n.Right = &treeNode{20, nil, nil}
	n.Right.Left = &treeNode{15, nil, nil}
	n.Right.Right = &treeNode{7, nil, nil}

	res := preorderTraversal(n)
	fmt.Println(res)

}
