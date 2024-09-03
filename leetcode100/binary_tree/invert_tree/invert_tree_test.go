package invert_tree

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	// 示例 1
	root1 := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}}
	fmt.Print("输入：")
	printTree(root1)
	invertedRoot1 := invertTree(root1)
	fmt.Print("输出：")
	printTree(invertedRoot1)

	// 示例 2
	root2 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Print("输入：")
	printTree(root2)
	invertedRoot2 := invertTree(root2)
	fmt.Print("输出：")
	printTree(invertedRoot2)

	// 示例 3
	var root3 *TreeNode
	fmt.Print("输入：")
	printTree(root3)
	invertedRoot3 := invertTree(root3)
	fmt.Print("输出：")
	printTree(invertedRoot3)
}
