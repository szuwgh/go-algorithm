package maximum_depth_of_binary_tree

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	n := &treeNode{3, nil, nil}
	n.Left = &treeNode{9, nil, nil}
	n.Right = &treeNode{20, nil, nil}
	n.Right.Left = &treeNode{15, nil, nil}
	n.Right.Right = &treeNode{7, nil, nil}
	level := maxDepth(n)
	fmt.Println(level)
}
