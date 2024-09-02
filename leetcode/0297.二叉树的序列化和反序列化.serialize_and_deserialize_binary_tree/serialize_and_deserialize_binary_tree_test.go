package serialize_and_deserialize_binary_tree

import (
	"fmt"
	"testing"
)

func Test_serialize(t *testing.T) {
	n := &treeNode{3, nil, nil}
	n.Left = &treeNode{9, nil, nil}
	n.Right = &treeNode{20, nil, nil}
	n.Left.Right = &treeNode{7, nil, nil}
	n.Right.Left = &treeNode{22, nil, nil}
	n.Right.Right = &treeNode{21, nil, nil}
	res := serialize(n)
	fmt.Println(res)
}
