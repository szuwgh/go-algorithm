package validate_binary_search_tree

import (
	"fmt"
	"testing"
)

func Test_validateBinarySearchTree(t *testing.T) {
	n := &treeNode{2, nil, nil}
	n.Left = &treeNode{1, nil, nil}
	n.Right = &treeNode{2, nil, nil}
	res := validateBinarySearchTree(n)
	fmt.Print(res)

}
