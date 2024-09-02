package symmetric_tree

import (
	"fmt"
	"testing"
)

func Test_SymmtricTree(t *testing.T) {
	n := &treeNode{1, nil, nil}
	n.Left = &treeNode{2, nil, nil}
	n.Right = &treeNode{2, nil, nil}
	n.Left.Left = &treeNode{3, nil, nil}
	n.Left.Right = &treeNode{4, nil, nil}
	n.Right.Left = &treeNode{4, nil, nil}
	n.Right.Right = &treeNode{3, nil, nil}
	res := isSymmetricTree(n)
	fmt.Println(res)
}
