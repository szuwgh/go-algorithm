package binary_tree_inorder_traversal

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	n := &treeNode{3, nil, nil}
	n.Left = &treeNode{9, nil, nil}
	n.Right = &treeNode{20, nil, nil}
	n.Left.Left = &treeNode{8, nil, nil}
	n.Left.Left.Left = &treeNode{10, nil, nil}
	n.Left.Left.Right = &treeNode{11, nil, nil}

	n.Left.Right = &treeNode{6, nil, nil}
	n.Left.Right.Left = &treeNode{12, nil, nil}
	n.Left.Right.Right = &treeNode{13, nil, nil}

	n.Right.Left = &treeNode{15, nil, nil}
	n.Right.Left.Left = &treeNode{14, nil, nil}
	n.Right.Left.Right = &treeNode{16, nil, nil}

	n.Right.Right = &treeNode{7, nil, nil}
	n.Right.Right.Left = &treeNode{17, nil, nil}
	n.Right.Right.Right = &treeNode{18, nil, nil}
	res := inorderTraversal(n)
	fmt.Println(res)
}
